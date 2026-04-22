package main

import (
	"bufio"
	"fmt"
	"strings"
	"strconv"
	"os"
)

func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	var result strings.Builder
	fmt.Fprintf(&result, "%d", slice[0])
	for _, value := range slice[1:] {
		fmt.Fprintf(&result, "%s%d", sep, value)
	}
	return result.String()
}


type Vector struct {
	data     []int
	size     int
	capacity int
}

//init
func NewVector(capacity int) *Vector {
	return &Vector{
		data:     make([]int, capacity), // nunca use len(data) ou cap(data) ou qq método do go de manipulação de array
		size:     0,
		capacity: capacity,
	}
}

/*---------------------------------------------Status---------------------------------------------*/
func (v *Vector) Status() string{
	return fmt.Sprintf("size:%d capacity:%d", v.size, v.capacity)
}

/*String---------------------------------------------*/
// bem parecido com os conceitos de poo na hora de exibir os elementos em uma formatação específica
func (v *Vector) String() string{

	return fmt.Sprintf("[%s]",Join(v.data[:v.size], ", "))

}

/*---------------------------------------------PushBack---------------------------------------------*/
func (v *Vector) PushBack( value int) {

	if(v.size == v.capacity){
		newCapacity := 1

		if(v.capacity > 0){
			newCapacity = v.capacity * 2
		}

		v.Reserve(newCapacity)
	}

	v.data[v.size] = value 
	v.size++
	
	// v.data = append(v.data, value)
	
}
/*---------------------------------------------PopBack---------------------------------------------*/

func(v *Vector) PopBack() (int, error) {

	if(v.size == 0){
		return 0, fmt.Errorf("vector is empty")
	}

	value := v.data[v.size-1]
	v.size--

	return value, nil
}
/*---------------------------------------------Insert---------------------------------------------*/

func (v *Vector) Insert(index int, value int) error {

	if(index < 0 || index > v.size){
		return fmt.Errorf("index out of range")
	}

	if(v.size == v.capacity){
		newCapacity := 1

		if(v.capacity > 0 ){
			newCapacity = v.capacity * 2 
		}

		v.Reserve(newCapacity)
	}

	for i := v.size ; i > index ;i-- {
		v.data[i] = v.data[i-1]
	}
	
	v.data[index] = value
	v.size++

	return nil

}
/*---------------------------------------------Erase---------------------------------------------*/

func (v *Vector) Erase(index int) error{

	if(index < 0 || index > v.size){
		return fmt.Errorf("index out of range")
	}

	for i := index; i < v.size -1; i++{
		v.data[i] = v.data[i+1]
	}

	v.size--
	
	return nil

}

/*---------------------------------------------Reserve---------------------------------------------*/
func (v *Vector) Reserve(newCapacity int) () {

	if(newCapacity <= v.capacity){
		return
	}

	newData := make([]int, newCapacity)

	for i := 0; i < v.size; i++{
		newData[i] = v.data[i]
	}

	v.data = newData
	v.capacity = newCapacity

}

/*---------------------------------------------At---------------------------------------------*/

func (v *Vector) At(index int) (int,error) {

	if( index < 0 || index >= v.size){
		return 0, fmt.Errorf("index out of range")
	}

	return v.data[index], nil

}

/*---------------------------------------------Set---------------------------------------------*/
func (v *Vector) Set(index int, newValue int ) error{
	/*
		if index < 0 || index >= v.size {
			return fmt.Errorf("index out of range")
		}
		v.data[index] = value
		return nil
	*/
	_, err := v.At(index)
	if(err != nil){
		return err
	}

	v.data[index] = newValue
	return nil
}
/*---------------------------------------------Get---------------------------------------------*/

// func ()

/*---------------------------------------------IndexOf---------------------------------------------*/

func (v *Vector) IndexOf(value int) int{

	// for index, valueData := range v.data[:v.size]{
	// 	if(value == valueData){
	// 		return index
	// 	}
	// }

	for i := 0; i < v.size; i++{
		if(value == v.data[i]){
			return i
		}
	}

	return -1

}

/*---------------------------------------------Contains---------------------------------------------*/
func (v *Vector) Contains(value int) bool{

	// return v.IndexOf(value) != -1

	for _,valueData := range v.data[:v.size]{

		if(valueData == value){
			return true
		}

	}

	return false

}
/*---------------------------------------------Slice---------------------------------------------*/
func(v *Vector) Slice(start int, end int) *Vector{


	//parecido com o exercício de nome princesa se não me falha a memória
	start = ((start % v.size) + v.size) % v.size
	end = ((end % v.size) + v.size) % v.size

	length := (end - start + v.size ) % v.size

	newData := make([]int,length)

	i := start
	for j := 0; j < length; j++{
		newData[j] = v.data[i]
		i = (i + 1) % v.size
	}


	return &Vector{
		data : newData,
		size :	length,
		capacity: length,
	}


	
}

/*---------------------------------------------Clear---------------------------------------------*/
func (v *Vector) Clear() {
	v.size = 0
}

/*---------------------------------------------Capacity--------J-------------------------------------*/
func (v *Vector) Capacity() int{
	return v.capacity
}


func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	v := NewVector(0)
	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
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
			v = NewVector(value)
		case "push":
			for _, part := range parts[1:] {
				value, _ := strconv.Atoi(part)
				v.PushBack(value)
			}
		case "show":
			// bem parecido com os conceitos de poo na hora de exibir os elementos em uma formatação específica
			fmt.Println(v)
		case "status":
			fmt.Println(v.Status())
			// Println(Status())
		case "pop":
			_,err := v.PopBack()
			if err != nil {
				fmt.Println(err)
			}
		case "insert":
			index, _ := strconv.Atoi(parts[1])
			value, _ := strconv.Atoi(parts[2])
			err := v.Insert(index, value)
			if err != nil {
				fmt.Println(err)
			}
		case "erase":
			index, _ := strconv.Atoi(parts[1])
			err := v.Erase(index)
			if err != nil {
				fmt.Println(err)
			}
		case "indexOf":
			value, _ := strconv.Atoi(parts[1])
			index := v.IndexOf(value)
			fmt.Println(index)
		case "contains":
			value, _ := strconv.Atoi(parts[1])
			if v.Contains(value) {
				fmt.Println("true")
			} else {
				fmt.Println("false")
			}
		case "clear":
			v.Clear()
		case "capacity":
			fmt.Println(v.Capacity())
		case "get":
			index, _ := strconv.Atoi(parts[1])
			value, err := v.At(index)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(value)
			}
		case "set":
			index, _ := strconv.Atoi(parts[1])
			value, _ := strconv.Atoi(parts[2])
			err := v.Set(index, value)
			if err != nil {
				fmt.Println(err)
			}
			
		case "reserve":
			newCapacity, _ := strconv.Atoi(parts[1])
			v.Reserve(newCapacity)
		case "slice":
			start, _ := strconv.Atoi(parts[1])
			end, _ := strconv.Atoi(parts[2])
			slice := v.Slice(start, end)
			fmt.Println(slice)
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
