package agent

import (
	"vec"
	"sync"
	"math"
)

type Parasite struct{
	*Agent
}

func (para *Parasite) Act(time_delta float64) {
	if para.age > para.lifespan  {
		para.alive = false
		return
	} else {
		para.calculateMovement(time_delta)
	}
}

func SpawnParasite(odin *Odin, id int64, pos vec.Vec3) *Parasite{
	target := vec.Vec3Add(pos, vec.Vec3Random(9000))
	agent := Agent{
		sync.RWMutex{},
		odin,
		"parasite",
		id,

		pos,
		target,
		vec.Vec3{0.0, 0.0, -1.0},
		vec.Vec3{0.0, 1.0, 0.0},
		vec.Vec3{1.0, 0.0, 0.0},
		0.0,
		vec.Euler{math.Pi * 1.5, 0.0, 0.0}, 

		vec.Vec3{0.0, 0.0, 0.0},
		vec.Vec3{0.0, 0.0, 0.0},
		vec.Vec3{0.0, 0.0, 0.0},
		90.0, 
		300.0,

		true, 
		0, 
		10000,

		CreateRocket()}
	return &Parasite{&agent}
}