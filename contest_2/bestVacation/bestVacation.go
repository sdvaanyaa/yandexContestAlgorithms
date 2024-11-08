package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func maxim(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func twoPointers(slice []int, k int) int {
	sort.Ints(slice)
	maxDist := 0
	nowDist := 0
	last := 0
	for first := 0; first < len(slice); first++ {
		for last < len(slice) && abs(slice[last]-slice[first]) <= k {
			last++
			nowDist = last - first
		}
		maxDist = maxim(maxDist, nowDist)
		nowDist = 0
	}
	return maxDist
}

func main() {
	in, _ := os.Open("input.txt")
	defer func() { _ = in.Close() }()
	out, _ := os.Create("output.txt")
	defer func() { _ = out.Close() }()
	scanner := bufio.NewScanner(in)
	scanner.Buffer(make([]byte, 6*1024*1024), 6*1024*1024)
	writer := bufio.NewWriter(out)
	defer func() { _ = writer.Flush() }()

	for scanner.Scan() {
		nK := strings.Split(scanner.Text(), " ")
		n, _ := strconv.Atoi(nK[0])
		k, _ := strconv.Atoi(nK[1])
		scanner.Scan()
		nums := strings.Split(scanner.Text(), " ")
		slice := make([]int, n)
		for i := 0; i < n; i++ {
			slice[i], _ = strconv.Atoi(nums[i])
		}

		_, _ = fmt.Fprintln(writer, twoPointers(slice, k))
	}
}
