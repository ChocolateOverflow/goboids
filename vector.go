package main

import "math"

type Vector struct {
	x float64
	y float64
}

func (this *Vector) add(v Vector) *Vector {
	this.x += v.x
	this.y += v.y
	return this
}

func sum(v1, v2 Vector) Vector {
	return Vector{v1.x + v2.x, v1.y + v2.y}
}

func difference(minuend, subtrahend Vector) Vector {
	return Vector{minuend.x - subtrahend.x, minuend.y - subtrahend.y}
}

func (this *Vector) scale(scalar float64) *Vector {
	this.x *= scalar
	this.y *= scalar
	return this
}

func dotProduct(v1, v2 *Vector) float64 {
	return v1.x*v2.x + v1.y*v2.y
}

func distance(v, u Vector) float64 {
	return math.Sqrt((v.x-u.x)*(v.x-u.x) + (v.y-u.y)*(v.y-u.y))
}

func (this *Vector) magnitude() float64 {
	return math.Sqrt(this.x*this.x + this.y*this.y)
}

func (this *Vector) angle() float64 {
	return math.Atan2(this.y, this.x)
}
