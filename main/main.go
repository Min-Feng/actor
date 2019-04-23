package main

import (
	"fmt"
	"time"

	"github.com/Min-Feng/actor"
)

type circle struct {
	r     float32
	index int
}

// 由於actor的mailbox通道，我設為interface
// 擁有Behavior()才允許通過
func (c circle) Behavior() {
	fmt.Printf("paint %d index circle that area is %f \n", c.index, c.r*c.r*3.14)
}

func main() {
	Sys := actor.System()

	Game := Sys.ActorOf("Game")
	Chess := Game.ActorOf("Chess")
	Werewolf := Game.ActorOf("Werewolf")

	Painter := Sys.ActorOf("Painter")
	Assistant := Painter.ActorOf("Assistant")

	for i := 0; i < 10; i++ {
		Sys.SendTo(Painter, circle{r: float32(i), index: i})
		if i == 5 {
			Painter.Stop()
			Chess.Stop()
		}
	}

	fmt.Printf("{%s}:\n\tis Dead? %t\n", Sys.Name(), Sys.IsDead())
	fmt.Printf("{%s}:\n\tis Dead? %t\n", Game.Name(), Game.IsDead())
	fmt.Printf("{%s}:\n\tis Dead? %t\n", Chess.Name(), Chess.IsDead())
	fmt.Printf("{%s}:\n\tis Dead? %t\n", Werewolf.Name(), Werewolf.IsDead())
	fmt.Printf("{%s}:\n\tis Dead? %t\n", Painter.Name(), Painter.IsDead())
	fmt.Printf("{%s}:\n\tis Dead? %t\n", Assistant.Name(), Assistant.IsDead())

	time.Sleep(1 * time.Second)
}
