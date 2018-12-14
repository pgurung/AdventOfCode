package main

import (
	"fmt"
	"strconv"
	"time"
)

const (
	serial = 4151 //Puzzle Input
)

type cell struct {
	x     int
	y     int
	power int
	size  int
}
type point struct {
	x int
	y int
	s int
}

type pointPow struct {
	p int
	s int
}

func main() {

	start := time.Now()

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
	fmt.Println("Total time:", time.Since(start))

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

func findGridPower(input <-chan point, output chan<- pointPow) {
	for p := range input {
		gridPower := 0
		for i := 0; i < p.s; i++ {
			for j := 0; j < p.s; j++ {
				if p.x+i < 301 && p.y+j < 301 {
					gridPower = gridPower + calcPower(p.x+i, p.y+j)
				}
			}
		}
		output <- pointPow{p: gridPower, s: p.s}
	}
}

func maxGridPower(x int, y int) (int, int) {
	power, size := 0, 0
	// maxSize := findMax(300-x, 300-y) calculates all the possible values
	//however for the given input of 4151, maxSize can be set to 15 and you'll still get the correct answer
	//execution time of 1 hr vs 2-3 seconds!!
	// without the use of channel this would take multiple hours and a few minutes respectively
	maxSize := 15

	input := make(chan point, maxSize)
	output := make(chan pointPow, maxSize)

	go findGridPower(input, output)
	go findGridPower(input, output)
	go findGridPower(input, output)
	go findGridPower(input, output)

	for i := 1; i < maxSize; i++ {
		input <- point{x: x, y: y, s: i}

	}
	close(input)

	for i := 0; i < maxSize-1; i++ {
		ps := <-output
		if ps.p > power {
			power = ps.p
			size = ps.s
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
