package silhouette

import (
	"fmt"
	"os"
	"testing"

	"github.com/calbim/ray-tracer/src/canvas"
	"github.com/calbim/ray-tracer/src/color"
	"github.com/calbim/ray-tracer/src/light"
	"github.com/calbim/ray-tracer/src/material"
	"github.com/calbim/ray-tracer/src/ray"
	"github.com/calbim/ray-tracer/src/shape"
	"github.com/calbim/ray-tracer/src/tuple"
)

func TestSilhouette(t *testing.T) {
	c := canvas.New(200, 200)
	sphere := shape.NewSphere()
	sphere.Material = material.New()
	sphere.Material.Color = color.New(1, 0.2, 1)
	light := light.PointLight(tuple.Point(-10, 10, -10), color.White)
	rayOrigin := tuple.Point(0, 0, -5)
	wallZ := 20.0
	wallSize := 14.0
	half := wallSize / 2
	pixelSize := wallSize / 200
	for y := 0; y < 200; y++ {
		worldY := float64(half - pixelSize*float64(y))
		for x := 0; x < 200; x++ {
			worldX := float64(-half + pixelSize*float64(x))
			position := tuple.Point(worldX, worldY, wallZ)
			diff := position.Subtract(rayOrigin)
			r := ray.New(rayOrigin, diff.Normalize())
			xs := sphere.Intersect(r)
			hit := shape.Hit(xs)
			if hit != nil {
				p := r.Position(hit.Value)
				normalv := hit.Object.Normal(p)
				eyev := r.Direction.Negate()
				color := sphere.Material.Lighting(light, p, eyev, *normalv)
				c.WritePixel(x, y, color)
			}
		}
	}
	ppm := c.ToPPM()
	file, err := os.Create("silhouette.ppm")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(file, ppm)
}