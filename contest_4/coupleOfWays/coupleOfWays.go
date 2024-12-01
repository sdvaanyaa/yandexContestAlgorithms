package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func getMax(values ...int) int {
	max := values[0]
	for _, v := range values[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

func findDepth(tree [][]int, startNode int) (int, int, []int) {
	stack := []struct {
		node   int
		parent int
		depth  int
	}{{startNode, -1, 0}}

	parents := make(map[int]int)
	parents[startNode] = -1

	maxDepth, deepestNode := 0, startNode

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if current.depth > maxDepth {
			maxDepth, deepestNode = current.depth, current.node
		}

		for _, neighbor := range tree[current.node] {
			if neighbor != current.parent {
				parents[neighbor] = current.node
				stack = append(stack, struct {
					node   int
					parent int
					depth  int
				}{neighbor, current.node, current.depth + 1})
			}
		}
	}

	path := []int{}
	for cur := deepestNode; cur != -1; cur = parents[cur] {
		path = append([]int{cur}, path...)
	}

	return deepestNode, maxDepth, path
}

func solve(n int, edges [][]int) int {
	if n == 2 {
		return 0
	}

	tree := make([][]int, n+1)
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		tree[a] = append(tree[a], b)
		tree[b] = append(tree[b], a)
	}

	root, _, _ := findDepth(tree, 1)
	_, _, path := findDepth(tree, root)

	if len(path) < 4 {
		return 0
	}

	visitedNodes := make(map[int]bool)
	for _, node := range path {
		visitedNodes[node] = true
	}

	removedPaths := [][]int{}
	maxPathLength := 0
	_, _, maxPathLength, _ = calculatePath(tree, visitedNodes, root, &removedPaths, &maxPathLength, true)

	result := maxPathLength * (len(path) - 1)
	return calculateResult(path, removedPaths, result)
}

func calculatePath(tree [][]int, visitedNodes map[int]bool, root int, removedPaths *[][]int, maxPathLength *int, isStart bool) (int, int, int, int) {
	stack := []struct {
		node    int
		parent  int
		start   bool
		visited bool
	}{{root, -1, isStart, false}}

	res := make(map[int][4]int)

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if !current.visited {
			stack = append(stack, struct {
				node    int
				parent  int
				start   bool
				visited bool
			}{current.node, current.parent, current.start, true})

			for _, neighbor := range tree[current.node] {
				if neighbor != current.parent {
					stack = append(stack, struct {
						node    int
						parent  int
						start   bool
						visited bool
					}{neighbor, current.node, false, false})
				}
			}
		} else {
			depths := [4]int{}
			if len(tree[current.node]) == 1 && !current.start {
				res[current.node] = [4]int{1, current.node, *maxPathLength, 0}
				tree[current.node] = nil
				continue
			}

			for _, neighbor := range tree[current.node] {
				if neighbor != current.parent {
					childRes := res[neighbor]
					if visitedNodes[childRes[1]] {
						depths[2] = childRes[0]
					} else if childRes[0] > depths[0] {
						depths[1] = depths[0]
						depths[0] = childRes[0]
					} else if childRes[0] > depths[1] {
						depths[1] = childRes[0]
					}
				}
			}
			tree[current.node] = nil

			if depths[2] == 0 {
				*maxPathLength = getMax(*maxPathLength, depths[0]+depths[1])
			} else {
				depths[3] = getMax(
					res[current.node][3],
					depths[0]+depths[1],
					depths[2]+depths[0],
					depths[2]+depths[1],
				)
				*removedPaths = append(*removedPaths, depths[:])
			}
			res[current.node] = [4]int{getMax(depths[0], getMax(depths[1], depths[2])) + 1, current.node, *maxPathLength, depths[3]}
		}
	}
	return res[root][0], res[root][1], res[root][2], res[root][3]
}

func calculateResult(path []int, removedPaths [][]int, preResult int) int {
	lmp, leftMax := 0, 0
	for i := 1; i < len(path)-2; i++ {
		lmp = getMax(
			removedPaths[len(removedPaths)-(i+1)][0],
			removedPaths[len(removedPaths)-(i+1)][1],
			lmp+1,
		)
		leftMax = getMax(
			leftMax,
			removedPaths[len(removedPaths)-(i+1)][0]+removedPaths[len(removedPaths)-(i+1)][1],
			lmp+removedPaths[len(removedPaths)-(i+1)][0],
			lmp+removedPaths[len(removedPaths)-(i+1)][1],
		)
		preResult = getMax(
			preResult,
			leftMax*removedPaths[len(removedPaths)-(i+2)][3],
		)
	}
	return preResult
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	edges := [][]int{}
	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		a, _ := strconv.Atoi(parts[0])
		b, _ := strconv.Atoi(parts[1])
		edges = append(edges, []int{a, b})
	}

	result := solve(n, edges)

	outputFile, _ := os.Create("output.txt")
	defer outputFile.Close()
	outputFile.WriteString(strconv.Itoa(result) + "\n")
}
