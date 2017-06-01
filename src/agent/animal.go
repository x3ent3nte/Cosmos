package agent

import (
	"vec"
	"math"
	"sync"
)

type Animal struct{
	*Agent
}

func (ani *Animal) Act(time_delta float64) {
	if ani.age > ani.lifespan  {
		ani.alive = false
		return
	} else {
		ani.calculateMovement(time_delta)
		if ani.age == (ani.lifespan / 10) {
			ani.Mitosis()
		}
		//ent.age++
	}
}

func (ani *Animal) Mitosis() {

}

func SpawnAnimal(id int64, pos vec.Vec3) *Animal{
	target := vec.Vec3Add(pos, vec.Vec3Random(9000))
	agent := Agent{
		sync.RWMutex{},
		"animal",
		id,

		pos,
		target,
		vec.Vec3{0.0, 0.0, -1.0},
		0.0,
		vec.Euler{math.Pi * 1.5, 0.0, 0.0}, 

		vec.Vec3{0.0, 0.0, 0.0},
		vec.Vec3{0.0, 0.0, 0.0},
		90.0, 
		300.0,

		true, 
		0, 
		10000}
	return &Animal{&agent}
}