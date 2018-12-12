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
	size  int
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

	fmt.Println(maxPower)

}

func createGrid() []cell {
	var powerGrid []cell

	for i := 1; i < 301; i++ {
		for j := 1; j < 301; j++ {
			power, size := maxGridPower(i, j)
			powerGrid = append(powerGrid, cell{x: i, y: j, power: power, size: size})
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

func findGridPower(x int, y int, s int) int {
	gridPower := 0
	for i := 0; i < s; i++ {
		for j := 0; j < s; j++ {
			if x+i < 301 && y+j < 301 {
				gridPower = gridPower + calcPower(x+i, y+j)
			}
		}
	}

	return gridPower
}

func maxGridPower(x int, y int) (int, int) {
	power, size := 0, 0

	// maxSize := findMax(300-x, 300-y)

	for i := 1; i < 50; i++ {
		if findGridPower(x, y, i) > power {
			power = findGridPower(x, y, i)
			size = i
		}

	}

	return power, size

}

func findMax(a, b int) int {
	if a > b {
		return a
	}

	return b
}
