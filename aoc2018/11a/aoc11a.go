package main

import (
	"fmt"
	"strconv"
)

const (
	maxX   = 300
	maxY   = 300
	serial = 4151
)

type cell struct {
	x     int
	y     int
	power int
}

func main() {

	powerGrid := createGrid()
	var maxPower cell
	max := 0

	for _, v := range powerGrid {
		if v.power > max {
			max = v.power
			maxPower = v
		}
	}

	fmt.Println(maxPower, max)

}

func createGrid() []cell {
	var powerGrid []cell

	for i := 1; i < 301; i++ {
		for j := 1; j < 301; j++ {
			gridPower := powerThree(i, j)
			powerGrid = append(powerGrid, cell{x: i, y: j, power: gridPower})

		}
	}

	return powerGrid
}

func calcPower(x int, y int) int {
	rackID := x + 10
	pow := strconv.Itoa(((rackID * y) + serial) * rackID)
	hundreds := string(pow[len(pow)-3])
	third, _ := strconv.Atoi(hundreds)
	power := third - 5

	return power

}

func powerThree(x int, y int) int {
	gridPower := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if x+i < 301 && y+j < 301 {
				gridPower = gridPower + calcPower(x+i, y+j)
			} else {
				gridPower = 0
				return gridPower
			}
		}
	}

	return gridPower
}
