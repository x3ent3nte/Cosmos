package vec

import (
	"math"
	"math/rand"
	"strconv"
)

type Vec3 struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

func Vec3Add(a Vec3, b Vec3) Vec3 {
	return Vec3{a.X + b.X, a.Y + b.Y, a.Z + b.Z}
}

func Vec3Sub(a Vec3, b Vec3) Vec3 {
	return Vec3{a.X - b.X, a.Y - b.Y, a.Z - b.Z}
}

func Vec3Scale(a Vec3, f float64) Vec3 {
	return Vec3{a.X * f, a.Y * f, a.Z * f}
}

func Vec3Mag(a Vec3) float64 {
	return math.Pow(math.Pow(a.X, 2) + math.Pow(a.Y, 2) + math.Pow(a.Z, 2), 0.5)
}

func Vec3Normal(a Vec3) Vec3 {
	return Vec3Scale(a, 1.0 / Vec3Mag(a))
}

func Vec3Dot(a Vec3, b Vec3) float64 {
	return (a.X * b.X) + (a.Y * b.Y) + (a.Z * b.Z);
}

func Vec3Cross(a Vec3, b Vec3) Vec3 {

	var cx = (a.Y * b.Z) - (a.Z * b.Y)
	var cy = (a.Z * b.X) - (a.X * b.Z)
	var cz = (a.X * b.Y) - (a.Y * b.X)
	return Vec3{cx, cy, cz}
}

func Vec3AngleBetween(a Vec3, b Vec3) float64 {
	var dot = Vec3Dot(a, b)
	var mag_mult = Vec3Mag(a) * Vec3Mag(b)
	if mag_mult == 0 {
		return 0.0
	}
	var ratio = dot / mag_mult
	if ratio > 1 {
		ratio = 1.0
	}
	if ratio < -1 {
		ratio = -1.0
	}
	return math.Acos(ratio)
}

func Vec3Yaw(a Vec3) float64 {
	return -math.Atan2(a.X, -a.Z)
}

func Vec3Pitch(a Vec3) float64 {
	return math.Asin(a.Y)
}

func Vec3XZFrame(a Vec3) Vec3 {
	return Vec3{a.X, 0.0, a.Z}
}

func Vec3XYFrame(a Vec3) Vec3 {
	return Vec3{a.X, a.Y, 0.0}
}

func Vec3YZFrame(a Vec3) Vec3 {
	return Vec3{0.0, a.Y, a.Z}
}

func Vec3DistanceBetween(a Vec3, b Vec3) float64 {
	return math.Pow(math.Pow(a.X - b.X, 2) + math.Pow(a.Y - b.Y, 2) + math.Pow(a.Z - b.Z, 2), 0.5)
}

func Vec3CosineSimilarity(a Vec3, b Vec3) float64 {
	denom := Vec3Mag(a) * Vec3Mag(b)
	if denom == 0 {
		return 0.0
	} else {
		return Vec3Dot(a, b) / denom
	}
}

func Vec3Random(scope float64) Vec3 {
	var yaw = rand.Float64() * 2 * math.Pi;
	var pitch = (rand.Float64() * 2 * math.Pi) - math.Pi;

	var x = math.Cos(yaw) * math.Cos(pitch)
	var y = math.Sin(pitch)
	var z = math.Sin(yaw) * math.Cos(pitch)

	var vec = Vec3{x, y, z}
	return Vec3Scale(vec, RandomFloat64(scope))
}

func RandomFloat64(scope float64) float64 {
	return (rand.Float64() * scope) - (scope / 2)
}

func degree2Radian(degree float64) float64{
	return (degree * math.Pi) / 180
}

func radian2Degree(radian float64) float64{
	return (180 * radian) / math.Pi
}

func Vec3ToString(v Vec3) string {
	return strconv.FormatFloat(v.X, 'f', -1, 64) + " " +
	strconv.FormatFloat(v.Y, 'f', -1, 64) + " " +
	strconv.FormatFloat(v.Z, 'f', -1, 64)
}







