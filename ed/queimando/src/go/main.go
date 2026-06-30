package main

import (
	"bufio"
	"fmt"
	// v2 "math/rand/v2"
	"os"
)

type Pos struct{
	l int
	c int
}

func burnTrees(grid [][]rune, l, c int) {
	
	stack := []Pos{{l,c}}

	for len(stack) > 0 {
		
		atual := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if(atual.l < 0 || atual.c < 0 || atual.l >= len(grid) || atual.c >= len(grid[0])){
			continue
		}
		
		if(grid[atual.l][atual.c] != '#'){
			continue
		}

		grid[atual.l][atual.c] = 'o'

		stack = append(stack,
			Pos{atual.l - 1, atual.c},
			Pos{atual.l + 1, atual.c},
			Pos{atual.l, atual.c - 1},
			Pos{atual.l, atual.c + 1},
			)

	}
	
	
	
	
	// stack := NewStack[Pos]()
	// _ , _ , _ = mat, l, c

	// Essa função deve usar uma list como pilha
	// e marcar as árvores na matriz como queimados
	// Uma sugestão de como fazer isso é:
	// - adicionar a primeira posição na pilha
	// - enquanto a pilha não estiver vazia:
	//   - retirar o elemento do topo
	//   - se puder ser queimado, queime e adicione seus vizinhos à pilha

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc, lfire, cfire int
	fmt.Sscanf(line, "%d %d %d %d", &nl, &nc, &lfire, &cfire)

	grid := make([][]rune, 0, nl)
	for range nl {
		scanner.Scan()
		line := []rune(scanner.Text())
		grid = append(grid, line)
	}
	burnTrees(grid, lfire, cfire)
	showGrid(grid)
}

func showGrid(mat [][]rune) {
	for _, linha := range mat {
		fmt.Println(string(linha))
	}
}
