package agent

import (
	"vec"
	"math"
	"fmt"
	"sync"
)

type Parasite struct{
	*Agent
}

func (ani *Parasite) Act(time_delta float64) {
	fmt.Println("Parasite Acted")
}

func SpawnParasite(pos vec.Vec3) *Parasite{
	target := vec.Vec3Add(pos, vec.Vec3Random(9000))
	agent := Agent{
		sync.RWMutex{},
		0,
		pos,
		target,
		vec.Vec3{0.0, 0.0, -1.0},
		0.0,
		vec.Euler{math.Pi * 1.5, 0.0, 0.0}, 
		vec.Vec3{0.0, 0.0, 0.0},
		vec.Vec3{0.0, 0.0, 0.0},
		500000.0, 
		300.0,

		true, 
		0, 
		10000}
	return &Parasite{&agent}
}