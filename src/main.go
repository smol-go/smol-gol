package main

import (
	"image/color"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	scale  = 2
	width  = 300
	height = 300
)

var (
	blue   = color.RGBA{69, 145, 196, 255}
	yellow = color.RGBA{255, 230, 120, 255}
	grid   [width][height]uint8
	buffer [width][height]uint8
	count  int
)

type Game struct{}

func (g *Game) Update() error {
	count++
	if count == 20 {
		for x := 1; x < width-1; x++ {
			for y := 1; y < height-1; y++ {
				buffer[x][y] = 0
				n := grid[x-1][y-1] + grid[x-1][y+0] + grid[x-1][y+1] +
					grid[x+0][y-1] + grid[x+0][y+1] +
					grid[x+1][y-1] + grid[x+1][y+0] + grid[x+1][y+1]
				if grid[x][y] == 0 && n == 3 {
					buffer[x][y] = 1
				} else if n < 2 || n > 3 {
					buffer[x][y] = 0
				} else {
					buffer[x][y] = grid[x][y]
				}
			}
		}
		grid, buffer = buffer, grid
		count = 0
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(blue)
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if grid[x][y] == 1 {
				for i := 0; i < scale; i++ {
					for j := 0; j < scale; j++ {
						screen.Set(x*scale+i, y*scale+j, yellow)
					}
				}
			}
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return width * scale, height * scale
}

func main() {
	for x := 1; x < width-1; x++ {
		for y := 1; y < height-1; y++ {
			if rand.Float32() < 0.5 {
				grid[x][y] = 1
			}
		}
	}

	ebiten.SetWindowSize(width*scale, height*scale)
	ebiten.SetWindowTitle("Game of Life")

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
