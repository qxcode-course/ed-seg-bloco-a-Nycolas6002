package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type Node struct{
	value int
	next *Node
	prev *Node
}

type ListNode struct{
	root *Node
	size int

}

func NewLList() *ListNode{

	root := &Node{}
	root.next = root
	root.prev = root

	return &ListNode{
		root : root,
		size : 0,
	}

}

func (ll *ListNode) Size() int{
	return ll.size
}

func (ll *ListNode) PushFront(value int) {
	
	newNode := &Node{
		value: value,
	}

	firstElement := ll.root.next

	newNode.next = firstElement
	newNode.prev = ll.root
	firstElement.prev = newNode

	ll.root.next = newNode //como de fosse o head

	if(ll.size == 0){
		ll.root.prev = newNode
	}
	
	ll.size++
}


func (ll *ListNode) PushBack(value int) {

	newNode := &Node{
		value: value,
	}

	lastElement := ll.root.prev
	newNode.prev = lastElement
	newNode.next = ll.root
	
	lastElement.next = newNode
	ll.root.prev = newNode
	
	if(ll.size == 0){
		ll.root.next = newNode
	}

	ll.size++
}

func (ll *ListNode) PopFront() {

	if(ll.Empty()){
		return
	}


	firstElement := ll.root.next
	newfirstElement := firstElement.next

	newfirstElement.prev = ll.root
	ll.root.next = newfirstElement
	
	
	if(ll.size == 1){
		ll.root.next = ll.root 
	}

	ll.size--
}

func (ll *ListNode) PopBack() {

	if(ll.Empty()){
		return
	}


	lastElement := ll.root.prev
	newLastElement := lastElement.prev

	newLastElement.next = ll.root
	ll.root.prev = newLastElement
	
	
	if(ll.size == 1){
		ll.root.next = ll.root 
	}

	ll.size--
}

func (ll *ListNode) Clear() {
	
	ll.root.next = ll.root
	ll.root.prev = ll.root
	ll.size = 0
}



func (ll *ListNode) Empty() bool {
	return ll.size == 0
}

func (ll *ListNode) String()  {


	fmt.Print("[")

	current := ll.root.next
	
	for current != ll.root{
		fmt.Print(current.value)
		
		if current.next != ll.root{
			fmt.Print(", ")
		}
		current = current.next
	}

	fmt.Println("]")
}


func main() {
	scanner := bufio.NewScanner(os.Stdin)
	ll := NewLList()

	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		fmt.Println(line)
		args := strings.Fields(line)

		if len(args) == 0 {
			continue
		}

		cmd := args[0]

		switch cmd {
		case "show":
			ll.String()
		case "size":
			fmt.Println(ll.Size())
		case "push_back":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushBack(num)
			}
		case "push_front":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushFront(num)
			}
		case "pop_back":
			ll.PopBack()
		case "pop_front":
			ll.PopFront()
		case "clear":
			ll.Clear()
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
