package main
import "fmt"

type info struct {
    gasolina int
    distancia int
}

func main() {

    var n int
    fmt.Scan(&n)

    bombas := make([]info, n)

    for i := 0; i < n; i++ {
        fmt.Scan(&bombas[i].gasolina, &bombas[i].distancia)
    }

    indiceProcurado := -1

    //iniciando em cada bomba
    for inicio := 0; inicio < n; inicio++ {

        tanque := 0
        consegueChegar := true

        //percorrendo todas as bombas
        for passo := 0; passo < n; passo++ {

            passoAtual := (inicio + passo) % n
            // fmt.Println("passo",passoAtual)
            // fmt.Println("gasolina",bombas[passoAtual].gasolina)
            // fmt.Println("distancia",bombas[passoAtual].distancia)
            // fmt.Println("------------------")

            tanque += bombas[passoAtual].gasolina

            tanque -= bombas[passoAtual].distancia

            if(tanque < 0){
                consegueChegar = false
                break
            }
        }

        if(consegueChegar){
            indiceProcurado = inicio
        }

    }

    fmt.Println(indiceProcurado)
    

}