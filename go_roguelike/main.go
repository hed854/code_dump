package main

import "fmt"

var (
	width  = 14
	height = 8
)

type Tile string

const (
	PLA = "@"
	WAL = "#"
	FLR = "."
	DOR = "+"
	GRA = ","
)

type Grid struct {
	width   int
	height  int
	content [][]Tile
}

func (g *Grid) draw() {
	for i := range g.content {
		for j := range g.content[i] {
			fmt.Printf("%s", g.content[i][j])
		}
		fmt.Printf("\n")
	}
}

type Player struct {
	x    int
	y    int
	repr Tile
}

func NewPlayer(x, y int) *Player {
	// Fake constructor pattern
	p := new(Player)
	p.x = x
	p.y = y
	p.repr = PLA
	return p
}

func (p *Player) draw(g *Grid) {
	// no need to return anything
	// we add the player to the game screen
	g.content[p.x][p.y] = p.repr
	return
}

func update(g *Grid) {
}

func main() {
	input := ""
	// Basic map
	game := Grid{height, width, [][]Tile{
		{GRA, GRA, GRA, GRA, GRA, GRA, GRA, GRA, GRA, GRA, GRA, GRA, GRA, GRA},
		{GRA, WAL, WAL, WAL, WAL, WAL, GRA, GRA, GRA, GRA, GRA, GRA, GRA, GRA},
		{GRA, WAL, FLR, FLR, FLR, WAL, GRA, GRA, GRA, GRA, GRA, GRA, GRA, GRA},
		{GRA, WAL, FLR, FLR, FLR, WAL, GRA, GRA, GRA, GRA, GRA, GRA, GRA, GRA},
		{GRA, WAL, WAL, DOR, WAL, WAL, GRA, GRA, GRA, GRA, GRA, GRA, GRA, GRA},
		{GRA, GRA, GRA, GRA, GRA, GRA, GRA, GRA, GRA, GRA, GRA, GRA, GRA, GRA},
		{GRA, GRA, GRA, GRA, GRA, GRA, GRA, GRA, GRA, GRA, GRA, GRA, GRA, GRA},
		{GRA, GRA, GRA, GRA, GRA, GRA, GRA, GRA, GRA, GRA, GRA, GRA, GRA, GRA},
	}}
	player := NewPlayer(6, 6)

	// Send the game screen by pointer
	player.draw(&game)
	game.draw()

	fmt.Printf("> ")
	fmt.Scanln(&input)

	// need console clear
}
