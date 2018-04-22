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

func VecEquals(v1 Vector, v2 Vector) bool {
	return v1.X == v2.X && v1.Y == v2.Y
}

func VecScale(v1 Vector, scalar int) Vector {
	return Vector{X: v1.X * scalar, Y: v1.Y * scalar}
}
