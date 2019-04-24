package main

import (
	"fmt"
	"time"

	"github.com/Min-Feng/actor"
)

type MyMsg struct {
	msg       string
	agentName string
}

func (m MyMsg) Behavior() {
	fmt.Printf("agent ID {%s} receive message number{%s}\n", m.agentName, m.msg)
}

func main() {
	Sys := actor.System()
	agentMax := 1000
	agents := make([]*actor.Actor, 0, agentMax)

	for i := agentMax - 1; 0 <= i; i-- {
		agentID := fmt.Sprintf("%d", i)
		agent := Sys.ActorOf(agentID)
		agents = append(agents, agent)
	}

	for i, agent := range agents {
		msg := MyMsg{
			msg:       fmt.Sprintf("%d", i),
			agentName: agent.Name(),
		}
		go Sys.SendTo(agent, msg)
	}

	time.Sleep(1 * time.Second)
}
