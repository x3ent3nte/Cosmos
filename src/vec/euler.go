package vec

import (
	"math"
	"strconv"
)

const two_pi = 2 * math.Pi

type Euler struct {
	Yaw float64 `json:"yaw"`
	Pitch float64 `json:"pitch"`
	Roll float64 `json:"roll"`
}

func CreateEuler() Euler {
	return Euler{0, 0, 0}
}

func (euler *Euler) ForwardVector() Vec3 {
    var x = math.Cos(euler.Yaw) * math.Cos(euler.Pitch)
    var y = math.Sin(euler.Pitch)
    var z = math.Sin(euler.Yaw) * math.Cos(euler.Pitch)
    return Vec3{x, y, z}
}

func (euler *Euler) RotateYaw(delta float64) {
	euler.Yaw += delta

	if euler.Yaw > two_pi {
		//euler.Yaw -= two_pi
	} 
}

func (euler *Euler) RotatePitch(delta float64) {
	euler.Pitch += delta
}

func (euler *Euler) RotateRoll(delta float64) {
	euler.Roll += delta
}

func EulerToString(euler Euler) string {
	return "Yaw: " + strconv.FormatFloat(euler.Yaw, 'f', -1, 64) + 
	" Pitch: " + strconv.FormatFloat(euler.Pitch, 'f', -1, 64) + 
	" Roll: " + strconv.FormatFloat(euler.Roll, 'f', -1, 64)
}
