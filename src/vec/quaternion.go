package vec

import (
	"math"
	"strconv"
	"fmt"
)

type Quaternion struct {
	R float64
	I float64
	J float64
	K float64
}

func QuaternionCreate(theta float64, axis Vec3) Quaternion {
	var ijk = Vec3Scale(axis, math.Sin(theta / 2))
	return QuaternionNormal(Quaternion{math.Cos(theta / 2), ijk.X, ijk.Y, ijk.Z})
}

func AxisAngleRotation(point Vec3, theta float64, axis Vec3) Vec3{
	var q = QuaternionCreate(theta, axis)
	return QuaternionRotation(point, q)
}

func QuaternionRotation(point Vec3, q Quaternion) Vec3 {
	var q_inverse = QuaternionInverse(q)
	var point_quaternion = QuaternionFromVec3(point)
	var rotated = HamiltonProduct(HamiltonProduct(q, point_quaternion), q_inverse)
	return Vec3Normal(IJKFromQuaternion(rotated))
}

func QuaternionInverse(q Quaternion) Quaternion {
	return Quaternion{q.R, -q.I, -q.J, -q.K}
}

func QuaternionFromVec3(v Vec3) Quaternion {
	return Quaternion{0, v.X, v.Y, v.Z}
}

func IJKFromQuaternion(q Quaternion) Vec3 {
	return Vec3{q.I, q.J, q.K}
}

func FURFromPYR(pyr Vec3) (Vec3, Vec3, Vec3) {
	pitch := pyr.X
	yaw := pyr.Y
	roll := pyr.Z

	z_axis := Vec3{0.0, 0.0, -1.0}	
	y_axis := Vec3{0.0, 1.0, 0.0}
	x_axis := Vec3{1.0, 0.0, 0.0}

	q_yaw := QuaternionCreate(yaw, y_axis)
	q_pitch := QuaternionCreate(pitch, x_axis)
	q_roll := QuaternionCreate(-roll, z_axis)

	q_combine := HamiltonProduct(q_yaw, HamiltonProduct(q_pitch, q_roll))
	forward := QuaternionRotation(z_axis, q_combine)
	up := QuaternionRotation(y_axis, q_combine)
	right := QuaternionRotation(x_axis, q_combine)
	return forward, up, right
}

func FURFromPYR2(pyr Vec3) (Vec3, Vec3, Vec3) {
	pitch := pyr.X
	yaw := pyr.Y
	roll := pyr.Z

	z_axis := Vec3{0.0, 0.0, -1.0}	
	y_axis := Vec3{0.0, 1.0, 0.0}
	x_axis := Vec3{1.0, 0.0, 0.0}

	z_axis2 := AxisAngleRotation(z_axis, yaw, y_axis)
	x_axis2 := AxisAngleRotation(x_axis, yaw, y_axis)

	y_axis2 := AxisAngleRotation(y_axis, pitch, x_axis2)
	z_axis3 := AxisAngleRotation(z_axis2, pitch, x_axis2)

	x_axis3 := AxisAngleRotation(x_axis2, -roll, z_axis3)
	y_axis3 := AxisAngleRotation(y_axis2, -roll, z_axis3)

	return z_axis3, y_axis3, x_axis3
}

func AxisAngleFromQuaternion(q Quaternion) (float64, Vec3) {
	fmt.Println("q: ", q)
	theta := math.Acos(q.R) * 2
	if theta == 0.0 {
		return theta, Vec3{0.0, 0.0, 0.0}
	} else {
		return theta, Vec3Scale(IJKFromQuaternion(q), 1 / math.Sin(theta / 2))
	}
}

func QuaternionBetweenVectors(a Vec3, b Vec3) Quaternion {
	cross := Vec3Cross(a, b)
	q := QuaternionFromVec3(cross)
	q.R = math.Sqrt(math.Pow(Vec3Mag(a), 2) * math.Pow(Vec3Mag(b), 2)) + Vec3Dot(a, b)
	return QuaternionNormal(q)
}

func QuaternionNormal(q Quaternion) Quaternion {
	return QuaterionScale(q, 1 / QuaterionMag(q))
}

func QuaterionMag(q Quaternion) float64 {
	return math.Sqrt((q.R * q.R) +  (q.I * q.I) + (q.J * q.J) + (q.K * q.K))
}

func QuaterionScale(q Quaternion, f float64) Quaternion {
	return Quaternion{q.R * f, q.I * f, q.J * f, q.K * f}
}
 
func HamiltonProduct(a Quaternion, b Quaternion) Quaternion {
	var r = (a.R * b.R) - (a.I * b.I) - (a.J * b.J) - (a.K * b.K)
	var i = (a.R * b.I) + (a.I * b.R) + (a.J * b.K) - (a.K * b.J)
	var j = (a.R * b.J) - (a.I * b.K) + (a.J * b.R) + (a.K * b.I)
	var k = (a.R * b.K) + (a.I * b.J) - (a.J * b.I) + (a.K * b.R)

	return Quaternion{r, i, j, k}
}

func QuaternionToString(q Quaternion) string {
	return strconv.FormatFloat(q.R, 'f', -1, 64) + " " + 
	strconv.FormatFloat(q.I, 'f', -1, 64) + " " +
	strconv.FormatFloat(q.J, 'f', -1, 64) + " " +
	strconv.FormatFloat(q.K, 'f', -1, 64)
}









