package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func twoPointers(slice []int) string {
	n := len(slice)
	mid := (n + 1) / 2
	i, j := mid-1, mid
	sort.Ints(slice)
	var builder strings.Builder

	if n%2 != 0 {
		builder.WriteString(strconv.Itoa(slice[i]))
		builder.WriteByte(' ')
		i--
	}

	for i >= 0 || j < n {
		if i >= 0 {
			builder.WriteString(strconv.Itoa(slice[i]))
			builder.WriteByte(' ')
			i--
		}
		if j < n {
			builder.WriteString(strconv.Itoa(slice[j]))
			j++
			if j != n {
				builder.WriteByte(' ')
			}
		}
	}
	return builder.String()
}

func main() {
	in, _ := os.Open("input.txt")
	defer func() { _ = in.Close() }()
	out, _ := os.Create("output.txt")
	defer func() { _ = out.Close() }()
	scanner := bufio.NewScanner(in)
	scanner.Buffer(make([]byte, 1024*1024*2), 1024*1024*2)
	writer := bufio.NewWriter(out)
	defer func() { _ = writer.Flush() }()

	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		nums := strings.Split(scanner.Text(), " ")
		slice := make([]int, n)
		for i := range nums {
			slice[i], _ = strconv.Atoi(nums[i])
		}

		_, _ = fmt.Fprintln(writer, twoPointers(slice))
	}
}
