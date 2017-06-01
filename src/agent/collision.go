package agent

import (
	"vec"
)

func bodyCollide(one Entity, two Entity) {
	impact_vec := vec.Vec3Sub(two.GetPos(), one.GetPos())

	one_velocity_relative := vec.Vec3Sub(one.GetVelocity(), two.GetVelocity())
	one_velocity_ratio := vec.Vec3CosineSimilarity(impact_vec, one_velocity_relative)
	one_velocity_impact := vec.Vec3Scale(one_velocity_relative, one_velocity_ratio)
	if one_velocity_ratio < 0 {
		return
	}

	two_velocity_delta := vec.Vec3Scale(one_velocity_impact, one.GetMass() / two.GetMass())
	momentum_transfer := vec.Vec3Scale(two_velocity_delta, two.GetMass())
	one_velocity_delta := vec.Vec3Scale(vec.Vec3Scale(momentum_transfer, 1 / one.GetMass()), -1)

	one_new_pos := vec.Vec3Add(two.GetPos(), vec.Vec3Scale(vec.Vec3Normal(impact_vec), -(one.GetRadius() + two.GetRadius())))
	one.SetPos(one_new_pos)
	
	one.AddVelocity(one_velocity_delta)
	two.AddVelocity(two_velocity_delta)
}