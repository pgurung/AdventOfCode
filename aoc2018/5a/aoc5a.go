package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {

	atz := "abcdefghijklmnopqrstuvwxyz"

	lower := strings.Split(atz, "")
	upper := strings.Split(strings.ToUpper(atz), "")

	dict := createDict(lower, upper)

	in, _ := ioutil.ReadFile("./input.txt")

	inStr := string(in)

	polymers := strings.Split(inStr, "")

	units := filterpolymers(dict, polymers)

	fmt.Println(units)
}

func createDict(sl []string, su []string) map[string]string {

	dict := make(map[string]string)

	for i := 0; i < 26; i++ {
		dict[sl[i]] = su[i]
		dict[su[i]] = sl[i]
	}

	return dict

}

func filterpolymers(d map[string]string, p []string) int {
	po := p

	var stack []string
	changes := 0

	for {
		for _, s := range po {
			if len(stack) == 0 {
				stack = append(stack, s)
				continue
			}

			l := len(stack)
			v, ok := d[s]
			if !ok {
				continue
			}

			if ok && v == stack[l-1] {
				stack = stack[:l-1]
				changes++
				continue
			}
			stack = append(stack, s)
		}

		if changes == 0 {
			break
		}

		po = stack[0:]

		stack = stack[:0]
		changes = 0

	}

	return len(stack)

}
