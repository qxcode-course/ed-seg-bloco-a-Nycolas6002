package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Value int
	next  *Node
	prev  *Node
	root  *Node
}

type LList struct {
	root *Node
	size int
}

func NewLList() *LList {
	list := &LList{}
	list.root = &Node{root: nil}
	list.root.next = list.root
	list.root.prev = list.root
	list.root.root = list.root // nó sentinela aponta pra si mesmo
	return list
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

func (n *Node) Next() *Node {
	if n.next == n.root {
		return nil
	}
	return n.next
}

func (n *Node) Prev() *Node {
	if n.prev == n.root {
		return nil
	}
	return n.prev
}

// func (l *LList) PushFront(value int) {
// 	l.insertBefore(l.root.next, value)
// }

func (l *LList) PushBack(value int) {
	l.insertBefore(l.root, value)
}

func (l *LList) insertBefore(mark *Node, value int) {
	n := &Node{
		Value: value,
		root:  l.root,
	}
	n.prev = mark.prev
	n.next = mark
	mark.prev.next = n
	mark.prev = n

	l.size++
}


func str2list(serial string) *LList {
	serial = serial[1 : len(serial)-1]
	ll := NewLList()
	if serial == "" {
		return ll
	}
	for _, p := range strings.Split(serial, ",") {
		value, _ := strconv.Atoi(p)
		ll.PushBack(value)
	}
	return ll
}

func (l *LList) Empty() bool{ return l.size == 0}


func (ll *LList) Size() int{ return ll.size}

func (l *LList) Insert(node *Node, value int){
	if(l.Empty()){
		return
	}

	l.insertBefore(node, value)
}

func (l *LList) String() {

	if(l.Empty()){
		fmt.Print("[ ]")
		return 
	}

	fmt.Print("[")
	
	for it := l.Front(); it != nil; it = it.Next(){
		fmt.Print(it.Value)
		
		if it.Next() != nil{
			fmt.Print(", ")
		}
	}

	fmt.Println("]")

}

func addsorted(ll *LList, value int) {

	for it := ll.Front(); it != nil; it = it.Next() {

		if value < it.Value {
			ll.Insert(it, value)
			return
		}
	}

	ll.PushBack(value)
}

func equals(lla,llb *LList) bool {

	if(lla.Size() != llb.Size()){
		return false
	}

	ita := lla.Front() 
	itb := llb.Front()


	for ita != nil && itb != nil {

		if(ita.Value != itb.Value ){
			return false
		}

		ita = ita.Next()
		itb =itb.Next()
	
	}

	return true

}

func reverse(ll *LList) {

	current := ll.root

	for{
		current.next, current.prev = current.prev, current.next

		current = current.prev

		if(current == ll.root){
			break
		}
	}

}
func merge(lla, llb *LList) *LList{

	result := NewLList()

	ita := lla.Front()
	itb := llb.Front()

	for ita != nil && itb != nil {

		if(ita.Value < itb.Value){
			result.PushBack(ita.Value)
			ita = ita.Next()
		}else{
			result.PushBack(itb.Value)
			itb = itb.Next()
		}

	}

		for ita != nil {
		result.PushBack(ita.Value)
		ita = ita.Next()
	}

	for itb != nil {
		result.PushBack(itb.Value)
		itb = itb.Next()
	}

	return result

	// for it := ll.Front(); it != nil; it = it.Next() {

	// 	if value < it.Value {
	// 		ll.Insert(it, value)
	// 		return
	// 	}
	// }

	// ll.PushBack(value)
}



func main() {
	scanner := bufio.NewScanner(os.Stdin)

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
		case "compare":
			lla := str2list(args[1])
			llb := str2list(args[2])
			if equals(lla, llb) {
				fmt.Println("iguais")
			} else {
				fmt.Println("diferentes")
			}
		case "addsorted":
			lla := NewLList()
			for i := 1; i < len(args); i++ {
				value, _ := strconv.Atoi(args[i])
				addsorted(lla, value)
			}
			lla.String()
			// fmt.Println(lla)
		case "reverse":
			lla := str2list(args[1])
			reverse(lla)

			if(lla.Empty()){
				fmt.Println("[]")
				break
			}
			lla.String()
			// fmt.Println(lla)
		case "merge":
			lla := str2list(args[1])
			llb := str2list(args[2])
			merged := merge(lla, llb)
			merged.String()
			// fmt.Println(merged)
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
