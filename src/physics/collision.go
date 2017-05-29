package physics

import (
	"agent"
	"vec"
)

func entityCollide(one agent.Entity, two agent.Entity) {
	impact_vec := vec.Vec3Sub(two.GetPos(), one.GetPos())

	one_velocity_relative := vec.Vec3Sub(one.Velocity(), two.Velocity())
	one_velocity_ratio := vec.Vec3CosineSimilarity(impact_vec, one_velocity_relative)
	one_velocity_impact := vec.Vec3Scale(one_velocity_relative, one_velocity_ratio)
	if one_velocity_ratio < 0 {
		return
	}

	two_velocity_delta := vec.Vec3Scale(one_velocity_impact, one.Mass() / two.Mass())
	momentum_transfer := vec.Vec3Scale(two_velocity_delta, two.Mass())
	one_velocity_delta := vec.Vec3Scale(vec.Vec3Scale(momentum_transfer, 1 / one.Mass()), -1)

	one_new_pos := vec.Vec3Add(two.GetPos(), vec.Vec3Scale(vec.Vec3Normal(impact_vec), -(one.Radius() + two.Radius())))
	one.SetPos(one_new_pos)
	
	one.AddVelocity(one_velocity_delta)
	two.AddVelocity(two_velocity_delta)
}