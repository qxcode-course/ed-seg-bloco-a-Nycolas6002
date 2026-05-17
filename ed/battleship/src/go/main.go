package main

import (
	"bufio"
	"fmt"
	"os"
)

// Função que será chamada no LeetCode
func countBattleships(board [][]byte) int {

	rows := len(board)
	cols := len(board[0])
	ships := 0
	

    for i := 0; i < rows; i++ {

        for j := 0; j < cols; j++ {

			if(board[i][j] ==  'X'){
				if( i > 0 && board[i-1][j] ==  'X'){
					continue
				}else if(j > 0 && board[i][j-1] ==  'X'){
					continue
				}else{
					ships++
				}
			}

        }
    }

	return ships

}

// Não modifique a função main
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	line := scanner.Text()

	var nl, nc int
	fmt.Sscanf(line, "%d %d", &nl, &nc)

	board := make([][]byte, nl)

	for i := 0; i < nl; i++ {
		scanner.Scan()
		board[i] = []byte(scanner.Text())
	}

	result := countBattleships(board)

	fmt.Println(result)
}