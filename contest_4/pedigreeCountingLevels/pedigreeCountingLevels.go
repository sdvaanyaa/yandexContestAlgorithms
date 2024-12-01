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

func calculationTheHeight(node *treeNode, level int, heightMap map[string]int) {
	heightMap[node.name] = level
	for _, child := range node.children {
		calculationTheHeight(child, level+1, heightMap)
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	pairs := make([]string, n-1)
	scanner := bufio.NewScanner(os.Stdin)

	names := make(map[string]struct{})
	for i := 0; i < n-1; i++ {
		scanner.Scan()
		pairs[i] = scanner.Text()
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

	heightMap := make(map[string]int)
	calculationTheHeight(root, 0, heightMap)

	file, _ := os.Create("output.txt")
	defer file.Close()

	for _, name := range nameList {
		fmt.Fprintf(file, "%s %d\n", name, heightMap[name])
	}
}
