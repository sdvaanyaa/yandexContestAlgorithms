package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func dfs(node int, subordinates [][]int, coins []int) (int, int) {
	totalCoins := 0
	totalDescendants := 0

	for _, child := range subordinates[node] {
		childCoins, childDescendants := dfs(child, subordinates, coins)
		totalCoins += childCoins
		totalDescendants += childDescendants
	}

	totalCoins += totalDescendants + 1

	totalDescendants += 1

	coins[node] = totalCoins

	return totalCoins, totalDescendants
}

func calculateCoins(n int, a []int) []int {
	subordinates := make([][]int, n+1)
	for i := 2; i <= n; i++ {
		boss := a[i-2]
		subordinates[boss] = append(subordinates[boss], i)
	}

	coins := make([]int, n+1)

	dfs(1, subordinates, coins)

	return coins[1:]
}

func main() {
	in, _ := os.Open("input.txt")
	defer in.Close()
	out, _ := os.Create("output.txt")
	defer out.Close()
	reader := bufio.NewReader(in)
	writer := bufio.NewWriter(out)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(line))

	line, _ = reader.ReadString('\n')
	parts := strings.Fields(strings.TrimSpace(line))
	parents := make([]int, n-1)
	for i, p := range parts {
		parents[i], _ = strconv.Atoi(p)
	}

	result := calculateCoins(n, parents)
	for _, coin := range result {
		fmt.Fprintf(writer, "%d ", coin)
	}
}
