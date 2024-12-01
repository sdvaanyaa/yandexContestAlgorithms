package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type treeNode struct {
	name     string
	children []*treeNode
}

func buildTree(pairs []string) *treeNode {
	parentToChildren := make(map[string][]string)
	allChildren := make(map[string]bool)

	for _, pair := range pairs {
		parts := strings.Fields(pair)
		child, parent := parts[0], parts[1]
		parentToChildren[parent] = append(parentToChildren[parent], child)
		allChildren[child] = true
	}

	var root string
	for parent := range parentToChildren {
		if !allChildren[parent] {
			root = parent
			break
		}
	}

	var build func(name string) *treeNode
	build = func(name string) *treeNode {
		node := &treeNode{name: name}
		for _, child := range parentToChildren[name] {
			node.children = append(node.children, build(child))
		}
		return node
	}

	return build(root)
}

func calculateChildrenCount(node *treeNode, childrenCount map[string]int) int {
	count := 0
	for _, child := range node.children {
		count++
		count += calculateChildrenCount(child, childrenCount)
	}
	childrenCount[node.name] = count
	return count
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	reader := bufio.NewReader(file)

	var n int
	line, _ := reader.ReadString('\n')
	fmt.Sscanf(line, "%d", &n)

	pairs := make([]string, n-1)
	names := make(map[string]struct{})
	for i := 0; i < n-1; i++ {
		s, _ := reader.ReadString('\n')
		pairs[i] = strings.TrimSpace(s)
		parts := strings.Fields(pairs[i])
		names[parts[0]] = struct{}{}
		names[parts[1]] = struct{}{}
	}

	var nameList []string
	for name := range names {
		nameList = append(nameList, name)
	}
	sort.Strings(nameList)

	root := buildTree(pairs)

	childrenCount := make(map[string]int)
	calculateChildrenCount(root, childrenCount)

	outputFile, _ := os.Create("output.txt")
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)

	for _, name := range nameList {
		fmt.Fprintln(writer, name, childrenCount[name])
	}

	writer.Flush()
}
