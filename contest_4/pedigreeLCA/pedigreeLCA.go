package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func findNodeByName(node *treeNode, name string) *treeNode {
	if node.name == name {
		return node
	}

	for _, child := range node.children {
		found := findNodeByName(child, name)
		if found != nil {
			return found
		}
	}

	return nil
}

func calculateLCA(v, p, q *treeNode, res *[]string) int {
	k := 0

	if v == p || v == q {
		k += 1
	}

	for _, child := range v.children {
		k += calculateLCA(child, p, q, res)
	}

	if k == 2 && len(*res) == 0 {
		*res = append(*res, v.name)
	}

	return k
}

func printTree(node *treeNode, depth int) {
	fmt.Printf("%s%s\n", strings.Repeat("  ", depth), node.name)
	for _, child := range node.children {
		printTree(child, depth+1)
	}
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	pairs := make([]string, n-1)
	for i := 0; i < n-1; i++ {
		scanner.Scan()
		pairs[i] = strings.TrimSpace(scanner.Text())
	}

	root := buildTree(pairs)
	printTree(root, 0)

	var results []string
	for scanner.Scan() {
		str := scanner.Text()
		fields := strings.Fields(str)
		p := findNodeByName(root, fields[0])
		q := findNodeByName(root, fields[1])
		if p == q {
			results = append(results, fields[0])
			continue
		}
		var res []string
		calculateLCA(root, p, q, &res)
		if len(res) > 0 {
			results = append(results, res[0])
		}
	}

	outputFile, _ := os.Create("output.txt")
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)

	for _, name := range results {
		fmt.Fprintln(writer, name)
	}

	writer.Flush()
}
