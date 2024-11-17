package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

type tokenType int

const (
	number tokenType = iota
	operator
	openBracket
	closeBracket
	unknown
)

var priority = map[string]int{
	"+": 1,
	"-": 1,
	"*": 2,
}

func getTokenType(r rune) tokenType {
	switch {
	case unicode.IsDigit(r):
		return number
	case r == '+' || r == '-' || r == '*':
		return operator
	case r == '(':
		return openBracket
	case r == ')':
		return closeBracket
	default:
		return unknown
	}
}

func isValidNext(prev, curr tokenType, numberCollected bool) bool {
	switch prev {
	case number:
		if numberCollected && curr == number {
			return false
		}
		return curr == operator || curr == closeBracket || curr == number
	case operator:
		return curr == number || curr == openBracket
	case openBracket:
		return curr == number || curr == openBracket
	case closeBracket:
		return curr == operator || curr == closeBracket
	default:
		return false
	}
}

func checkString(s string) bool {
	re := regexp.MustCompile(`\d+\s+\d+`)
	return !re.MatchString(s)
}

func correctBracketSequence(s string) bool {
	var stack []rune

	for _, v := range s {
		if v == '(' {
			stack = append(stack, '(')
		} else if v == ')' {
			if len(stack) > 0 {
				value := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if value != '(' {
					return false
				}
			} else {
				return false
			}
		}
	}
	if len(stack) > 0 {
		return false
	}
	return true
}

func validateExpression(s string) (bool, []string) {
	noSpaces := checkString(s)
	corrBackSeq := correctBracketSequence(s)
	if !noSpaces || !corrBackSeq {
		return false, nil
	}

	prevType := unknown
	var numBuffer strings.Builder
	var parts []string
	numberCollected := false
	for _, v := range s {
		if unicode.IsSpace(v) {
			continue
		}

		currType := getTokenType(v)
		if currType == unknown {
			return false, nil
		}

		if prevType != unknown && !isValidNext(prevType, currType, numberCollected) {
			return false, nil
		}

		if currType == number {
			numBuffer.WriteRune(v)
			numberCollected = false
		} else {
			if numBuffer.Len() > 0 {
				parts = append(parts, numBuffer.String())
				numBuffer.Reset()
				numberCollected = true
			}
			parts = append(parts, string(v))
		}
		prevType = currType
	}
	if numBuffer.Len() > 0 {
		parts = append(parts, numBuffer.String())
	}
	return prevType == number || prevType == closeBracket, parts
}

func infixToPostfix(s string) ([]string, error) {
	ok, expression := validateExpression(s)
	if !ok {
		return nil, fmt.Errorf("no valid expression")
	}
	var stack []string
	var result []string
	for _, v := range expression {
		if _, err := strconv.Atoi(v); err == nil {
			result = append(result, v)
		}
		switch getTokenType(rune(v[0])) {
		case operator:
			for len(stack) > 0 && priority[stack[len(stack)-1]] >= priority[v] {
				result = append(result, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, v)
		case openBracket:
			stack = append(stack, v)
		case closeBracket:
			for len(stack) > 0 && stack[len(stack)-1] != "(" {
				result = append(result, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1]
		}
	}
	for len(stack) > 0 {
		result = append(result, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}
	return result, nil
}

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

func calculation(s string) (int, error) {
	slice, err := infixToPostfix(s)
	if err != nil {
		return 0, err
	}
	return postfixRecord(slice), nil
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
		if num, err := calculation(s); err == nil {
			_, _ = fmt.Fprintln(writer, num)
		} else {
			_, _ = fmt.Fprintln(writer, "WRONG")
		}
	}
}
