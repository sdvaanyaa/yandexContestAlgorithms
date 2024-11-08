package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func calculateDistance(x, y, x1, y1 float64) float64 {
	return math.Sqrt(math.Pow(x1-x, 2) + math.Pow(y1-y, 2))
}

func minDistance(slice []float64) string {
	x1, y1, x2, y2, x, y := slice[0], slice[1], slice[2], slice[3], slice[4], slice[5]
	distances := []float64{
		calculateDistance(x, y, x1, y2), // NW
		calculateDistance(x, y, x1, y1), // SW
		calculateDistance(x, y, x2, y1), // SE
		calculateDistance(x, y, x2, y2), // NE,
		// W
		// E
		// N
		// S
	}

	S, N, W, E := 1000000.0, 1000000.0, 1000000.0, 1000000.0

	if x1 < x && x2 > x {
		if y < y1 {
			S = y1 - y
		} else if y > y2 {
			N = y - y2
		}
	} else if y > y1 && y < y2 {
		if x < x1 {
			W = x1 - x
		} else if x > x2 {
			E = x - x2
		}
	}

	distances = append(distances, W, E, N, S)

	directions := []string{"NW", "SW", "SE", "NE", "W", "E", "N", "S"}

	minDistance := distances[0]
	minIndex := 0
	for i, v := range distances {
		if v < minDistance {
			minDistance = v
			minIndex = i
		}
	}

	return directions[minIndex]
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
		for i := 0; i < 6; i++ {
			if !scanner.Scan() {
				return
			}
			line := scanner.Text()
			num, err := strconv.ParseFloat(line, 64)
			if err != nil {
				log.Fatalf("Conversion error: %v", err)
			}
			numbers = append(numbers, num)
		}

		if len(numbers) == 6 {
			result := minDistance(numbers)
			_, err := fmt.Fprintln(writer, result)
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
