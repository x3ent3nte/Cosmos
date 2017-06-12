package agent

import (
	"vec"
	"sync"
	"bit"
	"fmt"
	//"math"
)

type Player struct {
	*Agent
	client_id int64
	keycode int
}

func (player *Player) Act(time_delta float64) {
	fmt.Println("Speed: ", vec.Vec3Mag(player.velocity), "m/s")
	player.keyActions(time_delta)
}

func (player *Player) keyActions(time_delta float64) {
	fmt.Println("Orientation: ", player.Euler)
	fmt.Println("Vyaw: ", vec.Vec3Yaw(player.Forward), " Vpitch: ", vec.Vec3Pitch(player.Forward))
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

	if bit.IsBitOneAt(player.keycode, 6) { // I pitch down
		/*forward := vec.AxisAngleRotation(player.Forward, -0.01, player.right)
		up := vec.AxisAngleRotation(player.up, -0.01, player.right)
		pyr := YPRfromForwardUpRight(forward, up, player.right)
		player.Euler = pyr*/
		player.Euler.X -= 0.01
	} 
	if bit.IsBitOneAt(player.keycode, 7) { // K pitch up
		/*forward := vec.AxisAngleRotation(player.Forward, 0.01, player.right)
		up := vec.AxisAngleRotation(player.up, 0.01, player.right)
		pyr := YPRfromForwardUpRight(forward, up, player.right)
		player.Euler = pyr*/
		player.Euler.X += 0.01
	} 
	if bit.IsBitOneAt(player.keycode, 8) { // J yaw left
		/*forward := vec.AxisAngleRotation(player.Forward, +0.01, player.up)
		player.Euler.X = vec.Vec3Pitch(forward)
		player.Euler.Y = vec.Vec3Yaw(forward)*/
		player.Euler.Y += 0.01
	}
	if bit.IsBitOneAt(player.keycode, 9) { // L yaw right
		/*forward := vec.AxisAngleRotation(player.Forward, -0.01, player.up)
		player.Euler.X = vec.Vec3Pitch(forward)
		player.Euler.Y = vec.Vec3Yaw(forward)*/
		player.Euler.Y -= 0.01
	}
	if bit.IsBitOneAt(player.keycode, 10) { // U roll left
		player.Euler.Z += 0.01
	}
	if bit.IsBitOneAt(player.keycode, 11) { // O roll right
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










