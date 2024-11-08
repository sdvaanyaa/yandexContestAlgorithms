package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type criterion struct {
	a int
	b int
	i int
}

func customStableSort(n int, a, b, p []int) ([]criterion, []criterion) {
	sortByA := make([]criterion, n)
	sortByB := make([]criterion, n)

	for i := 0; i < n; i++ {
		sortByA[i] = criterion{a: a[i], b: b[i], i: i + 1}
		sortByB[i] = criterion{a: a[i], b: b[i], i: i + 1}
	}

	sort.SliceStable(sortByA, func(i, j int) bool {
		if sortByA[i].a != sortByA[j].a {
			return sortByA[i].a > sortByA[j].a
		}
		if sortByA[i].b != sortByA[j].b {
			return sortByA[i].b > sortByA[j].b
		}
		return sortByA[i].i < sortByA[j].i
	})

	sort.SliceStable(sortByB, func(i, j int) bool {
		if sortByB[i].b != sortByB[j].b {
			return sortByB[i].b > sortByB[j].b
		}
		if sortByB[i].a != sortByB[j].a {
			return sortByB[i].a > sortByB[j].a
		}
		return sortByB[i].i < sortByB[j].i
	})

	return sortByA, sortByB
}

func optimizedSolution(n int, a, b, p []int) []int {
	sortByA, sortByB := customStableSort(n, a, b, p)
	result := make([]int, n)
	alreadyBeen := make([]bool, n+1)

	indexA, indexB := 0, 0

	for i := 0; i < n; i++ {
		var index int
		if p[i] == 0 {
			for indexA < n {
				index = sortByA[indexA].i
				if !alreadyBeen[index] {
					result[i] = index
					alreadyBeen[index] = true
					indexA++
					break
				}
				indexA++
			}
		} else if p[i] == 1 {
			for indexB < n {
				index = sortByB[indexB].i
				if !alreadyBeen[index] {
					result[i] = index
					alreadyBeen[index] = true
					indexB++
					break
				}
				indexB++
			}
		}
	}
	return result
}

func main() {
	in, _ := os.Open("input.txt")
	defer func() { _ = in.Close() }()
	out, _ := os.Create("output.txt")
	defer func() { _ = out.Close() }()
	scanner := bufio.NewScanner(in)
	scanner.Buffer(make([]byte, 1024*1024*4), 1024*1024*4)
	writer := bufio.NewWriter(out)
	defer func() { _ = writer.Flush() }()

	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		strA := strings.Split(scanner.Text(), " ")
		scanner.Scan()
		strB := strings.Split(scanner.Text(), " ")
		scanner.Scan()
		strP := strings.Split(scanner.Text(), " ")
		a := make([]int, n)
		b := make([]int, n)
		p := make([]int, n)
		for i := 0; i < n; i++ {
			a[i], _ = strconv.Atoi(strA[i])
			b[i], _ = strconv.Atoi(strB[i])
			p[i], _ = strconv.Atoi(strP[i])
		}
		slice := optimizedSolution(n, a, b, p)
		var str strings.Builder
		for i := 0; i < n; i++ {
			str.WriteString(strconv.Itoa(slice[i]))
			str.WriteByte(' ')
		}
		_, _ = fmt.Fprintln(writer, str.String())
	}
}
