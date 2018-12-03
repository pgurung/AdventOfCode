package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	ids := read("./input.txt")

	for i := 0; i < len(ids); i++ {
		for j := 0; j < len(ids); j++ {
			diff, c := findDiff(ids[i], ids[j])

			if diff == 1 {
				fmt.Println(c)
				log.Fatal("Found it!")
			}
		}
	}

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

func findDiff(s1 string, s2 string) (int, string) {
	diff := 0
	c := ""

	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			diff++
		}
		if s1[i] == s2[i] {
			c = c + string(s1[i])
		}
	}

	return diff, c
}
