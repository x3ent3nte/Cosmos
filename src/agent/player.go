package agent

import (
	"vec"
	"sync"
	"math"
)

type Player struct {
	*Agent
	pressed *[]bool
}

func (player *Player) Act(time_delta float64) {
	if (*player.pressed)[87] { //W
		player.rocket.Thrust(vec.Vec3Normal(player.forward), 1.0, time_delta)
	}
	if (*player.pressed)[83] { //S
	}
	if (*player.pressed)[65] { //A
	}
	if (*player.pressed)[68] { //D
	}
	if (*player.pressed)[81] { //Q
	}
	if (*player.pressed)[69] { //E
	}
}

func (player *Player) UpdatePressed(new_pressed *[]bool) {
	player.pressed = new_pressed
}

func CreatePlayer(odin *Odin, id int64, pos vec.Vec3) *Player {
	agent := Agent{
		sync.RWMutex{},
		odin,
		"player",
		id,

		pos,
		vec.Vec3{0.0, 0.0, 0.0},
		vec.Vec3{0.0, 0.0, -1.0},
		0.0,
		vec.Euler{math.Pi * 1.5, 0.0, 0.0}, 

		vec.Vec3{0.0, 0.0, 0.0},
		vec.Vec3{0.0, 0.0, 0.0},
		90.0, 
		300.0,

		true, 
		0, 
		10000,
		CreateRocket()}
	pressed := make([]bool, 222)
	return &Player{&agent, &pressed}
}










