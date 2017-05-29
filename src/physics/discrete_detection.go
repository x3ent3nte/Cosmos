package physics

import (
	"agent"
	"vec"
	"mysort"
)

func DetectCollisions(ents []agent.Entity) {
	mysort.Insertion(agent.SortEntityXPos(ents))
	for i := 0; i < len(ents); i++ {
		one := ents[i]
		one_pos := one.GetPos()
		for j := i + 1; j < len(ents); j++ {
			two := ents[j]
			two_pos := two.GetPos()
			if two_pos.X - one_pos.X > one.Radius() + two.Radius() {
				break
			} else {
				if vec.Vec3DistanceBetween(one_pos, two_pos) < one.Radius() + two.Radius() {
					entityCollide(one, two)
				}
			}
		}
	}
}