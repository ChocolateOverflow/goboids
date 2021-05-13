package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

// config
const (
	boidCount                        = 100
	width                    float64 = 1000
	height                   float64 = 1000
	maxSpeed                 float64 = 5
	initVelocityScalar       float64 = 1  // how fast should the be at the start
	margin                   float64 = 10 // distance from edge to start avoiding edge
	cohesionRangeMin         float64 = 50
	cohesionRangeMax         float64 = 100
	cohesionForce            float64 = 0.001
	separationRange          float64 = 20
	separationForce          float64 = 0.08
	velocityMatchingRangeMin float64 = 50
	velocityMatchingRangeMax float64 = 100
	velocityMatchingForce    float64 = 0.01
	edgeAvoidanceForce       float64 = 1
)

func main() {
	ebiten.SetWindowSize(int(width), int(height))
	ebiten.SetWindowTitle("Boids in go!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
