package agent

import (
	"vec"
	"math"
)

type Rocket struct {
	fuel_mass float64
	mass_flow_rate float64
	specific_impulse float64
}

func (rocket *Rocket) Thrust(forward vec.Vec3, usage float64, time_delta float64) vec.Vec3 {
	fuel_used := rocket.mass_flow_rate * time_delta * usage
	fuel_used = math.Min(rocket.fuel_mass, fuel_used)	
	impulse := vec.Vec3Scale(forward, fuel_used * rocket.specific_impulse)
	rocket.fuel_mass -= fuel_used
	return impulse
}

func (rocket *Rocket) isLowOnFuel() bool {
	return rocket.fuel_mass < 100000
}

func CreateRocket() Rocket {
	return Rocket{1500000, 7000, 80}
}