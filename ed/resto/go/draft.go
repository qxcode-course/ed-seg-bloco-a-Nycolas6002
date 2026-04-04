package main
import "fmt"

// func printar(numero int, resto int ){
//     fmt.Println(numero, resto)
// }

func dividir(numero int){

    //caso base
    if(numero == 0){
        return
    }

    novoNumero := numero / 2
    resto := numero % 2

    dividir(novoNumero)
    fmt.Println(novoNumero, resto)
    
}

func main() {
    numero := 0

    fmt.Scan(&numero)
    dividir(numero)
    // numeroDivido := numero/2
    // dividir(numeroDivido, resto)

}
