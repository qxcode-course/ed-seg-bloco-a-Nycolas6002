package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	row    int
	column int
}

func search(grid [][]rune, pos, endPos Pos) bool {
	row, column := pos.row, pos.column

	// 1. fora da matriz
	if row < 0 || row >= len(grid) || column < 0 || column >= len(grid[0]) {
		return false
	}

	// 2. parede ou já visitado
	if grid[row][column] != ' ' {
		return false
	}
	
	grid[row][column] = '.'

	// 3. chegou no destino
	if row == endPos.row && column == endPos.column {
		return true
	}

	// 4. marca como visitado temporariamente

	// 5. tenta os 4 lados
	if search(grid, Pos{row - 1, column}, endPos) ||
		search(grid, Pos{row + 1, column}, endPos) ||
		search(grid, Pos{row, column - 1}, endPos) ||
		search(grid, Pos{row, column + 1}, endPos) {

		return true
	}

	// 6. backtracking (desfaz)
	grid[row][column] = ' '
	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nl_nc := scanner.Text()
	var nl, nc int
	fmt.Sscanf(nl_nc, "%d %d", &nl, &nc)
	grid := make([][]rune, nl)

	// Lê a gridriz
	for i := range nl {
		scanner.Scan()
		grid[i] = []rune(scanner.Text())
	}

	// Procura posições de início e endPos e conserta para _
	var startPos, endPos Pos
	for l := range nl {
		for c := range nc {
			if grid[l][c] == 'I' {
				grid[l][c] = ' '
				startPos = Pos{l, c}
			}
			if grid[l][c] == 'F' {
				grid[l][c] = ' '
				endPos = Pos{l, c}
			}
		}
	}

	search(grid, startPos, endPos)

	// Imprime o labirinto final
	for _, line := range grid {
		fmt.Println(string(line))
	}
}
