package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func twoPointers(slice []int, r int) int {
	monuments := 0
	last := 0
	for first := 0; first < len(slice); first++ {
		for last < len(slice) && slice[last]-slice[first] <= r {
			last++
		}
		monuments += len(slice) - last
	}
	return monuments
}

func main() {
	in, _ := os.Open("input.txt")
	defer func() { _ = in.Close() }()
	out, _ := os.Create("output.txt")
	defer func() { _ = out.Close() }()
	scanner := bufio.NewScanner(in)
	scanner.Buffer(make([]byte, 4*1024*1024), 4*1024*1024)
	writer := bufio.NewWriter(out)
	defer func() { _ = writer.Flush() }()

	for scanner.Scan() {
		nR := strings.Split(scanner.Text(), " ")
		n, _ := strconv.Atoi(nR[0])
		r, _ := strconv.Atoi(nR[1])
		scanner.Scan()
		nums := strings.Split(scanner.Text(), " ")
		slice := make([]int, n)
		for i := 0; i < n; i++ {
			slice[i], _ = strconv.Atoi(nums[i])
		}

		_, _ = fmt.Fprintln(writer, twoPointers(slice, r))
	}
}
