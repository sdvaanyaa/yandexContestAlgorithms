package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type node struct {
	data  int
	left  *node
	right *node
}

func insertNode(root *node, data int) (*node, bool) {
	if root == nil {
		return &node{data: data}, true
	}
	if root.data == data {
		return root, false
	}
	var inserted bool
	if data < root.data {
		root.left, inserted = insertNode(root.left, data)
	} else {
		root.right, inserted = insertNode(root.right, data)
	}
	return root, inserted
}

func searchNode(root *node, data int) bool {
	if root == nil {
		return false
	}
	if root.data == data {
		return true
	}
	if data <= root.data {
		return searchNode(root.left, data)
	}
	return searchNode(root.right, data)
}

func inorder(w io.Writer, node *node, level int) {
	if node == nil {
		return
	}
	inorder(w, node.left, level+1)
	fmt.Fprintf(w, "%s%d\n", strings.Repeat(".", level), node.data)
	inorder(w, node.right, level+1)
}

func main() {
	in, _ := os.Open("input.txt")
	defer in.Close()
	out, _ := os.Create("output.txt")
	defer out.Close()
	scanner := bufio.NewScanner(in)
	writer := bufio.NewWriter(out)
	defer writer.Flush()

	var root *node
	for scanner.Scan() {
		s := strings.Fields(scanner.Text())

		command := s[0]
		var data int
		if len(s) > 1 {
			data, _ = strconv.Atoi(s[1])
		}
		switch command {
		case "ADD":
			var inserted bool
			root, inserted = insertNode(root, data)
			if inserted {
				fmt.Fprintln(writer, "DONE")
			} else {
				fmt.Fprintln(writer, "ALREADY")
			}
		case "SEARCH":
			if searchNode(root, data) {
				fmt.Fprintln(writer, "YES")
			} else {
				fmt.Fprintln(writer, "NO")
			}
		case "PRINTTREE":
			inorder(writer, root, 0)
		}
	}
}
