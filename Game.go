package main

import (
	"image/color"
	"image/png"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	background = color.Black
	boidImage  *ebiten.Image
)

type Game struct {
	boidFactory *BoidFactory
	options     ebiten.DrawImageOptions
	inited      bool
}

func (this *Game) init() {
	defer func() {
		this.inited = true
	}()

	// set up boids
	this.boidFactory = newBoidFactory(width/2, height/2)
	this.boidFactory.spawnBoids(boidCount)
}

func (this *Game) Update() error {
	if !this.inited {
		this.init()
	}
	this.boidFactory.updateBoids()
	return nil
}

func (this *Game) Draw(screen *ebiten.Image) {
	w, h := boidImage.Size()
	for _, boid := range this.boidFactory.boids {
		this.options.GeoM.Reset()
		this.options.GeoM.Translate(-float64(w)/2, -float64(h)/2)
		this.options.GeoM.Rotate(boid.vel.angle())
		this.options.GeoM.Translate(float64(w)/2, float64(h)/2)
		this.options.GeoM.Translate(boid.pos.x, boid.pos.y)
		screen.DrawImage(boidImage, &this.options)
	}
}

func (this *Game) Layout(outerWidth, outerHeight int) (int, int) {
	return int(width), int(height)
}

func init() {
	imgFile, err := os.Open("boid.png")
	if err != nil {
		log.Fatal(err)
	}
	defer imgFile.Close()

	img, err := png.Decode(imgFile)
	if err != nil {
		log.Fatal(err)
	}
	origImage := ebiten.NewImageFromImage(img)

	w, h := origImage.Size()
	boidImage = ebiten.NewImage(w, h)

	options := &ebiten.DrawImageOptions{}
	options.ColorM.Scale(1, 1, 1, 1)

	boidImage.DrawImage(origImage, options)
}
