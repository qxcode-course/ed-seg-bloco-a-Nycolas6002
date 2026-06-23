package main

import (
	"fmt"
)
func main() {
    
    var expressao string
    var pilha []string
    
    fmt.Scan(&expressao)
    // fmt.Println(expressao)

    // fmt.Println(string(expressao[0]))

    for _, valor := range expressao{
        //push
        caractere := string(valor)
        if(caractere == "(" || caractere == "["){
            pilha = append(pilha, caractere)
        }

        //pop

        if(caractere == "]" || caractere == ")"){

            if(len(pilha) == 0){
                fmt.Println("nao balanceado")
                return
            }

            //topo
            topo := pilha[len(pilha)-1]
            //remove
            pilha = pilha[:len(pilha)-1]
            
            if(caractere == ")" && topo != "(" || caractere == "]" && topo != "["){
                fmt.Println("nao balanceado")
               return
            }

        }

    }

    if(len(pilha) > 0){
        fmt.Println("nao balanceado")
    }else{
        fmt.Println("balanceado")
    }
    


}