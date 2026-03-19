package main
import (
    "fmt"
    "math"
)

func main() {
    
    var numeroDeCompetidores int
    fmt.Scan(&numeroDeCompetidores)

    var indiceDoCompetidorVencedor = -1
    var resultadoDoValorAbsoluto = 0

    for i := 0; i < numeroDeCompetidores; i++{
        var a,b float64
        fmt.Scan(&a,&b)
        if(a < 10 || b < 10){
            continue
        }else{{
           var valorAbsoluto = math.Abs(a - b)
           if(valorAbsoluto >= float64(resultadoDoValorAbsoluto)){
                indiceDoCompetidorVencedor = i
           }
        }

    }
}

    if(indiceDoCompetidorVencedor == -1){
        fmt.Println("sem ganhador")
    }else{
        fmt.Println(indiceDoCompetidorVencedor)
    }

}