package main
import "fmt"

func formasDeSubir(lista []int,n int) int{

    //cassos base
    if(n == 1){
        // lista[0] = 1
        return 1
    }

    if(n == 2){
        // lista[1] = 1
        return 1
    }

    if(n == 3){
        // lista[2] = 2
        return 2
    }

    if(lista[n] !=  0){
        return lista[n]
    }

    resultado := formasDeSubir(lista, n-1) + formasDeSubir(lista, n-3)

    lista[n] = resultado

    return resultado

}

func main() {
    var qtdDegraus int
    fmt.Scan(&qtdDegraus)
    listaResultados := make([]int,qtdDegraus+1)
    // fmt.Println(listaResultados)
    fmt.Println(formasDeSubir(listaResultados,qtdDegraus))
}
