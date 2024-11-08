package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func slowSolution(slice []int) int {
	minTransfer := math.MaxInt
	nowTransfer := 0
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice); j++ {
			if i != j {
				nowTransfer += slice[j] * abs(i-j)
			}
		}
		if nowTransfer < minTransfer {
			minTransfer = nowTransfer
		}
		nowTransfer = 0
	}
	return minTransfer
}

func optimizedSolution(slice []int) int {
	prefixSumsStraight := make([]int, len(slice))
	prefixSumsStraight[0] = slice[0]
	prefixSumsReverse := make([]int, len(slice))
	prefixSumsReverse[len(slice)-1] = slice[len(slice)-1]
	for i := 1; i < len(slice); i++ {
		prefixSumsStraight[i] = prefixSumsStraight[i-1] + slice[i]
		prefixSumsReverse[len(slice)-1-i] = prefixSumsReverse[len(slice)-i] + slice[len(slice)-1-i]
	}

	transition := make([]int, len(slice))
	sums := 0
	for i := 1; i < len(slice); i++ {
		sums += slice[i] * i
		transition[0] = sums
	}

	minSums := transition[0]
	for i := 1; i < len(slice); i++ {
		transition[i] = transition[i-1] - prefixSumsReverse[i] + prefixSumsStraight[i-1]
		if transition[i] < minSums {
			minSums = transition[i]
		}
	}
	return minSums
}

// Генерация случайного массива
func generateRandomSlice(size, maxVal int) []int {
	slice := make([]int, size)
	for i := range slice {
		slice[i] = rand.Intn(maxVal) + 1
	}
	return slice
}

// Функция для стресс-теста
func stressTest(iterations, size, maxVal int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < iterations; i++ {
		slice := generateRandomSlice(size, maxVal)

		expected := slowSolution(slice)
		result := optimizedSolution(slice)

		if expected != result {
			fmt.Println("Mismatch found!")
			fmt.Println("Input:", slice)
			fmt.Println("Expected:", expected, "Got:", result)
			return
		}
	}
	fmt.Println("Stress test passed!")
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
		for i := range slice {
			slice[i], _ = strconv.Atoi(nums[i])
		}

		//_, _ = fmt.Fprintln(writer, optimizedSolution(slice))
		//_, _ = fmt.Fprintln(writer, slowSolution(slice))
	}
	stressTest(5, 100000, 1000000000)
}
