package main

import (
	"fmt"

	"github.com/ardwiinoo/godyssey/linked_list"
	stackqueue "github.com/ardwiinoo/godyssey/stack_queue"
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

	myStack := stackqueue.NewStack()
	myStack.Push(100)
	myStack.Push(200)
	myStack.Push(300)

	fmt.Println(myStack)
	myStack.Pop()
	fmt.Println(myStack)

	myQueue := stackqueue.NewQueue()
	myQueue.Enqueue(1000)
	myQueue.Enqueue(2000)
	myQueue.Enqueue(3000)

	fmt.Println(myQueue)
	myQueue.Dequeue()
	fmt.Println(myQueue)
}
