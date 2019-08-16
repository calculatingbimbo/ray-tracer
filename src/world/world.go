package world

import (
	"sort"

	"github.com/calbim/ray-tracer/src/color"
	"github.com/calbim/ray-tracer/src/light"
	"github.com/calbim/ray-tracer/src/material"
	"github.com/calbim/ray-tracer/src/ray"
	"github.com/calbim/ray-tracer/src/shape"
	"github.com/calbim/ray-tracer/src/transforms"
	"github.com/calbim/ray-tracer/src/tuple"
)

// World is a collection of objects and a light source
type World struct {
	Objects []shape.Shape
	Light   *light.Light
}

// Default returns a default World object
func Default() World {
	light := light.PointLight(tuple.Point(-10, 10, -10), color.New(1, 1, 1))
	s1 := shape.NewSphere()
	m := material.New()
	m.Color = color.New(0.8, 1.0, 0.6)
	m.Diffuse = 0.7
	m.Specular = 0.2
	s1.Material = m
	s2 := shape.NewSphere()
	s2.Transform = transforms.Scaling(0.5, 0.5, 0.5)
	return World{
		Light:   &light,
		Objects: []shape.Shape{s1, s2},
	}
}

// Intersect returns the intersections of a collection of objects with a ray
func (w *World) Intersect(r ray.Ray) []shape.Intersection {
	list := []shape.Intersection{}
	for _, o := range w.Objects {
		intersections := o.Intersect(r)
		list = append(list, intersections...)
	}
	sort.Sort(byValue(list))
	return list
}

//ShadeHit returns the shade of a hit
func (w *World) ShadeHit(c shape.Computation) color.Color {
	m := c.Object.GetMaterial()
	l := w.Light
	return m.Lighting(*l, c.Point, c.Eyev, c.Normal, false)
}

//ColorAt returns the color of an intersection
func (w *World) ColorAt(r ray.Ray) color.Color {
	intersections := w.Intersect(r)
	hit := shape.Hit(intersections)
	if hit == nil {
		return color.Black
	}
	comps := hit.PrepareComputations(r)
	return w.ShadeHit(comps)
}

//IsShadowed determines if a point is shadowed in a world
func (w *World) IsShadowed(p tuple.Tuple) bool {
	dV := w.Light.Position.Subtract(p)
	distance := dV.Magnitude()
	r := ray.New(p, dV.Normalize())
	intersections := w.Intersect(r)
	hit := shape.Hit(intersections)
	if hit != nil && hit.Value < distance {
		return true
	}
	return false
}

type byValue []shape.Intersection

func (s byValue) Len() int {
	return len(s)
}
func (s byValue) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byValue) Less(i, j int) bool {
	return s[i].Value < s[j].Value
}
