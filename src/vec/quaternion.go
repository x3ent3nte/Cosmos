package vec

import (
	"math"
	"strconv"
)

type Quaternion struct {
	R float64
	I float64
	J float64
	K float64
}

func QuaternionCreate(theta float64, axis Vec3) Quaternion {
	var ijk = Vec3Scale(axis, math.Sin(theta / 2))
	return Quaternion{math.Cos(theta / 2), ijk.X, ijk.Y, ijk.Z}
}

func AxisAngleRotation(point Vec3, theta float64, axis Vec3) Vec3{
	var q = QuaternionCreate(theta, axis)
	return QuaternionRotation(point, q)
}

func QuaternionRotation(point Vec3, q Quaternion) Vec3 {
	var q_inverse = QuaternionInverse(q)
	var point_quaternion = QuaternionFromVec3(point)
	var rotated = HamiltonProduct(HamiltonProduct(q, point_quaternion), q_inverse)
	return Vec3FromQuaternion(rotated)
}

func QuaternionInverse(q Quaternion) Quaternion {
	return Quaternion{q.R, -q.I, -q.J, -q.K}
}

func QuaternionFromVec3(v Vec3) Quaternion {
	return Quaternion{0, v.X, v.Y, v.Z}
}

func Vec3FromQuaternion(q Quaternion) Vec3 {
	return Vec3{q.I, q.J, q.K}
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









