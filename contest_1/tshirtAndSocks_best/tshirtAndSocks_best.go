package main

import "fmt"

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var a, b, c, d int
	var sl [][]int
	fmt.Scanf("%d\n%d\n%d\n%d\n", &a, &b, &c, &d)
	if a > 0 && c > 0 {
		sl = append(sl, []int{b + 1, d + 1})
	}
	if b > 0 && d > 0 {
		sl = append(sl, []int{a + 1, c + 1})
	}
	if a > 0 && b > 0 {
		sl = append(sl, []int{maxInt(a, b) + 1, 1})
	}
	if c > 0 && d > 0 {
		sl = append(sl, []int{1, maxInt(c, d) + 1})
	}

	var minTshirt, minSocks int
	minSum := sl[0][0] + sl[0][1]
	for _, v := range sl {
		if v[0]+v[1] <= minSum {
			minTshirt = v[0]
			minSocks = v[1]
			minSum = v[0] + v[1]
		}
	}

	fmt.Println(minTshirt, minSocks)
}
