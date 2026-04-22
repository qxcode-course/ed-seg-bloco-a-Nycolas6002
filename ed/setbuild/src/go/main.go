package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type Set struct{
	data []int
	size int
	capacity int
}

/*-----------------------------------NewSet-----------------------------------*/
func NewSet( capacity int) *Set{

	return &Set{
		data: make([]int, capacity),
		size: 0,
		capacity: capacity,
	}

}
/*-----------------------------------String-----------------------------------*/
func (s *Set) String() string{

	if(s.size == 0){
		return "[]"
	}

	result := "["

	for index,value := range s.data[:s.size]{
		
		result += strconv.Itoa(value)
		if(index < s.size - 1){
			result += ", "
		}
	}
	result += "]"

	return result
}

/*-----------------------------------BinarySearch-----------------------------------*/
func (s *Set) BinarySearch(value int) int{

	left := 0
	rigth := s.size - 1
	
	for left <= rigth{
		middle := (left + rigth) / 2
		if(s.data[middle] ==  value){
			return middle
		}
		if(s.data[middle] < value){
			left = middle + 1
		}else{
			rigth = middle - 1
		}
	}

	return -1
}
/*-----------------------------------Contains-----------------------------------*/
func (s *Set) Contains(value int) bool{

	if(s.BinarySearch(value) != -1){
		return true
	}

	return false

}

/*-----------------------------------Insert-----------------------------------*/
func (s *Set) insert( index int, value int) {

	if(s.size == s.capacity){
		s.Reserve(s.capacity * 2)
	}

	s.size++
	s.data[index] = value

}

// func (s *Set) Insert(value int) {

// 	if(s.BinarySearch(value) != -1){
// 		return
// 	}


// 	s.size++
// 	s.data[0] = value

// }

/*-----------------------------------Reserve-----------------------------------*/
func (s *Set) Reserve(newCapacity int) {

	if(s.capacity <= newCapacity){
		return
	}

	newData := make([]int, newCapacity)

	for i := 0; i < s.size; i++{
		newData[i] = s.data[i]
	}

	s.data = newData
	s.capacity = newCapacity	

}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	s := NewSet(0)
	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		fmt.Println(line)
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		cmd = parts[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(parts[1])
			s = NewSet(value)
		case "insert":
			// for _, part := range parts[1:] {
			// 	value, _ := strconv.Atoi(part)
			// 	s.insert(value)
			// }
			for index, part := range parts[1:] {
				value, _ := strconv.Atoi(part)
				s.insert(strconv.Atoi(index),value)
			}
		case "show":
			fmt.Println(s.String())
			fmt.Println(s.size, s.capacity)
			// fmt.Println(s.data)
		case "erase":
			// value, _ := strconv.Atoi(parts[1])
		case "contains":
			value, _ := strconv.Atoi(parts[1])
			s.Contains(value)
		case "clear":
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
