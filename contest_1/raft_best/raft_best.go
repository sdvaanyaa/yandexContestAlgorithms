package main

import "fmt"

func main() {
	var x1, y1, x2, y2, x, y int
	fmt.Scanf("%d\n%d\n%d\n%d\n%d\n%d\n", &x1, &y1, &x2, &y2, &x, &y)
	ans := ""
	if y > y2 {
		ans += "N"
	} else if y < y1 {
		ans += "S"
	}

	if x > x2 {
		ans += "E"
	} else if x < x1 {
		ans += "W"
	}

	fmt.Println(ans)
}
