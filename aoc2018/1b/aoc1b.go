package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {

	start := time.Now()

	fd := read("./input.txt")
	frequency := 0
	m := make(map[int]bool)

	for {
		for _, v := range fd {
			num, _ := strconv.Atoi(v)
			frequency += num
			if m[frequency] {
				fmt.Println(frequency)
				log.Fatal("Done in: ", time.Since(start))
			}
			m[frequency] = true
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
