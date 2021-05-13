# goboids

## How does it work?

Each **boid** follows 3 simple rules:

1. Cohesion: Move towards the center of neighboring boids
2. Separation: Keep some distance to avoid crashing into other boids
3. Velocity matching: Match one's own velocity (speed & direction) with the neighboring boids

## Playing around with the parameters

All configurable parameters are set as constants in `main.go`.

## Setup

**goboids** uses [Ebiten](https://github.com/hajimehoshi/ebiten) which requires some dependencies to be installed separately. Please follow the instructions [here](https://ebiten.org/documents/install.html).
