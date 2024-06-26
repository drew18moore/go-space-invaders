package utils

import "math"

type Vector struct {
	X float64
	Y float64
}

func NewVector(x, y float64) Vector {
	return Vector{
		X: x,
		Y: y,
	}
}

func (v Vector) Normalize() Vector {
	mag := math.Sqrt(v.X*v.X + v.Y*v.Y)
	return Vector{v.X / mag, v.Y / mag}
}
