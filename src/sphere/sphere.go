package sphere

import (
	"errors"
	"math"

	"github.com/calbim/ray-tracer/src/ray"
	"github.com/calbim/ray-tracer/src/tuple"
	"github.com/calbim/ray-tracer/src/intersections"
	uuid "github.com/nu7hatch/gouuid"
)

// Sphere represents a unique sphere
type Sphere struct {
	id string
}

// New returns a new sphere
func New() Sphere {
	id, err := uuid.NewV4()
	if err != nil {
		errors.New("failed to generate a unique identifier for sphere")
	}
	return Sphere{
		id: id.String(),
	}
}

// Intersect returns the points at which a ray intersects a sphere
func Intersect(s Sphere, r ray.Ray) []intersections.Intersection {
	sphereToRay := tuple.Subtract(r.Origin, tuple.Point(0.0, 0.0, 0.0))
	a := tuple.DotProduct(r.Direction, r.Direction)
	b := 2 * tuple.DotProduct(r.Direction, sphereToRay)
	c := tuple.DotProduct(sphereToRay, sphereToRay) - 1
	d := b*b - 4*a*c
	if d < 0 {
		return []intersections.Intersection{}
	}
	i1 := intersections.Intersection{Value:(-b - math.Sqrt(d)) / (2 * a), Object:s}
	i2 := intersections.Intersection{Value:(-b + math.Sqrt(d)) / (2 * a), Object:s}

	return intersections.Intersections(i1,i2)
}
