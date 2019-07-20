package world

import (
	"errors"
	"fmt"
	"sort"

	"github.com/calbim/ray-tracer/src/ray"

	"github.com/calbim/ray-tracer/src/intersections"
	"github.com/calbim/ray-tracer/src/light"
	"github.com/calbim/ray-tracer/src/material"
	"github.com/calbim/ray-tracer/src/sphere"
	"github.com/calbim/ray-tracer/src/transformations"
	"github.com/calbim/ray-tracer/src/tuple"
)

//World contains a set of objects and a light source
type World struct {
	Objects []intersections.Object
	Light   *light.PointLight
}

//NewDefault creates a world with a default config of one light source and two spheres
func NewDefault() (*World, error) {
	w := World{
		Light: &light.PointLight{Intensity: tuple.Point(-10, -10, -10), Position: tuple.Point(1, 1, 1)},
	}
	m := material.Material{Color: tuple.Color(0.8, 1.0, 0.6), Diffuse: 0.7, Specular: 0.2}
	s1, err := sphere.New()
	if err != nil {
		return nil, errors.New("Could not create a sphere")
	}
	s1.Material = m
	s2, err := sphere.New()
	if err != nil {
		return nil, errors.New("Could not create a sphere")
	}
	s2.Material = m
	s2.SetTransform(transformations.NewScaling(0.5, 0.5, 0.5))
	w.Objects = []intersections.Object{s1, s2}
	return &w, nil
}

// ByIntersectionValue implements sort.Interface for []Intersection based on
// the Value field.
type ByIntersectionValue []intersections.Intersection

func (a ByIntersectionValue) Len() int           { return len(a) }
func (a ByIntersectionValue) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByIntersectionValue) Less(i, j int) bool { return a[i].Value < a[j].Value }

//Intersect returns a list of intersections in sorted order when a ray passes through a world
func (w World) Intersect(r ray.Ray) ([]intersections.Intersection, error) {
	list := []intersections.Intersection{}
	for _, o := range w.Objects {
		intersection, err := intersections.Intersect(o, r)
		if err != nil {
			return nil, fmt.Errorf("Error while computing intersection for object %v", err)
		}
		list = append(list, intersection...)
	}
	sort.Sort(ByIntersectionValue(list))
	return list, nil
}