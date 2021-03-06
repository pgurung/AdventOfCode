package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type coords struct {
	x int
	y int
}

func main() {
	claims := read("./input.txt")

	// var diff []coords
	// var common []coords

	mc := make(map[coords]bool)
	t := 0

	re := regexp.MustCompile(`(?m)\D+`)

	for _, e := range claims {
		claim := strings.Split(strings.TrimSpace(re.ReplaceAllString(e, " ")), " ")

		le, _ := strconv.Atoi(claim[1])
		te, _ := strconv.Atoi(claim[2])
		w, _ := strconv.Atoi(claim[3])
		h, _ := strconv.Atoi(claim[4])

		for i := te + 1; i <= te+h; i++ {
			for j := le + 1; j <= le+w; j++ {
				c := coords{j, i}
				if v, ok := mc[c]; ok && !v {
					t++
					mc[c] = true
				}
				if v, ok := mc[c]; ok && v {
					continue
				}

				mc[c] = false
			}
		}

	}

	fmt.Println(t)

}

func read(s string) []string {
	file, _ := os.Open(s)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var r []string

	for scanner.Scan() {
		r = append(r, scanner.Text())
	}

	return r

}
