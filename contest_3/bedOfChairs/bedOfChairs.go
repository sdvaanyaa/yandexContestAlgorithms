package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Chair struct {
	width  int
	height int
	index  int
}

func addToDeque(heightDeque *[]Chair, chairs []Chair, idx int) {
	previousHeight := 0
	if idx-1 >= 0 {
		previousHeight = chairs[idx-1].height
	} else {
		previousHeight = chairs[idx].height
	}
	currentDiff := chairs[idx].height - previousHeight

	for len(*heightDeque) > 0 && (*heightDeque)[len(*heightDeque)-1].height < currentDiff {
		*heightDeque = (*heightDeque)[:len(*heightDeque)-1]
	}

	*heightDeque = append(*heightDeque, Chair{height: currentDiff, index: idx})
}

func getMaxHeight(heightDeque []Chair, leftBound int) int {
	for len(heightDeque) > 0 && heightDeque[0].index <= leftBound {
		heightDeque = heightDeque[1:]
	}

	if len(heightDeque) == 0 {
		return 0
	}
	return heightDeque[0].height
}

func minDiscomfort(chairs []Chair, n int, H int) int {
	heightDeque := []Chair{}
	left := 0
	currentWidth := 0
	minDisc := math.MaxInt

	for right := 0; right < n; right++ {
		currentWidth += chairs[right].width
		addToDeque(&heightDeque, chairs, right)

		for currentWidth >= H {
			minDisc = min(minDisc, getMaxHeight(heightDeque, left))
			currentWidth -= chairs[left].width
			left++
		}
	}

	return minDisc
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	inp, _ := os.Open("input.txt")
	defer func() { _ = inp.Close() }()
	out, _ := os.Create("output.txt")
	defer func() { _ = out.Close() }()

	scanner := bufio.NewScanner(inp)
	scanner.Buffer(make([]byte, 1024*1024*2), 1024*1024*8)

	writer := bufio.NewWriter(out)
	defer func() { _ = writer.Flush() }()

	scanner.Scan()
	nH := strings.Split(scanner.Text(), " ")
	n, _ := strconv.Atoi(nH[0])
	H, _ := strconv.Atoi(nH[1])

	height := make([]int, n)
	width := make([]int, n)

	scanner.Scan()
	heightList := strings.Split(scanner.Text(), " ")
	for i := 0; i < n; i++ {
		height[i], _ = strconv.Atoi(heightList[i])
	}

	scanner.Scan()
	widthList := strings.Split(scanner.Text(), " ")
	for i := 0; i < n; i++ {
		width[i], _ = strconv.Atoi(widthList[i])
	}

	chairs := make([]Chair, n)
	for i := 0; i < n; i++ {
		chairs[i] = Chair{width: width[i], height: height[i], index: i}
	}

	sort.Slice(chairs, func(i, j int) bool {
		return chairs[i].height < chairs[j].height
	})

	minDisc := minDiscomfort(chairs, n, H)
	_, _ = fmt.Fprintln(writer, minDisc)
}
