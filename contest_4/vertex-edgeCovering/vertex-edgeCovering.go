package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var (
	adj    [][]int
	cost   []int
	dp     [][]int
	parent []int
	marked []bool
)

func dfs(v, p int) {
	parent[v] = p
	dp[v][0] = 0
	dp[v][1] = cost[v]

	for _, to := range adj[v] {
		if to == p {
			continue
		}
		dfs(to, v)
		dp[v][0] += dp[to][1]
		dp[v][1] += min(dp[to][0], dp[to][1])
	}
}

func restore(v, p, state int) {
	if state == 1 {
		marked[v] = true
	}
	for _, to := range adj[v] {
		if to == p {
			continue
		}
		if state == 0 {
			restore(to, v, 1)
		} else {
			if dp[to][0] < dp[to][1] {
				restore(to, v, 0)
			} else {
				restore(to, v, 1)
			}
		}
	}
}

func countMarked() int {
	count := 0
	for i := 1; i < len(marked); i++ {
		if marked[i] {
			count++
		}
	}
	return count
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	inp, _ := os.Open("input.txt")
	defer inp.Close()
	out, _ := os.Create("output.txt")
	defer out.Close()

	reader := bufio.NewReader(inp)
	writer := bufio.NewWriter(out)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(line))
	if n == 1 {
		line, _ = reader.ReadString('\n')
		cost, _ := strconv.Atoi(strings.TrimSpace(line))
		writer.WriteString(strconv.Itoa(cost) + " 1\n1\n")
		return
	}

	adj = make([][]int, n+1)
	cost = make([]int, n+1)
	dp = make([][]int, n+1)
	parent = make([]int, n+1)
	marked = make([]bool, n+1)

	for i := 0; i < n-1; i++ {
		line, _ = reader.ReadString('\n')
		parts := strings.Fields(line)
		u, _ := strconv.Atoi(parts[0])
		v, _ := strconv.Atoi(parts[1])
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	line, _ = reader.ReadString('\n')
	parts := strings.Fields(line)
	for i := 1; i <= n; i++ {
		cost[i], _ = strconv.Atoi(parts[i-1])
	}

	for i := 1; i <= n; i++ {
		dp[i] = make([]int, 2)
	}

	dfs(1, -1)

	minCost := 0
	if dp[1][0] < dp[1][1] {
		minCost = dp[1][0]
		restore(1, -1, 0)
	} else {
		minCost = dp[1][1]
		restore(1, -1, 1)
	}

	writer.WriteString(strconv.Itoa(minCost) + " " + strconv.Itoa(countMarked()) + "\n")
	for i := 1; i <= n; i++ {
		if marked[i] {
			writer.WriteString(strconv.Itoa(i) + " ")
		}
	}
	writer.WriteString("\n")
}
