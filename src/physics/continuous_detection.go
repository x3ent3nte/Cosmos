package physics

import (
	"agent"
	"mymath"
	"math"
	"vec"
	"mysort"
)

func HandleMovement(ents []agent.Entity, time_delta float64) {
	mysort.Insertion(agent.SortEntityXPos(ents))
	collisions := predictCollisions(ents, time_delta)
	mysort.QuickSort(SortEntityPairTime(collisions))
	time_passed := 0.0
	for _, col := range collisions {
		time_seconds := (col.time - time_passed) / 1000.0
		for _, ent := range ents {
			ent.Move(time_seconds)
		}
		entityCollide(col.one, col.two)
		time_passed += col.time
	}
	time_left := (time_delta - time_passed) / 1000.0
	for _, ent := range ents {
			ent.Move(time_left)
	}

}

func predictCollisions(ents []agent.Entity, time_delta float64) []EntityPairTime {
	time_seconds := time_delta / 1000.0
	collision_times := make([]EntityPairTime, 0)
	potential := potentialCollisions(ents, time_delta)
	for _, pair := range potential {
		happens, time1, time2 := TimeOfImpact(pair.one, pair.two)
		if happens {
			impact_time := math.Min(math.Max(time1, 0.0), math.Max(time2, 0.0))
			if impact_time > 0.0 && impact_time <= time_seconds {
				collision := EntityPairTime{pair.one, pair.two, impact_time}
				collision_times = append(collision_times, collision)
			}
		}
	}
	return collision_times
}

func potentialCollisions(ents []agent.Entity, time_delta float64) []EntityPairTime {
	pairs := make([]EntityPairTime, 0)
	for i := 0; i < len(ents); i++ {
		one := ents[i]
		for j := i + 1; j < len(ents); j++ {
			two := ents[j] 
			if math.Abs(one.GetPos().X - two.GetPos().X) > (one.Radius() + two.Radius()) * 3 {
				break;
			} else { 
				pairs = append(pairs, EntityPairTime{one, two, 0.0})
			}
		}
	}
	return pairs
}

func TimeOfImpact(one agent.Entity, two agent.Entity) (bool, float64, float64) {
	distance_sq := math.Pow(one.Radius() + two.Radius(), 2)

	x_quad := vec.Vec3{math.Pow(one.Velocity().X - two.Velocity().X, 2), 2 * (one.GetPos().X - two.GetPos().X) * (one.Velocity().X - two.Velocity().X), math.Pow(one.GetPos().X - two.GetPos().X, 2)}
	y_quad := vec.Vec3{math.Pow(one.Velocity().Y - two.Velocity().Y, 2), 2 * (one.GetPos().Y - two.GetPos().Y) * (one.Velocity().Y - two.Velocity().Y), math.Pow(one.GetPos().Y - two.GetPos().Y, 2)}
	z_quad := vec.Vec3{math.Pow(one.Velocity().Z - two.Velocity().Z, 2), 2 * (one.GetPos().Z - two.GetPos().Z) * (one.Velocity().Z - two.Velocity().Z), math.Pow(one.GetPos().Z - two.GetPos().Z, 2)}
	quad := vec.Vec3Add(vec.Vec3Add(x_quad, y_quad), z_quad)
	quad.Z -= distance_sq
	return mymath.SolveQuadratic(quad.X, quad.Y, quad.Z)
}















