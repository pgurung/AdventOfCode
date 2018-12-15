package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const (
	input = "293801"
)

type elf struct {
	score int
	index int
}

func main() {
	elf1 := elf{score: 3, index: 0}
	elf2 := elf{score: 7, index: 1}

	scoreboard := []string{"3", "7"}

	for len(scoreboard) < 40000000 {
		score := strings.Split(strconv.Itoa(elf1.score+elf2.score), "")
		scoreboard = append(scoreboard, score...)

		elf1 = move(elf1, scoreboard)
		elf2 = move(elf2, scoreboard)
	}

	fmt.Println(strings.Index(strings.Join(scoreboard, ""), input))

}

func move(e elf, s []string) elf {
	moves := e.score + 1

	var newElf elf

	if e.index+moves < len(s) {
		index := e.index + moves
		score, _ := strconv.Atoi(s[index])
		newElf.score = score
		newElf.index = index
	} else {
		l := int(math.Floor(float64(e.index+moves) / float64(len(s))))
		index := e.index + moves - l*len(s)
		score, _ := strconv.Atoi(s[index])
		newElf.score = score
		newElf.index = index
	}

	return newElf
}
