package main
import "fmt"
func main() {

    var qtdFigurasTotal,qtdFiguras int

    fmt.Scan(&qtdFiguras,&qtdFigurasTotal)

    listaFiguras := []int{}
    mapaFiguras := make(map[int]int)

    for i:=0; i < qtdFigurasTotal; i++{
        var elemento int
        fmt.Scan(&elemento)
        _, existe := mapaFiguras[elemento]
        if(existe){
            mapaFiguras[elemento]++
            // listaFiguras[i] = elemento
            listaFiguras = append(listaFiguras,elemento)
        }else{
            mapaFiguras[elemento] = 1
        }
    }
    

    if(len(listaFiguras) >= 1){
        for i:=0; i < len(listaFiguras); i++{
            if(i < len(listaFiguras)-1){
                fmt.Printf("%v ",listaFiguras[i])
            }else{
                fmt.Printf("%v\n",listaFiguras[i])
            }
        }

    }else{
        fmt.Println("N")
    }

    listaFaltantes := []int{}

    for i:=1; i <= qtdFiguras; i++{
        // valor, _ := 
        // fmt.Println(mapaFiguras[i])
        if(mapaFiguras[i] == 0){
            listaFaltantes = append(listaFaltantes, i)
        }
    }

    if(len(listaFaltantes) > 0){
        for i:=0; i < len(listaFaltantes); i++{
            if(i < len(listaFaltantes)-1){
                fmt.Printf("%v ",listaFaltantes[i])
            }else{
                fmt.Printf("%v\n",listaFaltantes[i])
            }
        }
        
    }else{
        fmt.Println("N")
    }


}
