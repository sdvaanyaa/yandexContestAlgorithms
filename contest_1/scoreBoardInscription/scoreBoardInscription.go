package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type model struct {
	name  string
	slice []int
}

func minMax(slice []int) (int, int) {
	minn := slice[0]
	maxx := slice[0]

	for _, value := range slice {
		if value < minn {
			minn = value
		}
		if value > maxx {
			maxx = value
		}
	}
	return minn, maxx
}

func form(matrix []string) []string {
	X := make([]int, 0)
	Y := make([]int, 0)
	for y := 0; y < len(matrix); y++ {
		for x := 0; x < len(matrix[y]); x++ {
			if matrix[y][x] == '#' {
				X = append(X, x)
			}
		}
	}

	for x := 0; x < len(matrix); x++ {
		for y := 0; y < len(matrix); y++ {
			if matrix[y][x] == '#' {
				Y = append(Y, y)
			}
		}
	}

	xMin, xMax := minMax(X)
	yMin, yMax := minMax(Y)
	newMatrix := make([]string, 0)
	for y := yMin; y < yMax+1; y++ {
		side := ""
		for x := xMin; x < xMax+1; x++ {
			side += string(matrix[y][x])
		}
		newMatrix = append(newMatrix, side)
	}
	return newMatrix
}

func sampling(matrix []string) []model {
	slice := make([]model, 0)

	if len(matrix) == 1 {
		if matrix[0][0] == '#' {
			slice = append(slice, model{"all", []int{1}})
			return slice
		} else if matrix[0][0] == '.' {
			slice = append(slice, model{"noname", []int{1}})
			return slice
		}
	}

	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] == '.' {
			slice = append(slice, model{"noname", []int{1}})
			return slice
		}
		countTrellis1 := 0
		countTrellis2 := 0
		countPoint := 0
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == '.' {
				countPoint++
			} else if matrix[i][j] == '#' && countPoint == 0 {
				countTrellis1++
			} else if matrix[i][j] == '#' && countPoint != 0 {
				countTrellis2++
			}
		}
		if countTrellis2 == 0 && countPoint == 0 {
			slice = append(slice, model{"all", []int{countTrellis1}})
		} else if countTrellis2 == 0 && countPoint != 0 {
			slice = append(slice, model{"left", []int{countTrellis1, countPoint}})
		} else if countTrellis2 != 0 && countTrellis1 != 0 && countPoint != 0 {
			slice = append(slice, model{"edge", []int{countTrellis1, countPoint, countTrellis2}})
		}
	}

	return slice
}

func slicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func uniqueModels(models []model) []model {
	unique := []model{models[0]}

	for i := 1; i < len(models); i++ {
		if models[i].name != models[i-1].name || !slicesEqual(models[i].slice, models[i-1].slice) {
			unique = append(unique, models[i])
		}
	}
	return unique
}

func scoreboardInscription(matrix []string) string {
	countResh := 0
	for _, value := range matrix {
		for _, element := range value {
			if element == '#' {
				countResh++
			}
		}
	}
	if countResh == 0 {
		return "X"
	}
	matrix = form(matrix)
	samlping := sampling(matrix)
	unique := uniqueModels(samlping)

	patterns := map[string]string{
		"all":               "I",
		"all,edge,all,left": "P",
		"edge,all,edge":     "H",
		"left,all":          "L",
		"all,left,all":      "C",
		"all,edge,all":      "O",
	}

	names := ""
	for i, m := range unique {
		if i > 0 {
			names += ","
		}
		names += m.name
	}

	if letter, exists := patterns[names]; exists {
		if letter == "H" {
			if !slicesEqual(unique[0].slice, unique[2].slice) {
				return "X"
			}
		} else if letter == "P" {
			if unique[1].slice[0] != unique[3].slice[0] {
				return "X"
			}
		}
		return letter
	} else {
		return "X"
	}
}

func main() {
	inputFile, _ := os.Open("input.txt")
	defer func() { _ = inputFile.Close() }()

	outputFile, _ := os.Create("output.txt")
	defer func() { _ = outputFile.Close() }()

	scanner := bufio.NewScanner(inputFile)
	writer := bufio.NewWriter(outputFile)
	defer func() { _ = writer.Flush() }()

	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())

		matrix := make([]string, n)
		for i := 0; i < n; i++ {
			scanner.Scan()
			matrix[i] = scanner.Text()
		}

		_, _ = fmt.Fprintln(writer, scoreboardInscription(matrix))

		_, _ = writer.WriteString("\n")
	}
}
