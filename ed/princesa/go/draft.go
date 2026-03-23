package main
import "fmt"

func imprimir (sliceDeNum []int, pos int){
     fmt.Print("[ ")
    for i,valor := range sliceDeNum{
        
       
        if(pos == i){
            fmt.Printf("%v> ", valor)
        }else{
            fmt.Printf("%v ", valor)
        }
    }
    fmt.Println("]")
}

func main() {
    var qtdNumeros,numEscolhido int

    fmt.Scan(&qtdNumeros,&numEscolhido)

    listaDeNumeros := make([]int, qtdNumeros)

    for i := 0; i < qtdNumeros; i++ {
        listaDeNumeros[i] = i+1
    }

    posAtual := numEscolhido - 1 

    for len(listaDeNumeros) > 0 {
           imprimir(listaDeNumeros, posAtual)

           if(len(listaDeNumeros) == 1){
            break
           }
            
           morto := (posAtual + 1) % len(listaDeNumeros)
           
           listaDeNumeros = append(listaDeNumeros[:morto], listaDeNumeros[morto+1:]...)

           posAtual = morto % len(listaDeNumeros)
        }

}