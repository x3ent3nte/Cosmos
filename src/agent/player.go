package agent

import (
	"vec"
	"sync"
	"bit"
	//"fmt"
	//"math"
)

type Player struct {
	*Agent
	client_id int64
	keycode int
}

func (player *Player) Act(time_delta float64) {
	//fmt.Println("Speed: ", vec.Vec3Mag(player.velocity), "m/s")
	player.keyActions(time_delta)
}

func (player *Player) keyActions(time_delta float64) {
	if bit.IsBitOneAt(player.keycode, 0) { // W
		impulse := player.rocket.Thrust(vec.Vec3Normal(player.Forward), 1.0, time_delta)
		player.applyImpulse(impulse)
	}
	if bit.IsBitOneAt(player.keycode, 1) { // S
		impulse := player.rocket.Thrust(vec.Vec3Normal(vec.Vec3Scale(player.Forward, -1.0)), 1.0, time_delta)
		player.applyImpulse(impulse)
	}
	if bit.IsBitOneAt(player.keycode, 2) { // A
		impulse := player.rocket.Thrust(vec.Vec3Normal(vec.Vec3Scale(player.right, -1.0)), 1.0, time_delta)
		player.applyImpulse(impulse)
	}
	if bit.IsBitOneAt(player.keycode, 3) { // D
		impulse := player.rocket.Thrust(vec.Vec3Normal(player.right), 1.0, time_delta)
		player.applyImpulse(impulse)
	}
	if bit.IsBitOneAt(player.keycode, 4) { //Q
		impulse := player.rocket.Thrust(vec.Vec3Normal(player.Up), 1.0, time_delta)
		player.applyImpulse(impulse)
	}
	if bit.IsBitOneAt(player.keycode, 5) { //E
		impulse := player.rocket.Thrust(vec.Vec3Normal(vec.Vec3Scale(player.Up, -1.0)), 1.0, time_delta)
		player.applyImpulse(impulse)
	}

	turn_rate := 1 * time_delta
	if bit.IsBitOneAt(player.keycode, 6) { // I pitch down
		player.Forward = vec.AxisAngleRotation(player.Forward, -turn_rate, player.right)
		player.Up = vec.AxisAngleRotation(player.Up, -turn_rate, player.right)
		//delta := vec.QuaternionCreate(-turn_rate, player.Right)
		//player.AddAngularVelocity(delta)
	} 
	if bit.IsBitOneAt(player.keycode, 7) { // K pitch up
		player.Forward = vec.AxisAngleRotation(player.Forward, turn_rate, player.right)
		player.Up = vec.AxisAngleRotation(player.Up, turn_rate, player.right)
		//delta := vec.QuaternionCreate(turn_rate, player.Right)
		//player.AddAngularVelocity(delta)
	} 
	if bit.IsBitOneAt(player.keycode, 8) { // J yaw left
		player.Forward = vec.AxisAngleRotation(player.Forward, turn_rate, player.Up)
		player.right = vec.AxisAngleRotation(player.right, turn_rate, player.Up)
		//delta := vec.QuaternionCreate(turn_rate, player.Up)
		//player.AddAngularVelocity(delta)
	}
	if bit.IsBitOneAt(player.keycode, 9) { // L yaw right
		player.Forward = vec.AxisAngleRotation(player.Forward, -turn_rate, player.Up)
		player.right = vec.AxisAngleRotation(player.right, -turn_rate, player.Up)
		//delta := vec.QuaternionCreate(-turn_rate, player.Up)
		//player.AddAngularVelocity(delta)
	}
	if bit.IsBitOneAt(player.keycode, 10) { // U roll left
		player.Up = vec.AxisAngleRotation(player.Up, -turn_rate, player.Forward)
		player.right = vec.AxisAngleRotation(player.right, -turn_rate, player.Forward)
		//delta := vec.QuaternionCreate(-turn_rate, player.Forward)
		//player.AddAngularVelocity(delta)
	}
	if bit.IsBitOneAt(player.keycode, 11) { // O roll right
		player.Up = vec.AxisAngleRotation(player.Up, turn_rate, player.Forward)
		player.right = vec.AxisAngleRotation(player.right, turn_rate, player.Forward)
		//delta := vec.QuaternionCreate(turn_rate, player.Forward)
		//player.AddAngularVelocity(delta)
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

		vec.Vec3{0.0, 0.0, 0.0}, 
		vec.QuaternionCreate(0.0, vec.Vec3{0.0, 0.0, -1.0}),
		vec.QuaternionCreate(0.00, vec.Vec3{1.0, 0.0, 0.0}),

		90.0, 
		300.0,

		true, 
		0, 
		10000,
		CreateRocket()}
	return &Player{&agent, client_id, 0}
}










