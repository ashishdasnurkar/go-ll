package main

import "fmt"

type Cell struct {
	data string
	next *Cell
}

func main() {
	aCell := Cell{"Apple", nil}
	bCell := Cell{data: "Banana"}
	aCell.next = &bCell
	top := &aCell
	for cell := top; cell != nil; cell = cell.next {
		fmt.Printf("%s ", cell.data)
	}
	fmt.Println()
}
