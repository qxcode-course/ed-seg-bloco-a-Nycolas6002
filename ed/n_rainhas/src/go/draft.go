package main
import "fmt"

var totalQueens int

func isPositionValid(positions []int,row int,col int) bool{
    
    for r := 0; r < row; r++ {
        //checando coluna
        if(positions[r] == col){
            return false
        }

        //checando diagonal
        if(abs(r-row) == abs(positions[r]-col)){
            return false
        }
        
    }

    return true

}

func abs(x int) int{

    if(x < 0 ){
        return -x
    }
    return x

}

func totalOfQueens(numsQuens []int, row int, n int){
    //n reprensentando o qtd de rainhas e tamanho do tabuleiro
    if(row == n){
        totalQueens++
        return
    }

    for col := 0; col < n; col++ {

        if(isPositionValid(numsQuens,row,col)){

            numsQuens[row] = col

            totalOfQueens(numsQuens, row+1 , n)

        }
        
    }

}

func main() {
    var numsQuens int

    fmt.Scan(&numsQuens)

    position := make([]int, numsQuens)

    totalOfQueens(position, 0, numsQuens)

    fmt.Println(totalQueens)
}