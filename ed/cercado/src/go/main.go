package main

import (
	"bufio"
	"fmt"
	"os"
)

func dfs(board [][]byte, i int, j int){

	rows := len(board)
	cols := len(board[0])

	if(i < 0 || j < 0 || i >= rows || j >= cols){
		return
	}

	if(board[i][j] == 'X' || board[i][j] == 'T'){
		return
	}

	board[i][j] = 'T'

	dfs(board, i+1,j)//down
	dfs(board, i-1,j)//up
	dfs(board, i,j+1)//right
	dfs(board, i,j-1)//left
}

// NÃO ALTERE A ASSINATURA DA FUNÇÃO solve
func solve(board [][]byte) {

	rows := len(board)	
	cols := len(board[0])

	/*percorrendo as bordas do quadrado*/

	/*parte de cima*/
	for j := 0; j < cols;j++{
		if(board[0][j] == 'O'){
			dfs(board, 0 , j)
		}
	}

	/*lado esquerda*/
	for i := 0; i < rows;i++{
		if(board[i][0] == 'O'){
			dfs(board, i , 0)
		}
	}

	/*lado direito*/
	for i := 0; i < rows;i++{
		if(board[i][cols - 1] == 'O'){
			dfs(board, i , cols - 1)
		}
	}

	/*parte de baixo*/
	for j := 0; j < cols; j++{
		if(board[rows-1][j] == 'O'){
			dfs(board, rows - 1, j)
		}
	}

	for i := 0; i < rows;i++{
		for j := 0; j < cols; j++{
			if(board[i][j] == 'T'){
				board[i][j] = 'O'
			}else{
				board[i][j] = 'X'
			}

		}
	}


	// return board

	// _ = board
}

// NÃO ALTERE A MAIN
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var nrows, ncols int
	fmt.Sscanf(scanner.Text(), "%d %d", &nrows, &ncols)
	board := make([][]byte, nrows)
	for i := 0; i < nrows; i++ {
		scanner.Scan()
		board[i] = []byte(scanner.Text())
	}
	solve(board)
	for _, row := range board {
		fmt.Println(string(row))
	}
}
