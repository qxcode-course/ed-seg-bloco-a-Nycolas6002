package main
import (
    "fmt"
    "strings"
)


func diagonal(s string, k int){
    if len(s) == 0 {
        return
    }
    // jefferson 9

    // Println
    diagonal(s[1:],k-1)
    fmt.Print(strings.Repeat(" ", k))
    fmt.Println(string(s[0]))

}

func main() {
    var palavra string
    fmt.Scan(&palavra)
    qtdCaracteres := len(palavra)
    // fmt.Println(diagonal(palavra, qtdCaracteres))
    diagonal(palavra, qtdCaracteres)
}
