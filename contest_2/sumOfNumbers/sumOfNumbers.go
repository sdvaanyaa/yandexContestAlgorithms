package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func twoPointers(slice []int, k int) int {
	count := 0
	sums := 0
	left := 0

	for right := 0; right < len(slice); right++ {
		sums += slice[right]

		for sums > k && left <= right {
			sums -= slice[left]
			left++
		}
		if sums == k {
			count++
		}
	}
	return count
}

func main() {
	inputFile, _ := os.Open("input.txt")
	defer func() { _ = inputFile.Close() }()

	outputFile, _ := os.Create("output.txt")
	defer func() { _ = outputFile.Close() }()

	scanner := bufio.NewScanner(inputFile)
	scanner.Buffer(make([]byte, 1024*1024), 1024*1024)
	writer := bufio.NewWriter(outputFile)
	defer func() { _ = writer.Flush() }()

	for scanner.Scan() {
		NandK := strings.Split(scanner.Text(), " ")
		n, _ := strconv.Atoi(NandK[0])
		k, _ := strconv.Atoi(NandK[1])
		scanner.Scan()
		nums := strings.Split(scanner.Text(), " ")
		slice := make([]int, n)
		for i := 0; i < n; i++ {
			slice[i], _ = strconv.Atoi(nums[i])
		}

		_, _ = fmt.Fprintln(writer, twoPointers(slice, k))
	}
}
