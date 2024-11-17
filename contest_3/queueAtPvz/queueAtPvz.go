package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calculation(n int, b uint64, clients []uint64) uint64 {
	var totalTime uint64
	var waitingClients uint64

	for i := 0; i < n; i++ {
		waitingClients += clients[i]

		if waitingClients <= b {
			totalTime += waitingClients
			waitingClients = 0
		} else {
			totalTime += b
			waitingClients -= b
		}

		totalTime += waitingClients
	}

	totalTime += waitingClients
	return totalTime
}

func main() {
	in, _ := os.Open("input.txt")
	defer func() { _ = in.Close() }()
	out, _ := os.Create("output.txt")
	defer func() { _ = out.Close() }()
	scanner := bufio.NewScanner(in)
	scanner.Buffer(make([]byte, 1024*1024*2), 1024*1024*10)
	writer := bufio.NewWriter(out)
	defer func() { _ = writer.Flush() }()

	for scanner.Scan() {
		nB := strings.Split(scanner.Text(), " ")
		n, _ := strconv.Atoi(nB[0])
		b, _ := strconv.ParseUint(nB[1], 10, 64)
		scanner.Scan()
		strA := strings.Split(scanner.Text(), " ")
		clients := make([]uint64, n)
		for i := 0; i < n; i++ {
			clients[i], _ = strconv.ParseUint(strA[i], 10, 64)
		}

		_, _ = fmt.Fprintln(writer, calculation(n, b, clients))
	}
}
