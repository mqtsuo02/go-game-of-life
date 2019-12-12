package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	width    = 60
	height   = 20
	lifespan = 100
	duration = 100
	clear    = "\033[2J"
	head     = "\033[1;1H"
)

type cells [width + 2][height + 2]int

var cs, ncs cells

func main() {
	initialize()
	for i := 0; i < lifespan; i++ {
		render()
		update()
		time.Sleep(100 * time.Millisecond)
	}
	end()
}

func initialize() {
	fmt.Print(clear)
	rand.Seed(time.Now().UnixNano())
	for y := 1; y < height+1; y++ {
		for x := 1; x < width+1; x++ {
			cs[x][y] = rand.Intn(2)
		}
	}
}

func render() {
	var screen string
	for y := 0; y < height+2; y++ {
		for x := 0; x < width+2; x++ {
			c := " "
			if cs[x][y] == 1 {
				c = "▉"
			}
			screen += c
		}
		screen += "\n"
	}
	fmt.Print(head)
	fmt.Print(screen)
}

func update() {
	for y := 1; y < height+1; y++ {
		for x := 1; x < width+1; x++ {
			ncs[x][y] = 0
			cnt := cs[x-1][y-1] + cs[x][y-1] + cs[x+1][y-1] + cs[x-1][y] + cs[x+1][y] + cs[x-1][y+1] + cs[x][y+1] + cs[x+1][y+1]
			if cnt == 2 && cs[x][y] == 1 || cnt == 3 {
				ncs[x][y] = 1
			}
		}
	}
	cs = ncs
}

func end() {
	fmt.Print("Press any key to end. ")
	bufio.NewScanner(os.Stdin).Scan()
	fmt.Print(clear)
	fmt.Print(head)
}
