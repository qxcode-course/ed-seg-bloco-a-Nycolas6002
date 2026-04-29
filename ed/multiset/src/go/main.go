package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	result := fmt.Sprintf("%d", slice[0])
	for _, value := range slice[1:] {
		result += sep + fmt.Sprintf("%d", value)
	}
	return result
}

type MultiSet struct{
	data []int
	size int
	capacity int
}

//Init
func NewMultiSet(capacity int) *MultiSet{
	return &MultiSet{
		data: make([]int,capacity),
		size: 0,
		capacity: capacity,
	}
}

//Show
func Show(ms *MultiSet){
	// fmt.Println("[",Join(ms.data[:ms.size],", "),"]")
	fmt.Printf("[%s]\n", Join(ms.data[:ms.size], ", "))
}

func (ms* MultiSet) findPosition(value int) int{

	left := 0
	right := ms.size

	for left < right{
		middle := (left + right) / 2

		if(ms.data[middle] < value){
			left = middle + 1
		}else{
			right = middle
		}

	}

	return left

}

func (ms *MultiSet) Expand(newCapacity int){

	if(newCapacity <= ms.capacity){
		return
	}

	newData := make([]int, newCapacity)

	for i := 0; i < ms.size; i++{
		newData[i] = ms.data[i]
	}

	ms.data = newData
	ms.capacity = newCapacity

}

func (ms *MultiSet) insert(value int, index int){

	if(ms.size == ms.capacity){
		newCapacity := 1

		if(ms.capacity > 0){
			newCapacity = ms.capacity * 2
		}

		ms.Expand(newCapacity)
	}

	for i := ms.size; i > index; i--{
		ms.data[i] = ms.data[i-1]

	}

	ms.data[index] = value
	ms.size++

}


func (ms *MultiSet) Insert(value int){

	index := ms.findPosition(value)
	ms.insert(value, index)
}


func (ms *MultiSet) Contains(value int) bool{

	for i := 0; i < ms.size;i++{
		if(ms.data[i] == value){
			return true
		}
	}

	return false
	
}

func (ms *MultiSet) Count(value int) int{

	var count int
	for i := 0; i < ms.size; i++{
		if(ms.data[i] == value){
			count++
		}
	}

	return count

}

func (ms *MultiSet) erase(value int) error {

	for i := 0 ; i < ms.size; i++ {
		if(ms.data[i] == value){
			return nil
		}
	}

	return fmt.Errorf("value not found")

}


func (ms *MultiSet) Erase(value int) {

	returnFunction := ms.erase(value)

	if(returnFunction == nil ){

		var index int

		for i := 0; i < ms.size;i++{
			if(ms.data[i] == value){
				index = i
				break
			}
		}

		for index < ms.size-1 {
			ms.data[index] = ms.data[index + 1]
			index++
		}

		ms.size--
	}else{
		fmt.Println(returnFunction)
	}


}

func (ms *MultiSet) Unique() int{
	

	if(ms.size == 0){
		return 0
	}

	count := 1
	variableAux := ms.data[0]

	for i := 0; i < ms.size; i++{
		if(variableAux != ms.data[i]){
			count++
			variableAux = ms.data[i]
		}

	}

	return count

}

func (ms *MultiSet) Clear(){
	ms.size = 0
}


func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)
	ms := NewMultiSet(0)

	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		args := strings.Fields(line)
		fmt.Println(line)
		if len(args) == 0 {
			continue
		}
		cmd = args[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(args[1])
			ms = NewMultiSet(value)
		case "insert":
			for _, part := range args[1:] {
				value, _ := strconv.Atoi(part)
				ms.Insert(value)
			}
		case "show":
			Show(ms)
		case "erase":
			value, _ := strconv.Atoi(args[1])
			ms.Erase(value)
		case "contains":
			value, _ := strconv.Atoi(args[1])
			fmt.Println(ms.Contains(value))
		case "count":
			value, _ := strconv.Atoi(args[1])
			fmt.Println(ms.Count(value))
		case "unique":
			fmt.Println(ms.Unique())
			// ms.Unique()
		case "clear":
			ms.Clear()

		default:
			fmt.Println("fail: comando invalido")
		}
	}
}