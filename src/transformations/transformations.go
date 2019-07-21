package transformations

import (
	"math"

	"github.com/calbim/ray-tracer/src/tuple"

	"github.com/calbim/ray-tracer/src/matrix"
)

// Pi = 3.14
var Pi = math.Pi

// NewTranslation returns a new 4x4 transformation matrix
func NewTranslation(x, y, z float64) [][]float64 {
	return matrix.New([]float64{1, 0, 0, x, 0, 1, 0, y, 0, 0, 1, z, 0, 0, 0, 1}, 4, 4)
}

// NewScaling returns a new 4x4 scaling matrix
func NewScaling(x, y, z float64) [][]float64 {
	return matrix.New([]float64{
		x, 0, 0, 0,
		0, y, 0, 0,
		0, 0, z, 0,
		0, 0, 0, 1},
		4, 4)

}

// RotationX returns a matrix that represents a rotation by r radians around the X axis
func RotationX(r float64) [][]float64 {
	return matrix.New([]float64{
		1, 0, 0, 0,
		0, math.Cos(r), -math.Sin(r), 0,
		0, math.Sin(r), math.Cos(r), 0,
		0, 0, 0, 1,
	}, 4, 4)
}

// RotationY returns a matrix that represents a rotation by r radians around the Y axis
func RotationY(r float64) [][]float64 {
	return matrix.New([]float64{
		math.Cos(r), 0, math.Sin(r), 0,
		0, 1, 0, 0,
		-math.Sin(r), 0, math.Cos(r), 0,
		0, 0, 0, 1,
	}, 4, 4)
}

// RotationZ returns a matrix that represents a rotation by r radians around the Z axis
func RotationZ(r float64) [][]float64 {
	return matrix.New([]float64{
		math.Cos(r), -math.Sin(r), 0, 0,
		math.Sin(r), math.Cos(r), 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}, 4, 4)
}

//NewShearing returns a matrix that represents a shearing (skew) operation
func NewShearing(Xy, Xz, Yx, Yz, Zx, Zy float64) [][]float64 {
	return matrix.New([]float64{
		1, Xy, Xz, 0,
		Yx, 1, Yz, 0,
		Zx, Zy, 1, 0,
		0, 0, 0, 1,
	}, 4, 4)
}

//ViewTransform returns a matrix that represents view parameters
func ViewTransform(from tuple.Tuple, to tuple.Tuple, up tuple.Tuple) [][]float64 {
	forward := tuple.Normalize(tuple.Subtract(to, from))
	left := tuple.CrossProduct(forward, tuple.Normalize(up))
	trueUp := tuple.CrossProduct(left, forward)
	orientation := matrix.New([]float64{
		left.X, left.Y, left.Z, 0,
		trueUp.X, trueUp.Y, trueUp.Z, 0,
		-forward.X, -forward.Y, -forward.Z, 0,
		0, 0, 0, 1}, 4, 4)
	return matrix.Multiply(orientation, NewTranslation(-from.X, -from.Y, -from.Z))
}
