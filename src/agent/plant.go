package agent

import (
	"vec"
	"sync"
	"math"
)

type Plant struct {
	*Agent
}

func (plant *Plant) Act(time_delta float64) {
	if plant.age > plant.lifespan  {
		plant.alive = false
		return
	} else {
		plant.calculateMovement(time_delta)
	}
}

func SpawnPlant(odin *Odin, id int64, pos vec.Vec3) *Plant{
	target := vec.Vec3Add(pos, vec.Vec3Random(9000))
	agent := Agent{
		sync.RWMutex{},
		odin,
		"plant",
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
		90.0, 
		300.0,

		true, 
		0, 
		10000,

		CreateRocket()}
	return &Plant{&agent}
}