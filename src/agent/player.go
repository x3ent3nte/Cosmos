package agent

import (
	"vec"
	"sync"
	"bit"
)

type Player struct {
	*Agent
	client_id int64
	keycode int
}

func (player *Player) Act(time_delta float64) {
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
		//player.angular_velocity.Y = 1
		impulse := player.rocket.Thrust(vec.Vec3Normal(player.up), 1.0, time_delta)
		player.applyImpulse(impulse)
	}
	if bit.IsBitOneAt(player.keycode, 5) { //E
		//player.angular_velocity.Y = -1
		impulse := player.rocket.Thrust(vec.Vec3Normal(vec.Vec3Scale(player.up, -1.0)), 1.0, time_delta)
		player.applyImpulse(impulse)
	}

	if bit.IsBitOneAt(player.keycode, 6) { // I
		//pitch down
		player.Euler.X -= 0.01
		//player.angular_velocity.X = -1
	} 
	if bit.IsBitOneAt(player.keycode, 7) { // K
		//pitch up
		player.Euler.X += 0.01
		//player.angular_velocity.X = 1
	} 
	if bit.IsBitOneAt(player.keycode, 8) { // J
		//yaw left
		player.Euler.Y += 0.01
		//player.angular_velocity.Y = 1
	}
	if bit.IsBitOneAt(player.keycode, 9) { // L
		//yaw right
		player.Euler.Y -= 0.01
		//player.angular_velocity.Y = -1
	}
	if bit.IsBitOneAt(player.keycode, 10) { // U
		//roll left
		player.Euler.Z += 0.01
	}
	if bit.IsBitOneAt(player.keycode, 11) { // O
		//roll right
		player.Euler.Z -= 0.01
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










