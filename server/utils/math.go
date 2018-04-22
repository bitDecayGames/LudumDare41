package utils

type Vector struct {
	X int
	Y int
}

func VecAdd(v1 Vector, v2 Vector) Vector {
	return Vector{
		X: v1.X + v2.X,
		Y: v1.Y + v2.Y,
	}
}
