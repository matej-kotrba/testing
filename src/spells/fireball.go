package spells

import (
	"fmt"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const speed = 5

type Fireball struct {
	X float32
	Y  float32
	SX float32
	SY float32
	R int16
	image rl.Image
	texture rl.Texture2D
}

func NewFireball(spawnDest Destination, dest Destination) {
	f := new(Fireball)

	f.R = 100

	f.image = *rl.LoadImage("static/imgs/spells/fireball.png")
	rl.ImageResize(&f.image, int32(f.R) * 2, int32(f.R) * 2)
	f.texture = rl.LoadTextureFromImage(&f.image)

	xLine := dest.X - spawnDest.X
	yLine := dest.Y - spawnDest.Y
	c := float32(math.Sqrt(math.Pow(float64(xLine), 2) + math.Pow(float64(yLine), 2)))

	f.Y  = spawnDest.Y
	f.X  = spawnDest.X
	f.SX = xLine / c
	f.SY = yLine / c
	
	Projectiles = append(Projectiles, f)
}

func (f *Fireball) Move() {
	f.X += f.SX * speed
	f.Y += f.SY * speed
}

func (f *Fireball) Render() {
	fmt.Printf(" (%v, %v) ", f.X, f.Y)
	rl.ImageResize(&f.image, int32(f.R * 2), int32(f.R * 2))
	rl.DrawTexture(f.texture, int32(f.X), int32(f.Y), rl.White)
}