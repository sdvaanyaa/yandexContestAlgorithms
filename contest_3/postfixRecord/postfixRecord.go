package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func postfixRecord(slice []string) int {
	var stack []int
	for _, v := range slice {
		if n, err := strconv.Atoi(v); err == nil {
			stack = append(stack, n)
		} else {
			if len(stack) >= 2 {
				var operationResult int
				switch v {
				case "+":
					operationResult = stack[len(stack)-1] + stack[len(stack)-2]
				case "-":
					operationResult = stack[len(stack)-2] - stack[len(stack)-1]
				case "*":
					operationResult = stack[len(stack)-1] * stack[len(stack)-2]
				}
				stack = stack[:len(stack)-2]
				stack = append(stack, operationResult)
			}
		}
	}
	return stack[0]
}

func main() {
	in, _ := os.Open("input.txt")
	defer func() { _ = in.Close() }()
	out, _ := os.Create("output.txt")
	defer func() { _ = out.Close() }()
	scanner := bufio.NewScanner(in)
	scanner.Buffer(make([]byte, 1024*1024*2), 1024*1024*4)
	writer := bufio.NewWriter(out)
	defer func() { _ = writer.Flush() }()
	for scanner.Scan() {
		slice := strings.Split(scanner.Text(), " ")
		_, _ = fmt.Fprintln(writer, postfixRecord(slice))
	}
}
