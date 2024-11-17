package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type stack[T any] []T

type pairs struct {
	index int
	value int
}

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

func correctBracketSequence(prices []int) string {
	var st stack[pairs]
	result := make([]int, len(prices))
	st.Push(pairs{0, prices[0]})
	for i := 1; i < len(prices); i++ {
		for len(st) > 0 && st[len(st)-1].value > prices[i] {
			p, _ := st.Pop()
			result[p.index] = i
		}
		st.Push(pairs{i, prices[i]})
	}

	for _, p := range st {
		result[p.index] = -1
	}

	var build strings.Builder
	for _, value := range result {
		build.WriteString(strconv.Itoa(value))
		build.WriteByte(' ')
	}

	return build.String()
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
		pricesStr := strings.Split(scanner.Text(), " ")
		prices := make([]int, n)
		for i := range pricesStr {
			num, _ := strconv.Atoi(pricesStr[i])
			prices[i] = num
		}
		_, _ = fmt.Fprintln(writer, correctBracketSequence(prices))
	}
}
