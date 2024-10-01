package main

import "fmt"

type Cell struct {
	data string
	next *Cell
}
type LinkedList struct {
	sentinel *Cell
}

func makeLinkedList() LinkedList {
	newList := LinkedList{}
	newList.sentinel = &Cell{"SENTINEL", nil}
	return newList
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
