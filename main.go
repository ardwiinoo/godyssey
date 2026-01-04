package main

import (
	"github.com/ardwiinoo/godyssey/linked_list"
)

func main() {
	myList := linked_list.NewLinkedList()
	myList.Prepend(linked_list.NewNode(1))
	myList.Prepend(linked_list.NewNode(2))
	myList.Prepend(linked_list.NewNode(5))
	myList.PrintListData()

	myList.DeleteWithValue(2)
	myList.DeleteWithValue(100)
	myList.PrintListData()
}
