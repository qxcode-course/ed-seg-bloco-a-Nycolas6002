package main
import (
    "fmt"
    // "strings"
    "sort"
)

func backtracking(caracteres []string, current []string, using []bool){

    if(len(current) == len(caracteres)){
        for _, value := range current {
            fmt.Print(value)
        }
        fmt.Println()

        return
    }

    for i := 0; i < len(caracteres); i++ {
        
        if(using[i]){
            continue
        }
        //Escolher
        using[i] = true
        current = append(current, caracteres[i])

        //Explorar
        backtracking(caracteres, current, using )

        //Voltar
        current = current[:len(current)-1]
        using[i] = false

    }

}

func main() {
    
    var word string

    fmt.Scan(&word)
    
    var caractere []string
    
    for _, value := range word {
        caractere = append(caractere, string(value))
    }
    sort.Strings(caractere)

    using := make([]bool, len(caractere))


    backtracking(caractere, []string{}, using )
}