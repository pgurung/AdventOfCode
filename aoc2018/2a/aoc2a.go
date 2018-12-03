package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	ids := read("./input.txt")

	twice := 0
	thrice := 0

	for _, s := range ids {
		u := uniqueChars(strings.Split(s, ""))

		t := false
		th := false

		for _, ch := range u {
			if strings.Count(s, ch) == 2 && !t {
				twice++
				t = true
			}
			if strings.Count(s, ch) == 3 && !th {
				thrice++
				th = true
			}
		}
	}

	fmt.Println(twice * thrice)

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

func uniqueChars(s []string) []string {
	seen := make(map[string]struct{}, len(s))
	j := 0
	for _, v := range s {
		if _, ok := seen[v]; ok {
			continue
		}
		seen[v] = struct{}{}
		s[j] = v
		j++
	}
	return s[:j]
}
