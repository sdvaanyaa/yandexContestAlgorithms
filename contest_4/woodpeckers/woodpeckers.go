package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calculateFreeTops(n, freeInput, k int) int {
	result := 1
	n += 2
	for i := 0; i < freeInput; i++ {
		result = (result * (n + i)) % k
	}
	return result
}

func validateGraph(graph *[][]int, input [][]int, freeTops *int) bool {
	visited := make([]bool, len(input))
	visited[0] = true

	for pointer := 1; pointer < len(input); pointer++ {
		if visited[pointer] {
			continue
		}

		pathTree := []int{}
		dfsQueue := newQueue()
		dfsQueue.push(dequeItem{curTop: pointer, prevTop: 0})

		for !dfsQueue.isEmpty() {
			item := dfsQueue.pop()
			curTop, prevTop := item.curTop, item.prevTop

			pathTree = append(pathTree, curTop)
			if visited[curTop] {
				return false
			}
			visited[curTop] = true

			subtreeCount := 0
			for _, top := range input[curTop] {
				if len(input[top]) > 1 {
					subtreeCount++
				}
				if top != prevTop {
					dfsQueue.push(dequeItem{curTop: top, prevTop: curTop})
				}
			}
			if subtreeCount > 2 {
				return false
			}
		}

		if len(pathTree) == 1 {
			*freeTops++
		} else {
			*graph = append(*graph, pathTree)
		}
	}

	return true
}

func factorial(val, k int) int {
	if val == 0 {
		return 1
	}
	result := 1
	for i := 1; i <= val; i++ {
		result = (result * i) % k
	}
	return result
}

func calculateAnswerToTree(graph [][]int, input [][]int, k int) int {
	result := 1
	for _, tree := range graph {
		localResult := 1
		dfsQueue := newQueue()
		dfsQueue.push(dequeItem{curTop: tree[0], prevTop: 0})

		processedFlag := false
		for !dfsQueue.isEmpty() {
			item := dfsQueue.pop()
			curTop, prevTop := item.curTop, item.prevTop

			subtreeCount, singleInput := 0, 0
			if len(input[curTop]) == 1 && len(tree) != 2 {
				dfsQueue.push(dequeItem{curTop: input[curTop][0], prevTop: 0})
				continue
			}
			for _, top := range input[curTop] {
				if len(input[top]) > 1 {
					subtreeCount++
				}
				if len(input[top]) == 1 {
					singleInput++
				}
				if top != prevTop && len(input[top]) > 1 {
					dfsQueue.push(dequeItem{curTop: top, prevTop: curTop})
				}
			}
			localResult = (localResult * factorial(singleInput, k)) % k
			if subtreeCount >= 1 && !processedFlag {
				localResult = (localResult * 2) % k
				processedFlag = true
			}
		}
		localResult = (localResult * 2) % k
		result = (result * localResult) % k
	}
	return result
}

func calculateFinalAnswer(graph [][]int, input [][]int, k int, n, freeTops int) {
	result := calculateAnswerToTree(graph, input, k)
	T := len(graph)
	result = (result * factorial(T, k)) % k
	result = (result * calculateFreeTops(n-freeTops, freeTops, k)) % k
	fmt.Println(result)
}

type dequeItem struct {
	curTop  int
	prevTop int
}

type queue struct {
	items []dequeItem
}

func newQueue() *queue {
	return &queue{items: []dequeItem{}}
}

func (q *queue) push(item dequeItem) {
	q.items = append(q.items, item)
}

func (q *queue) pop() dequeItem {
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func (q *queue) isEmpty() bool {
	return len(q.items) == 0
}

func main() {
	inFile, _ := os.Open("input.txt")
	defer inFile.Close()
	outFile, _ := os.Create("output.txt")
	defer outFile.Close()

	scanner := bufio.NewScanner(inFile)
	writer := bufio.NewWriter(outFile)
	defer writer.Flush()

	scanner.Scan()
	firstLine := strings.Split(scanner.Text(), " ")
	n, _ := strconv.Atoi(firstLine[0])
	m, _ := strconv.Atoi(firstLine[1])
	k, _ := strconv.Atoi(firstLine[2])

	input := make([][]int, n+1)
	var graph [][]int
	freeTops := 0

	for i := 0; i < m; i++ {
		scanner.Scan()
		edge := strings.Split(scanner.Text(), " ")
		top1, _ := strconv.Atoi(edge[0])
		top2, _ := strconv.Atoi(edge[1])
		input[top1] = append(input[top1], top2)
		input[top2] = append(input[top2], top1)
	}

	if validateGraph(&graph, input, &freeTops) {
		calculateFinalAnswer(graph, input, k, n, freeTops)
	} else {
		fmt.Println(0)
	}
}
