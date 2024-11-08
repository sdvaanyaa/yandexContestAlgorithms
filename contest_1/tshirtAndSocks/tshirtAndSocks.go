package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func minQuantity(slice []float64) (int, int) {
	A, B, C, D := slice[0], slice[1], slice[2], slice[3]
	var minTshirt, maxTshirt, minSocks, maxSocks float64
	var a, b int

	if A == 0 {
		a = 1
		b = int(C) + 1
		return a, b
	} else if B == 0 {
		a = 1
		b = int(D) + 1
		return a, b
	} else if C == 0 {
		b = 1
		a = int(A) + 1
		return a, b
	} else if D == 0 {
		b = 1
		a = int(B) + 1
		return a, b
	}

	if A >= B {
		minTshirt = B + 1
		maxTshirt = A + 1
	} else {
		minTshirt = A + 1
		maxTshirt = B + 1
	}

	if C >= D {
		minSocks = D + 1
		maxSocks = C + 1
	} else {
		minSocks = C + 1
		maxSocks = D + 1
	}

	if A < B && C > D {
		if minTshirt < B && minSocks < C {
			a = int(B) + 1
			b = int(C) + 1
			if a < b {
				b = 1
			} else {
				a = 1
			}
			return a, b
		}
	} else if A > B && C < D {
		if minTshirt < A && minSocks < D {
			a = int(A) + 1
			b = int(D) + 1
			if a < b {
				b = 1
			} else {
				a = 1
			}
			return a, b
		}
	}

	case1 := minTshirt + minSocks
	case2 := maxTshirt + 1
	case3 := maxSocks + 1

	minQ := math.Min(math.Min(minTshirt+minSocks, maxTshirt+1), maxSocks+1)

	switch minQ {
	case case1:
		a, b = int(minTshirt), int(minSocks)
	case case2:
		a, b = int(maxTshirt), 1
	case case3:
		a, b = 1, int(maxSocks)
	}
	return a, b
}

func main() {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Unable to open file for reading: %v", err)
	}
	defer func() { _ = inputFile.Close() }()

	outputFile, err := os.Create("output.txt")
	if err != nil {
		log.Fatalf("Unable to open file for writing: %v", err)
	}
	defer func() { _ = outputFile.Close() }()

	scanner := bufio.NewScanner(inputFile)
	writer := bufio.NewWriter(outputFile)

	for {
		numbers := make([]float64, 0)
		for i := 0; i < 4; i++ {
			if !scanner.Scan() {
				return
			}
			line := scanner.Text()
			num, err := strconv.ParseFloat(line, 64)
			if err != nil {
				log.Fatalf("Conversion error: %v", err)
			}
			numbers = append(numbers, num)

			if len(numbers) == 4 {
				a, b := minQuantity(numbers)
				_, err := fmt.Fprintln(writer, a, b)
				if err != nil {
					log.Fatalf("Unable to write result: %v", err)
				}
			}

			if err := scanner.Err(); err != nil {
				log.Fatalf("Scanner error: %v", err)
			}
			if err := writer.Flush(); err != nil {
				log.Fatalf("Error flushing writer: %v", err)
			}
		}
	}
}
