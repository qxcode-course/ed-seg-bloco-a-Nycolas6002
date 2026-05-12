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
	root *Node
}

type LList struct{
	root *Node
	size int
}

func (n *Node) Next() *Node{

	if(n.next == n.root){
		return nil
	}
	return n.next

}

func (n *Node) Prev() *Node{

	if(n.prev == n.root){
		return nil
	}
	return n.prev

}

func (ll *LList) Front() *Node{

	if(ll.Empty()){
		return nil
	}
	return ll.root.next

}

func (ll *LList) Back() *Node{

	if(ll.Empty()){
		return nil
	}
	return ll.root.prev

}

func (ll *LList) Empty() bool {
	return ll.size == 0
}

func NewLList() *LList{

	root := &Node{}
	root.prev = root
	root.next = root
	root.root = root

	return &LList{
		root : root,
		size: 0,
	}
}

func (ll *LList) Size() int{ return ll.size}

func (ll *LList) PushFront(value int) {

	firstElement := ll.root.next

	newNode := &Node{
		value: value,
		next:firstElement,
		prev: ll.root,
		root: ll.root,
	}

	firstElement.prev = newNode
	ll.root.next = newNode 

	if(ll.size == 0){
		ll.root.prev = newNode
	}

	ll.size++
}

func (ll *LList) PushBack(value int) {

	lastElement := ll.root.prev

	newNode := &Node{
		value: value,
		next:  ll.root,
		prev:  lastElement,
		root:  ll.root,
	}

	lastElement.next = newNode
	ll.root.prev = newNode

	if ll.size == 0 {
		ll.root.next = newNode
	}

	ll.size++
}

func (ll *LList) String() {

	if(ll.Empty()){
		fmt.Print("[ ]")
		return 
	}

	fmt.Print("[")
	
	for it := ll.Front(); it != nil; it = it.Next(){
		fmt.Print(it.value)
		
		if it.Next() != nil{
			fmt.Print(", ")
		}
	}

	fmt.Println("]")

}


func (ll *LList) Clear() {
	ll.root.next = ll.root
	ll.root.prev = ll.root
	ll.size = 0
}

func (ll *LList) Search(value int) *Node{
	for it := ll.Front(); it != nil; it = it.Next(){
		
		if(it.value == value){
			return it
		}
	}

	return nil
}


func (ll *LList) Insert(node *Node, value int){

	if(node == nil){
		fmt.Println("fail: not found")
		return
	}

	before := node.prev

	newNode := &Node{
		value: value,
		next: node,
		prev: before,
		root: ll.root,
	}

	before.next = newNode
	node.prev = newNode

	ll.size++

}

func (ll *LList) Replace(node *Node, value int){

	if(node == nil){
		fmt.Println("fail: not found")
		return
	}

	node.value = value

}

func (ll *LList) Remove(node *Node){

	if(node == nil){
		fmt.Println("fail: not found")
		return
	}

	before := node.prev
	after := node.next

	before.next = after
	after.prev = before

	ll.size--

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
			// ll.PopBack()
		case "pop_front":
			// ll.PopFront()
		case "clear":
			ll.Clear()
		case "walk":
			fmt.Print("[ ")
			for node := ll.Front(); node != nil; node = node.Next() {
				fmt.Printf("%v ", node.value)
			}
			fmt.Print("]\n[ ")
			for node := ll.Back(); node != nil; node = node.Prev() {
				fmt.Printf("%v ", node.value)
			}
			fmt.Println("]")
		case "replace":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
				node.value = newvalue
			} else {
				fmt.Println("fail: not found")
			}
		case "insert":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
				ll.Insert(node, newvalue)
			} else {
				fmt.Println("fail: not found")
			}
		case "remove":
			oldvalue, _ := strconv.Atoi(args[1])
			node := ll.Search(oldvalue)
			if node != nil {
				ll.Remove(node)
			} else {
				fmt.Println("fail: not found")
			}
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
