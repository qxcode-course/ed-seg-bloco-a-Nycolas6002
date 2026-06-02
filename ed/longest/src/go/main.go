package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func dfs(matrix [][]int, i int, j int){

	rows := len(matrix)
    cols := len(matrix[0])

    // fora da matriz
    if i < 0 || j < 0 || i >= rows || j >= cols {
        return
    }
	
	if(){
		
	}

	
	dfs(matrix, i+1,j )//down
	dfs(matrix, i-1,j )//up
	dfs(matrix, i,j+1 )//right
	dfs(matrix, i,j-1)//left

}

func longestIncreasingPath(matrix [][]int) int {

	rows := len(matrix)
	cols := len(matrix[0])

	count := 0

	for i := 0; i < rows; i++{
		for j := 0; j < cols; j++{
			count++
			dfs(matrix, i, j)
		}
	}

	return count

}

// Não modifique a função main
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	if !scanner.Scan() {
		return
	}
	parts := strings.Fields(scanner.Text())
	if len(parts) < 2 {
		return
	}
	nl, _ := strconv.Atoi(parts[0])
	nc, _ := strconv.Atoi(parts[1])

	matrix := make([][]int, nl)
	for i := 0; i < nl; i++ {
		if !scanner.Scan() {
			return
		}
		tokens := strings.Fields(scanner.Text())
		row := make([]int, nc)
		for j := 0; j < nc && j < len(tokens); j++ {
			v, _ := strconv.Atoi(tokens[j])
			row[j] = v
		}
		matrix[i] = row
	}

	fmt.Println(longestIncreasingPath(matrix))
}
