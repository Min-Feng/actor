// Package actor 模擬 erlang 的 訊息傳遞方式
package actor

import (
	"context"
	"fmt"
	"sync"
)

const (
	defaultMailboxSize = 10
	systemName         = "/"
)

type message interface {
	Behavior()
}

type Actor struct {
	name       string
	mu         sync.RWMutex
	recvChan   chan message
	cancelFunc context.CancelFunc
	ctx        context.Context
}

// System 為Actor模型的頂端節點
func System() *Actor {
	rootCtx := context.Background()

	sysActor := &Actor{
		name:     systemName,
		recvChan: make(chan message, defaultMailboxSize),
		ctx:      rootCtx,
	}

	go actorHandler(sysActor)

	return sysActor
}

// ActorOf 創造子Actor的方法
func (ac *Actor) ActorOf(name string) *Actor {
	childCtx, cancel := context.WithCancel(ac.ctx)

	var childName string
	if ac.name == systemName {
		childName = ac.name + name
	} else {
		childName = ac.name + "/" + name
	}

	childActor := &Actor{
		name:       childName,
		recvChan:   make(chan message, defaultMailboxSize),
		cancelFunc: cancel,
		ctx:        childCtx,
	}

	go actorHandler(childActor)

	return childActor
}

func actorHandler(ac *Actor) {
	for {
		select {
		// 接收其他 actor 送來的訊息
		case msg := <-ac.recvChan:
			msg.Behavior()

		// 停止 actor 的運作
		case <-ac.ctx.Done():
			fmt.Printf("Actor{%s} is closed\n", ac.name)
			close(ac.recvChan)

			// 關閉 actor 前,確保 mailbox 內的值都讀取出來執行一次
			for msg := range ac.recvChan {
				msg.Behavior()
			}
			return
		}
	}
}

// SendTo 指定 sender 發送訊息到 receiver
func (sender *Actor) SendTo(receiver *Actor, msg message) {

	//以此 actor 的身份 對 其他 actor 發送訊息
	sendTo(receiver, msg)
}

func sendTo(receiver *Actor, msg message) {
	// ac.RecvMsg阻塞發生時，利用mutex強制等待訊息接收完畢
	// 才允許stop Actor
	// 不然會發生 panic: send on closed channel
	receiver.mu.RLock()
	defer receiver.mu.RUnlock()

	select {
	case <-receiver.ctx.Done():
		fmt.Printf("Actor{%s} had been closed ,doesn't receive message = %#v \n", receiver.name, msg)
	default:
		receiver.recvChan <- msg
	}
}

func (ac *Actor) Stop() {
	if ac.name == "/" {
		fmt.Printf("ActorSystem don't stop\n")
		return
	}

	ac.mu.Lock()
	defer ac.mu.Unlock()

	select {
	case <-ac.ctx.Done():
		fmt.Printf("Actor{%s} had been closed\n", ac.name)
	default:
		ac.cancelFunc()
	}
}

func (ac *Actor) IsDead() bool {
	select {
	case <-ac.ctx.Done():
		return true
	default:
		return false
	}
}

func (ac *Actor) Name() string {
	return ac.name
}
