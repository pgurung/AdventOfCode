package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./1aInput.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	frequency := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())

		frequency = frequency + num
	}

	fmt.Println(frequency)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
