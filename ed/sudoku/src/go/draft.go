package main
import (
    "fmt"
    "math"
)
    


func estaNaLinha(tabuleiro[][] int,l int, element int) bool{

    for c := 0; c < len(tabuleiro); c++ {
        if(tabuleiro[l][c] == element){
            return true
        }
    }
    
    return false
}

func estaNaColuna(tabuleiro[][] int,c int, element int) bool{

    for l := 0; l < len(tabuleiro); l++ {
        if(tabuleiro[l][c] == element){
            return true
        }
    }
    
    return false
}

func estaNoQuadrante(tabuleiro[][] int, lin int, col int, element int) bool {

    dim := len(tabuleiro)
    lado := int(math.Sqrt(float64(dim)))
    StartL := ( lin / lado) * lado
    StartC := ( col / lado) * lado

    for i := 0; i < int(lado); i++ {
        for j := 0; j < int(lado); j++ {
            if(tabuleiro[StartL+i][StartC+j] == element){
                return true
            }
        }
    }

    return false
}

func resolver(tabuleiro[][]int, index int) bool {
    /*Obs:
        index = a quantidade de casas na qual eu já passei
        nl = a quantidade de linhas do meu tabuleiro
    */

    nl := len(tabuleiro)
    //condição de parada
    if(index == nl*nl){
        return true
    }

    //linhas e colunas
    l := index / nl
    c := index % nl

    if(tabuleiro[l][c] != 0) {
        return resolver(tabuleiro, index+1)
    }

    //backtracking

    for i := 1; i <= nl; i++ {
        //checar a linha,coluna e quadrante.
        element := i
        if(!estaNaLinha(tabuleiro, l, element) && !estaNaColuna(tabuleiro, c, element) && !estaNoQuadrante(tabuleiro, l, c, element)){
            tabuleiro[l][c] = element
            if (resolver(tabuleiro, index+1)) {
                return true
            }
            tabuleiro[l][c] = 0
        }
        
    }

    return false

}

func main() {
    
    var numberOfColAndLines int
    fmt.Scanln(&numberOfColAndLines)

    tabuleiro := make([][]int, numberOfColAndLines)

    for i := 0; i < numberOfColAndLines; i++ {

        //add linha
        tabuleiro[i] = make([]int, numberOfColAndLines)

        //pegando linha toda
        var row string
        fmt.Scan(&row)      

        for j := 0; j < numberOfColAndLines; j++ {  
            if(row[j] == '.'){
                tabuleiro[i][j] = 0
            }else{
                tabuleiro[i][j] = int(row[j] - '0')
            }

        }
    }

    resolver(tabuleiro, 0)


// fmt.Println("-----------------------")    
    for i := 0; i < numberOfColAndLines; i++ {
        for j := 0; j < numberOfColAndLines; j++ {  
            fmt.Print(tabuleiro[i][j])        
        }
        fmt.Println()
    }

}