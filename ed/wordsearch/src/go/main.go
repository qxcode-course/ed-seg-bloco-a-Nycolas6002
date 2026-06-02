package main

import (
	"bufio"
	"fmt"
	"os"
)

//faço todas as validações
func backtracking(grid [][]byte,word string,row int,col int ,index int) bool{
		
	if(row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0])){
		return false
	}

	if(grid[row][col] != word[index]){
		return false
	}

	if(grid[row][col]) == '#'{
		return false
	}

	if(index == len(word)-1){
		return true
	}

	temp := grid[row][col]
	grid[row][col] = '#'

	found := backtracking(grid,word,row-1,col,index+1) ||
	backtracking(grid,word,row+1,col,index+1) ||
	backtracking(grid,word,row,col+1,index+1) ||
	backtracking(grid,word,row,col-1,index+1)

	grid[row][col] =  temp

	return found

}

// Não mude a assinatura desta função, ela é a função chamada pelo LeetCode
func exist(grid [][]byte, word string) bool {

     for row := 0; row < len(grid);row++{
        for col := 0; col < len(grid[row]);col++{
			if(backtracking(grid, word, row, col, 0)){
				return true
			}
        }
    }

	return false

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var word string
	fmt.Sscanf(scanner.Text(), "%s", &word)
	grid := make([][]byte, 0)
	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
	}
	if exist(grid, word) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
