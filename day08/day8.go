package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Node struct {
	children []Node
	metadata []int
}

func (n *Node) metadataSum() int {
	total := 0
	for _, m := range n.metadata {
		total += m
	}
	for _, c := range n.children {
		total += c.metadataSum()
	}
	return total
}

func (n *Node) value() int {
	var total int
	if n.children == nil {
		for _, m := range n.metadata {
			total += m
		}
	} else {
		for _, m := range n.metadata {
			if m-1 < len(n.children) {
				total += n.children[m-1].value()
			}
		}
	}
	return total
}

func parseNode(nums []int) (Node, int) {
	numChildren := nums[0]
	numMetadata := nums[1]
	i := 2
	node := Node{}
	for c := 0; c < numChildren; c++ {
		child, numConsumed := parseNode(nums[i:])
		i += numConsumed
		node.children = append(node.children, child)
	}
	node.metadata = nums[i : i+numMetadata]
	return node, i + numMetadata
}

func part1(nums []int) int {
	tree, _ := parseNode(nums)
	return tree.metadataSum()
}

func part2(nums []int) int {
	tree, _ := parseNode(nums)
	return tree.value()
}

func main() {
	inputBytes, _ := ioutil.ReadFile("input/day08.txt")
	inputString := strings.TrimSpace(string(inputBytes))
	var input []int
	for _, item := range strings.Split(inputString, " ") {
		num, _ := strconv.Atoi(item)
		input = append(input, num)
	}

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
