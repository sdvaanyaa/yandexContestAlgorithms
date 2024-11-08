package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func maximum(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func prefixSumsBuildingEqual(n int, weights []int, which string) []int {
	prefixSums := make([]int, n)
	if which == "equal" {
		for i := 1; i < n; i++ {
			if weights[i-1] == weights[i] {
				prefixSums[i] = prefixSums[i-1] + 1
			} else {
				prefixSums[i] = prefixSums[i-1]
			}
		}
	} else if which == "bigger" {
		for i := 1; i < n; i++ {
			if weights[i-1] > weights[i] {
				prefixSums[i] = prefixSums[i-1] + 1
			} else {
				prefixSums[i] = prefixSums[i-1]
			}
		}
	}
	return prefixSums
}

func fastSearchBuilding(prefixSums []int, n int) map[int]int {
	fastSearch := make(map[int]int)
	for i := n - 1; i >= 0; i-- {
		fastSearch[prefixSums[i]] = i + 1
	}
	return fastSearch
}

func optimizedSolution(n, k int, weights, start []int) string {
	prefixSumsEqual := prefixSumsBuildingEqual(n, weights, "equal")
	fastSearchEqual := fastSearchBuilding(prefixSumsEqual, n)

	prefixSumsBigger := prefixSumsBuildingEqual(n, weights, "bigger")
	fastSearchBigger := fastSearchBuilding(prefixSumsBigger, n)

	var resultIndex int
	var build strings.Builder
	for _, i := range start {
		resultIndex = -1
		bigger := prefixSumsBigger[i-1]
		if fastSearchBigger[bigger] == 1 {
			resultIndex = maximum(resultIndex, fastSearchBigger[0])
		} else {
			resultIndex = maximum(resultIndex, fastSearchBigger[bigger])
		}

		equal := prefixSumsEqual[i-1]
		if equal-k <= 0 {
			resultIndex = maximum(resultIndex, fastSearchEqual[0])
		} else {
			resultIndex = maximum(resultIndex, fastSearchEqual[equal-k])
		}
		build.WriteString(strconv.Itoa(resultIndex))
		build.WriteByte(' ')
	}
	return build.String()
}

func main() {
	in, _ := os.Open("input.txt")
	defer func() { _ = in.Close() }()
	out, _ := os.Create("output.txt")
	defer func() { _ = out.Close() }()
	scanner := bufio.NewScanner(in)
	scanner.Buffer(make([]byte, 1024*1024*10), 1024*1024*10)
	writer := bufio.NewWriter(out)
	defer func() { _ = writer.Flush() }()

	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())

		scanner.Scan()
		strWeights := strings.Split(scanner.Text(), " ")

		scanner.Scan()
		mK := strings.Split(scanner.Text(), " ")
		m, _ := strconv.Atoi(mK[0])
		k, _ := strconv.Atoi(mK[1])

		scanner.Scan()
		strStart := strings.Split(scanner.Text(), " ")

		weights := make([]int, n)
		start := make([]int, m)
		for i := 0; i < n; i++ {
			weights[i], _ = strconv.Atoi(strWeights[i])
		}

		for i := 0; i < m; i++ {
			start[i], _ = strconv.Atoi(strStart[i])
		}

		_, _ = fmt.Fprintln(writer, optimizedSolution(n, k, weights, start))
	}
}
