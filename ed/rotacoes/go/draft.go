package main
import "fmt"
func main() {
    var tamanhoLista, elementos int

    fmt.Scan(&tamanhoLista,&elementos)

    listaDeNumeros := make([]int, tamanhoLista)

    for i := 0; i < len(listaDeNumeros); i++ {
        // var valor int
        fmt.Scan(&listaDeNumeros[i])
        // listaDenu
    }

    listaFinal := make([]int, tamanhoLista)

   
    for i := 0; i < len(listaDeNumeros); i++ {
        pos := (i + elementos) % len(listaDeNumeros)
        listaFinal[pos] = listaDeNumeros[i]
    }

    //imprimindo
    fmt.Print("[ ")
   
    for i := 0; i < len(listaFinal); i++ {
        fmt.Printf("%v ",listaFinal[i])
    }
    fmt.Println("]")
    // fmt.Println(listaFinal)
}
