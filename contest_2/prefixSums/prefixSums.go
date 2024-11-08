package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func makePrefixSums(nums []int) string {
	prefixSums := make([]int, len(nums)+1)
	for i := 1; i < len(nums)+1; i++ {
		prefixSums[i] = prefixSums[i-1] + nums[i-1]
	}
	s := ""
	for i := 1; i < len(prefixSums); i++ {
		if i == 1 {
			s += strconv.Itoa(prefixSums[i])
		} else {
			s += " "
			s += strconv.Itoa(prefixSums[i])
		}
	}
	return s
}

func main() {
	inputFile, _ := os.Open("input.txt")
	defer func() { _ = inputFile.Close() }()

	outputFile, _ := os.Create("output.txt")
	defer func() { _ = outputFile.Close() }()

	scanner := bufio.NewScanner(inputFile)
	writer := bufio.NewWriter(outputFile)
	defer func() { _ = writer.Flush() }()

	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		nums := strings.Split(scanner.Text(), " ")
		slice := make([]int, n)
		for i := 0; i < n; i++ {
			slice[i], _ = strconv.Atoi(nums[i])
		}
		_, _ = fmt.Fprintln(writer, makePrefixSums(slice))
	}

}
