package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var priceMap map[rune]int

func setPrice(w string) {
	priceMap = make(map[rune]int)
	for i, v := range w {
		priceMap[v] = i
	}
}

func comparePrice(r1, r2 rune) bool {
	var goStack bool
	if priceMap[r1] < priceMap[r2] {
		if r1 == '(' || r1 == '[' {
			goStack = true
		}
		return goStack
	}
	return goStack
}

func stackFilling(s string) ([]rune, *strings.Builder, int) {
	var stack []rune
	countOpenBrackets := 0
	builder := &strings.Builder{}

	for _, v := range s {
		switch v {
		case '(', '[':
			stack = append(stack, v)
			builder.WriteRune(v)
			countOpenBrackets++
		case ')':
			stack = stack[:len(stack)-1]
			builder.WriteRune(v)
		case ']':
			stack = stack[:len(stack)-1]
			builder.WriteRune(v)
		}
	}
	return stack, builder, countOpenBrackets
}

func minimumPsp(w, s string, n int) string {
	stack, builder, countOpenBrackets := stackFilling(s)

	var reverseBracket rune
	var goStack bool
	var cheapOpenBracket rune

	for _, v := range w {
		if v == '(' || v == '[' {
			cheapOpenBracket = v
			break
		}
	}

	if len(stack) == 0 {
		for len(stack) < 2 && builder.Len() < n {
			stack = append(stack, cheapOpenBracket)
			builder.WriteRune(cheapOpenBracket)
			countOpenBrackets++
			if countOpenBrackets == n/2 {
				break
			}
			top := stack[len(stack)-1]
			if top == '[' {
				reverseBracket = ']'
			} else if top == '(' {
				reverseBracket = ')'
			}

			goStack = comparePrice(cheapOpenBracket, reverseBracket)
			if goStack {
				stack = append(stack, cheapOpenBracket)
				builder.WriteRune(cheapOpenBracket)
				countOpenBrackets++
			} else {
				builder.WriteRune(reverseBracket)
				stack = stack[:len(stack)-1]
			}
		}
	}

	for len(stack) > 0 {
		if countOpenBrackets == n/2 {
			break
		}
		top := stack[len(stack)-1]
		if top == '[' {
			reverseBracket = ']'
		} else if top == '(' {
			reverseBracket = ')'
		}

		goStack = comparePrice(cheapOpenBracket, reverseBracket)
		if goStack {
			stack = append(stack, cheapOpenBracket)
			builder.WriteRune(cheapOpenBracket)
			countOpenBrackets++
		} else {
			builder.WriteRune(reverseBracket)
			stack = stack[:len(stack)-1]
		}
	}

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		if top == '[' {
			reverseBracket = ']'
		} else if top == '(' {
			reverseBracket = ')'
		}
		builder.WriteRune(reverseBracket)
		stack = stack[:len(stack)-1]
	}

	for builder.Len() < n {
		stack = append(stack, cheapOpenBracket)
		builder.WriteRune(cheapOpenBracket)
		top := stack[len(stack)-1]
		if top == '[' {
			reverseBracket = ']'
		} else if top == '(' {
			reverseBracket = ')'
		}

		goStack = comparePrice(cheapOpenBracket, reverseBracket)
		if goStack {
			stack = append(stack, cheapOpenBracket)
			builder.WriteRune(cheapOpenBracket)
			countOpenBrackets++
		} else {
			builder.WriteRune(reverseBracket)
			stack = stack[:len(stack)-1]
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
	scanner.Buffer(make([]byte, 200000), 200000)
	writer := bufio.NewWriter(out)
	defer func() { _ = writer.Flush() }()

	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		w := scanner.Text()
		setPrice(w)
		scanner.Scan()
		s := scanner.Text()
		_, _ = fmt.Fprintln(writer, minimumPsp(w, s, n))
	}
}
