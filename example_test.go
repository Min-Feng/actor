package actor_test

import (
	"fmt"
	"time"

	"github.com/Min-Feng/actor"
)

type circle struct {
	r     float32
	index int
}

func (c circle) Behavior() {
	fmt.Printf("paint %d index circle that area is %f \n", c.index, c.r*c.r*3.14)
}

func Example() {
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

	// Output:
	// Actor{/Painter} had been closed ,doesn't receive message = main.circle{r:6, index:6}
	// Actor{/Painter} had been closed ,doesn't receive message = main.circle{r:7, index:7}
	// Actor{/Painter} had been closed ,doesn't receive message = main.circle{r:8, index:8}
	// Actor{/Painter} had been closed ,doesn't receive message = main.circle{r:9, index:9}
	// Actor{/Game/Chess} is closed
	// Actor{/Painter/Assistant} is closed
	// Actor{/Painter} is closed
	// paint 0 index circle that area is 0.000000
	// paint 1 index circle that area is 3.140000
	// paint 2 index circle that area is 12.560000
	// paint 3 index circle that area is 28.260000
	// paint 4 index circle that area is 50.240002
	// paint 5 index circle that area is 78.500000
	// {/}:
	// 		is Dead? false
	// {/Game}:
	// 		is Dead? false
	// {/Game/Chess}:
	// 		is Dead? true
	// {/Game/Werewolf}:
	// 		is Dead? false
	// {/Painter}:
	// 		is Dead? true
	// {/Painter/Assistant}:
	// 		is Dead? true
}
