package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func dfs(image [][]int, sr int, sc int, color int, originColor int){

    rows := len(image)
    cols := len(image[0])

    if(sr < 0 || sc < 0 || sr >= rows || sc >= cols){
		return
	}

	if(image[sr][sc] != originColor || image[sr][sc] == color){
		return
	}

	image[sr][sc] = color

    dfs(image , sr + 1, sc, color, originColor)//down
    dfs(image , sr-1, sc, color, originColor)//up
    dfs(image , sr, sc+1, color, originColor)//right
    dfs(image , sr, sc-1, color, originColor)//left

}

// Não modifique a assinatura da função floodFill
func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	

	originColor := image[sr][sc]
    dfs(image, sr, sc, color, originColor )

    return image

}

// Não modifique a função main
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	parts := strings.Fields(line)
	nl, _ := strconv.Atoi(parts[0])
	nc, _ := strconv.Atoi(parts[1])

	image := make([][]int, nl)
	for i := 0; i < nl; i++ {
		scanner.Scan()
		rowStr := strings.Fields(scanner.Text())
		row := make([]int, nc)
		for j := 0; j < nc; j++ {
			row[j], _ = strconv.Atoi(rowStr[j])
		}
		image[i] = row
	}

	// Lê sr, sc e color
	scanner.Scan()
	params := strings.Fields(scanner.Text())
	sr, _ := strconv.Atoi(params[0])
	sc, _ := strconv.Atoi(params[1])
	color, _ := strconv.Atoi(params[2])

	result := floodFill(image, sr, sc, color)

	// Imprime a matriz resultante
	for _, row := range result {
		for j, val := range row {
			if j > 0 {
				fmt.Print(" ")
			}
			fmt.Print(val)
		}
		fmt.Println()
	}
}
