package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const mod = 1000000007

func prefixSums(slice []int) int {
	reverseSlice := make([]int, len(slice))
	for i := range slice {
		reverseSlice[i] = slice[len(slice)-1-i]
	}

	prefix := make([]int, len(reverseSlice)-1)
	for i := 1; i < len(reverseSlice)-1; i++ {
		prefix[i] = (prefix[i-1] + reverseSlice[i-1]) % mod
	}

	for i := 1; i < len(reverseSlice)-1; i++ {
		prefix[i] = prefix[i] * reverseSlice[i] % mod
	}

	for i := 1; i < len(slice)-1; i++ {
		prefix[i] = (prefix[i] + prefix[i-1]) % mod
	}

	reversePrefix := make([]int, len(prefix))
	for i := range prefix {
		reversePrefix[i] = prefix[len(prefix)-1-i]
	}

	result := 0
	for i := 0; i < len(slice)-2; i++ {
		reversePrefix[i] = reversePrefix[i] * slice[i] % mod
		result += reversePrefix[i] % mod
	}
	return result % mod
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
		nums := strings.Split(scanner.Text(), " ")
		slice := make([]int, n)
		for i := range slice {
			slice[i], _ = strconv.Atoi(nums[i])
		}
		_, _ = fmt.Fprintln(writer, prefixSums(slice))
	}
}
