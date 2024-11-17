package main

import (
	"bufio"
	"fmt"
	"os"
)

type stack[T any] []T

func (s *stack[T]) Push(value T) {
	*s = append(*s, value)
}

func (s *stack[T]) Pop() (T, bool) {
	if len(*s) == 0 {
		var zero T
		return zero, false
	}
	value := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return value, true
}

func correctBracketSequence(s string) string {
	var st stack[rune]
	if len(s) == 1 {
		return "no"
	}
	for _, v := range s {
		switch v {
		case '(', '{', '[':
			st.Push(v)
		case ')':
			value, ok := st.Pop()
			if !ok || value != '(' {
				return "no"
			}
		case '}':
			value, ok := st.Pop()
			if !ok || value != '{' {
				return "no"
			}
		case ']':
			value, ok := st.Pop()
			if !ok || value != '[' {
				return "no"
			}
		}
	}
	if _, ok := st.Pop(); ok {
		return "no"
	}
	return "yes"
}

func main() {
	in, _ := os.Open("input.txt")
	defer func() { _ = in.Close() }()
	out, _ := os.Create("output.txt")
	defer func() { _ = out.Close() }()
	scanner := bufio.NewScanner(in)
	scanner.Buffer(make([]byte, 150000), 150000)
	writer := bufio.NewWriter(out)
	defer func() { _ = writer.Flush() }()
	for scanner.Scan() {
		s := scanner.Text()
		_, _ = fmt.Fprintln(writer, correctBracketSequence(s))
	}
}
