package main

import (
	"bufio"
	"fmt"
	"os"
)

/*a função dfs vai servir apenas para eu percorrer todos os caminhos da ilha. E então eu irei "apagar onde tem 1"*/
func dfs(grid [][]byte, i int, j int){

	rows := len(grid)
	cols := len(grid[0])

	if(i < 0 || j < 0 || i >= rows || j >= cols ){
		return
	}

	if(grid[i][j] == '0'){
		return
	}

	grid[i][j] = '0'

	/*walk to right,left,up,down*/
	dfs(grid, i+1, j) //down
	dfs(grid, i-1, j) //up
	dfs(grid, i, j+1) //right
	dfs(grid, i, j-1) //left

	//condição de parada
	//passo recursivo
	//retorno

}

// Não modifique a assinatura da função numIslands
// Ela é a função que será chamada no LeetCode para resolver o problema
func numIslands(grid [][]byte) int {


	rows := len(grid)
	cols := len(grid[0])
	count := 0

	for i:= 0; i < rows ;i++{
		for j:= 0; j < cols ; j++{

			if(grid[i][j] == '1'){
				count++
				dfs(grid, i, j)
			}

		}
	}

	/*
		percorrer linha e coluna
			se achei 1
			aumento o count em 1 
			faço a busca em profundidade
	*/
	//
	// _ := grid
	// return 0

	return count
}

// Não modifique a função main
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc int
	fmt.Sscanf(line, "%d %d", &nl, &nc)
	grid := make([][]byte, nl)
	for i := 0; i < nl; i++ {
		scanner.Scan()
		grid[i] = []byte(scanner.Text())
	}
	result := numIslands(grid)
	fmt.Println(result)
}
