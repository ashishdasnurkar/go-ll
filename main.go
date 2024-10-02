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
func (me *Cell) addAfter(after *Cell) {
	after.next = me.next
	me.next = after
}
func main() {
	aCell := Cell{"Apple", nil}
	bCell := Cell{data: "Banana"}
	aCell.next = &bCell
	top := &aCell
	cCell := Cell{"Dragon Fruit", nil}
	(&bCell).addAfter(&cCell)
	for cell := top; cell != nil; cell = cell.next {
		fmt.Printf("%s ", cell.data)
	}
	fmt.Println()
}
