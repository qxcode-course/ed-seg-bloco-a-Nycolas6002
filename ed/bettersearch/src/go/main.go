package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BetterSearch(slice []int, value int) (bool, int) {

	low := 0 
	high := len(slice) - 1
	
	//kick = chute segundo o google tradutor 
	
	for low <= high {
		
		middle := (low + high) / 2
		kick := slice[middle]

		if(kick == value){
			return true, middle
		}

		if(kick > value){
			high = middle - 1
		}else{
			low = middle + 1
		}
	}

	return false, low
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Split(scanner.Text(), " ")
	slice := []int{}
	for _, elem := range parts[1 : len(parts)-1] {
		value, _ := strconv.Atoi(elem)
		slice = append(slice, value)
	}
	scanner.Scan()
	value, _ := strconv.Atoi(scanner.Text())
	found, result := BetterSearch(slice, value)
	if found {
		fmt.Println("V", result)
	} else {
		fmt.Println("F", result)
	}
}
