package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

//Node is ...
type Node struct {
	childNodes []Node
	metadata   []int
}

func main() {
	in, err := ioutil.ReadFile("./input.txt")
	check(err)

	input := strings.Split(string(in), " ")

	var intSlice []int

	for _, s := range input {
		num, _ := strconv.Atoi(s)
		intSlice = append(intSlice, num)
	}

	tree := createTree(intSlice)

	metaSum := sumMetadata(tree)

	fmt.Println("Metadata Sum: ", metaSum)

}

func createTree(nums []int) Node {

	pos := 0
	var tree Node

	children := nums[pos]
	pos++
	metaCount := nums[pos]
	pos++

	for i := 0; i < children; i++ {
		childNode := createTree(nums[pos:])
		tree.childNodes = append(tree.childNodes, childNode)
		pos = pos + getLength(childNode)
	}

	for i := 0; i < metaCount; i++ {
		tree.metadata = append(tree.metadata, nums[pos+i])
	}
	return tree

}

func getLength(Node Node) int {
	length := 2
	for i := 0; i < len(Node.childNodes); i++ {
		length = length + getLength(Node.childNodes[i])
	}
	length = length + len(Node.metadata)
	return length
}

func sumMetadata(tree Node) int {

	sum := 0

	for _, Nodes := range tree.childNodes {
		sum += sumMetadata(Nodes)
	}

	return sumInts(tree.metadata) + sum

}

func sumInts(nums []int) int {

	sum := 0

	for _, num := range nums {
		sum = sum + num
	}

	return sum
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
