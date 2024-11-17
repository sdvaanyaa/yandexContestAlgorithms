package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type deque []int

func (d *deque) PushFront(value int) {
	*d = append([]int{value}, *d...)
}

func (d *deque) PushBack(value int) {
	*d = append(*d, value)
}

func (d *deque) PopFront() {
	*d = (*d)[1:]
}

func (d *deque) PopBack() {
	*d = (*d)[:len(*d)-1]
}

func (d *deque) Front() int {
	return (*d)[0]
}

func (d *deque) Back() int {
	return (*d)[len(*d)-1]
}

func minimumOnSegment(nums []int, k int) []int {
	var d deque
	result := make([]int, 0, len(nums)-k+1)
	for i := 0; i < len(nums); i++ {
		for len(d) > 0 && d.Back() > nums[i] {
			d.PopBack()
		}
		d.PushBack(nums[i])
		if len(d) > 0 && i >= k-1 {
			result = append(result, d.Front())
			if d.Front() == nums[i-k+1] {
				d.PopFront()
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
	scanner.Buffer(make([]byte, 1024*1024*2), 1024*1024*4)
	writer := bufio.NewWriter(out)
	defer func() { _ = writer.Flush() }()

	for scanner.Scan() {
		nK := strings.Split(scanner.Text(), " ")
		n, _ := strconv.Atoi(nK[0])
		k, _ := strconv.Atoi(nK[1])

		scanner.Scan()
		slice := strings.Split(scanner.Text(), " ")
		nums := make([]int, n)
		for i := 0; i < n; i++ {
			num, _ := strconv.Atoi(slice[i])
			nums[i] = num
		}
		for _, v := range minimumOnSegment(nums, k) {
			_, _ = fmt.Fprintln(writer, v)
		}
	}
}
