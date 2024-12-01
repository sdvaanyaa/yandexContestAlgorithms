package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var sizes map[int]int

func calculationSize(node int, list map[int][]int, visited map[int]bool) int {
	size := 1
	visited[node] = true

	for _, neighbor := range list[node] {
		if !visited[neighbor] {
			size += calculationSize(neighbor, list, visited)
		}
	}
	sizes[node] = size
	return size
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	nStr, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(nStr))

	list := make(map[int][]int, n)

	for i := 0; i < n-1; i++ {
		line, _ := reader.ReadString('\n')
		parts := strings.Fields(line)
		u, _ := strconv.Atoi(parts[0])
		v, _ := strconv.Atoi(parts[1])

		list[u] = append(list[u], v)
		list[v] = append(list[v], u)
	}

	visited := make(map[int]bool, n)
	sizes = make(map[int]int)

	calculationSize(1, list, visited)
	var nodes []int
	for node := range sizes {
		nodes = append(nodes, node)
	}
	sort.Ints(nodes)
	for _, node := range nodes {
		fmt.Printf("%d ", sizes[node])
	}
	fmt.Println()
}
