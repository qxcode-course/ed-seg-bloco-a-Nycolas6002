package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
    

func verificaPrecendencia(valor string) int{

    switch(valor){
        case "^":
            return 3
        case "*", "/":
            return 2
        case "+", "-":
            return 1
        default:
            return 0
    }

}

func verificaOperador(valor string) bool{
    return valor == "+" ||
            valor == "-" ||
            valor == "*" ||
            valor == "/" ||
            valor == "^"
}


func main() {
    // var expressao string

    // fmt.Scanln(&expressao)

    // fmt.Println(expressao)
    scanner := bufio.NewScanner(os.Stdin)
    var sliceCaracteres []string

    if(scanner.Scan()){
        sliceCaracteres = strings.Fields(scanner.Text())
    }

    var pilha []string 
    var saida []string

    for _, valor := range sliceCaracteres  {
        
        //determina se é um operador
        if(!verificaOperador(valor)){
            saida = append(saida, valor)
            continue
        }
        
        for len(pilha) > 0 && verificaPrecendencia(pilha[len(pilha)-1]) >= verificaPrecendencia(valor) {
            
            saida = append(saida, pilha[len(pilha)-1])
            pilha = pilha[:len(pilha)-1]
            
        }

        pilha = append(pilha, valor)
    }


    // fmt.Println(pilha[0])
    // fmt.Println(saida[0])

    for len(pilha) > 0{
        saida = append(saida, pilha[len(pilha)-1])
        pilha = pilha[:len(pilha)-1]
    }

    fmt.Println(strings.Join(saida, " "))
}