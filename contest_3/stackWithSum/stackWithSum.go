package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var stack []int

func processing(cmd rune, num int, writer io.Writer) {
	switch cmd {
	case '+':
		if len(stack) > 0 {
			stack = append(stack, num+stack[len(stack)-1])
		} else {
			stack = append(stack, num)
		}
	case '-':
		if len(stack) >= 2 {
			_, _ = fmt.Fprintln(writer, stack[len(stack)-1]-stack[len(stack)-2])
		} else {
			_, _ = fmt.Fprintln(writer, stack[len(stack)-1])
		}
		stack = stack[:len(stack)-1]
	case '?':
		if len(stack) == num {
			_, _ = fmt.Fprintln(writer, stack[len(stack)-1])
		} else {
			_, _ = fmt.Fprintln(writer, stack[len(stack)-1]-stack[len(stack)-num-1])
		}
	}
}

func main() {
	in, _ := os.Open("input.txt")
	defer func() { _ = in.Close() }()
	out, _ := os.Create("output.txt")
	defer func() { _ = out.Close() }()
	scanner := bufio.NewScanner(in)
	writer := bufio.NewWriter(out)
	defer func() { _ = writer.Flush() }()

	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		for i := 0; i < n; i++ {
			scanner.Scan()
			line := scanner.Text()
			cmd := rune(line[0])
			var num int
			if len(line) > 1 {
				num, _ = strconv.Atoi(line[1:])
			}
			processing(cmd, num, writer)
		}
	}
}
