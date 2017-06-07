package agent

import (
	"vec"
	"sync"
	"math"
	"bit"
)

type Player struct {
	*Agent
	client_id int64
	keycode int
}

func (player *Player) Act(time_delta float64) {
	player.calculateDirectionVectors()
	if bit.IsBitOneAt(player.keycode, 0) {
		impulse := player.rocket.Thrust(vec.Vec3Normal(player.forward), 1.0, time_delta)
		player.applyImpulse(impulse)
	}
	if bit.IsBitOneAt(player.keycode, 1) {
		impulse := player.rocket.Thrust(vec.Vec3Normal(vec.Vec3Scale(player.forward, -1.0)), 1.0, time_delta)
		player.applyImpulse(impulse)
	}
	if bit.IsBitOneAt(player.keycode, 2) {
		impulse := player.rocket.Thrust(vec.Vec3Normal(vec.Vec3Scale(player.right, -1.0)), 1.0, time_delta)
		player.applyImpulse(impulse)
	}
	if bit.IsBitOneAt(player.keycode, 3) {
		impulse := player.rocket.Thrust(vec.Vec3Normal(player.right), 1.0, time_delta)
		player.applyImpulse(impulse)
	}
	if bit.IsBitOneAt(player.keycode, 4) {
		impulse := player.rocket.Thrust(vec.Vec3Normal(player.up), 1.0, time_delta)
		player.applyImpulse(impulse)
	}
	if bit.IsBitOneAt(player.keycode, 5) {
		impulse := player.rocket.Thrust(vec.Vec3Normal(vec.Vec3Scale(player.up, -1.0)), 1.0, time_delta)
		player.applyImpulse(impulse)
	}
}

func (player *Player) UpdateKeyCode(new_keycode int) {
	player.keycode = new_keycode
}

func SpawnPlayer(odin *Odin, client_id int64, id int64, pos vec.Vec3) *Player {
	agent := Agent{
		sync.RWMutex{},
		odin,
		"player",
		id,

		pos,
		vec.Vec3{0.0, 0.0, 0.0},

		vec.Vec3{0.0, 0.0, -1.0},
		vec.Vec3{0.0, 1.0, 0.0},
		vec.Vec3{1.0, 0.0, 0.0},

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
	return &Player{&agent, client_id, 0}
}










