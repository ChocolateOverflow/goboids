package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type Boid struct {
	pos Vector
	vel Vector
}

func newBoid(pos Vector) *Boid {
	return &Boid{pos, Vector{(rand.Float64() - 0.5) * initVelocityScalar, (rand.Float64() - 0.5) * initVelocityScalar}}
}

func (this *Boid) print() {
	fmt.Printf("{%f,%f} {%f,%f}\n", this.pos.x, this.pos.y, this.vel.x, this.vel.y)
}

func (this *Boid) doCohesion(boids []*Boid) *Boid {
	center := Vector{0, 0}
	num := 0.

	for _, boid := range boids {
		d := distance(this.pos, boid.pos)
		if this != boid && d >= cohesionRangeMin && d <= cohesionRangeMax {
			center.add(boid.pos)
			num++
		}
	}

	if num > 0 {
		accel := difference(*center.scale(1 / num), this.pos)
		this.vel.add(*accel.scale(cohesionForce))
	}

	return this
}

func (this *Boid) doSeparation(boids []*Boid) *Boid {
	movement := Vector{0, 0}
	for _, boid := range boids {
		if this != boid && distance(this.pos, boid.pos) <= separationRange {
			movement.add(difference(this.pos, boid.pos))
		}
	}
	this.vel.add(*movement.scale(separationForce))
	return this
}

func (this *Boid) matchVelocity(boids []*Boid) *Boid {
	targetVelocity := Vector{0, 0}
	num := 0.

	for _, boid := range boids {
		d := distance(this.pos, boid.pos)
		if this != boid && d >= velocityMatchingRangeMin && d <= velocityMatchingRangeMax {
			targetVelocity.add(boid.vel)
			num++
		}
	}

	if num > 0 {
		targetVelocity.scale(1 / num)
		accel := difference(targetVelocity, this.vel)
		this.vel.add(*accel.scale(velocityMatchingForce))
	}

	return this
}

func (this *Boid) stayWithinBounds() *Boid {
	if this.pos.x < margin {
		this.vel.x += edgeAvoidanceForce
	} else if this.pos.x > width-margin {
		this.vel.x -= edgeAvoidanceForce
	}

	if this.pos.y < margin {
		this.vel.y += edgeAvoidanceForce
	} else if this.pos.y > height-margin {
		this.vel.y -= edgeAvoidanceForce
	}

	return this
}

func (this *Boid) limitSpeed() *Boid {
	speed := this.vel.magnitude()
	if speed > maxSpeed {
		this.vel.x = (this.vel.x / speed) * maxSpeed
		this.vel.y = (this.vel.y / speed) * maxSpeed
	}

	return this
}

func (this *Boid) applyMovement() *Boid {
	this.pos.x += this.vel.x
	this.pos.y += this.vel.y
	return this
}

func (this *Boid) updateBoid(boids []*Boid, wg *sync.WaitGroup) {
	defer wg.Done()
	this.doCohesion(boids)
	this.doSeparation(boids)
	this.matchVelocity(boids)
	this.stayWithinBounds().limitSpeed().applyMovement()
}

type BoidFactory struct {
	pos   Vector
	boids []*Boid
}

func newBoidFactory(x float64, y float64) *BoidFactory {
	return &BoidFactory{Vector{x, y}, make([]*Boid, 0)}
}

func (this *BoidFactory) printBoids() {
	for _, boid := range this.boids {
		boid.print()
	}
}

func (this *BoidFactory) cloneBoids() []*Boid {
	result := make([]*Boid, len(this.boids))
	copy(result, this.boids)
	return result
}

func (this *BoidFactory) spawnBoids(c int) *BoidFactory {
	newBoids := make([]*Boid, c)
	for i := 0; i < c; i++ {
		newBoids[i] = newBoid(this.pos)
	}
	this.boids = append(this.boids, newBoids...)
	return this
}

func (this *BoidFactory) updateBoids() {
	var wg *sync.WaitGroup = new(sync.WaitGroup)
	wg.Add(len(this.boids))
	clone := this.cloneBoids()

	for _, boid := range this.boids {
		go boid.updateBoid(clone, wg)
	}
}
