package utils

var DeadVector = Vector{X: -1, Y: -1}

type Vector struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func NewVec(x, y int) Vector {
	return Vector{X: x, Y: y}
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
