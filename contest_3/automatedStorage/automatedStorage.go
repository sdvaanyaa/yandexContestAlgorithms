package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type rover struct {
	index int
	time  int
}

func roadsEmpty(roads [][]rover) bool {
	for _, road := range roads {
		if len(road) > 0 {
			return false
		}
	}
	return true
}

func sortRoversByTimes(roads [][]rover) [][]rover {
	for i := 0; i < 4; i++ {
		sort.Slice(roads[i], func(j, k int) bool {
			return roads[i][j].time < roads[i][k].time
		})
	}
	return roads
}

func calculation(a, b, n int, roads [][]rover) []int {
	var left = []int{3, 0, 1, 2}
	var right = []int{1, 2, 3, 0}
	var opposite = []int{2, 3, 0, 1}

	mainRoads := make([]bool, 4)
	mainRoads[a] = true
	mainRoads[b] = true

	currentTime := 1
	result := make([]int, n)
	passed := make([]bool, 4)

	canDriveMainRoad := func(i int) bool {
		return !mainRoads[left[i]] || len(roads[left[i]]) == 0 || roads[left[i]][0].time > currentTime
	}

	canDriveSecondaryRoad := func(i int) bool {
		return (!mainRoads[right[i]] || len(roads[right[i]]) == 0 || roads[right[i]][0].time > currentTime) &&
			(len(roads[left[i]]) == 0 || roads[left[i]][0].time > currentTime) &&
			(!mainRoads[opposite[i]] || len(roads[opposite[i]]) == 0 || roads[opposite[i]][0].time > currentTime)
	}

	for !roadsEmpty(roads) {
		for i := 0; i < 4; i++ {
			if len(roads[i]) > 0 && roads[i][0].time <= currentTime {
				if mainRoads[i] && canDriveMainRoad(i) {
					result[roads[i][0].index] = currentTime
					passed[i] = true
				} else if !mainRoads[i] && canDriveSecondaryRoad(i) {
					result[roads[i][0].index] = currentTime
					passed[i] = true
				}
			}
		}

		for i := 0; i < 4; i++ {
			if passed[i] {
				roads[i] = roads[i][1:]
				passed[i] = false
			}
		}
		currentTime++
	}

	return result
}

func main() {
	inp, _ := os.Open("input.txt")
	defer func() { _ = inp.Close() }()
	out, _ := os.Create("output.txt")
	defer func() { _ = out.Close() }()
	scanner := bufio.NewScanner(inp)
	writer := bufio.NewWriter(out)
	defer func() { _ = writer.Flush() }()

	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		aB := strings.Split(scanner.Text(), " ")
		a, _ := strconv.Atoi(aB[0])
		a--
		b, _ := strconv.Atoi(aB[1])
		b--

		roads := make([][]rover, 4)
		for i := 0; i < n; i++ {
			scanner.Scan()
			dT := strings.Split(scanner.Text(), " ")
			d, _ := strconv.Atoi(dT[0])
			t, _ := strconv.Atoi(dT[1])
			d--
			roads[d] = append(roads[d], rover{i, t})
		}

		roads = sortRoversByTimes(roads)
		result := calculation(a, b, n, roads)
		for i := 0; i < len(result); i++ {
			_, _ = fmt.Fprintln(writer, result[i])
		}
	}
}
