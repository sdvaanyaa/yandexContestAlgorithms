package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func twoPointers(str string, c, n int) int {
	var l, r, bestl, bestr int
	var rude, countA, countB int

	for r < n {
		if str[r] == 'a' {
			countA++
		} else if str[r] == 'b' {
			rude += countA
			countB++
		}
		r++

		for rude > c && l < r {
			if str[l] == 'a' {
				countA--
				rude -= countB
			} else if str[l] == 'b' {
				countB--
			}
			l++
		}

		if r-l > bestr-bestl {
			bestl, bestr = l, r
		}
	}

	return bestr - bestl
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
		firstNums := strings.Split(scanner.Text(), " ")
		n, _ := strconv.Atoi(firstNums[0])
		c, _ := strconv.Atoi(firstNums[1])
		scanner.Scan()
		str := scanner.Text()
		_, _ = fmt.Fprintln(writer, twoPointers(str, c, n))
	}
}
