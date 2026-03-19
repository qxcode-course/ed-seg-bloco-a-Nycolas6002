package main
import "fmt"
func main() {
    var qtdTotalFig, qtdCartasBaruel int

    fmt.Scan(&qtdTotalFig, &qtdCartasBaruel)
    
    albumBaruel := make([]int, qtdCartasBaruel)
    figurasRepetidas := make([]int, 0, qtdCartasBaruel)
    figurasUnicas := make(map[int]bool)

    for i := range qtdCartasBaruel{
        fmt.Scan(&albumBaruel[i])
    }

    //repetidos

    for _, figura := range albumBaruel{
        if(figurasUnicas[figura]){
            figurasRepetidas = append(figurasRepetidas, figura)
            
        }else{
            figurasUnicas[figura] = true
        }
    }

        if(len(figurasRepetidas) == 0){
            fmt.Println("N")
        }else{
            for i, valor := range figurasRepetidas{
                if(i < len(figurasRepetidas) -1 ){
                   fmt.Printf("%v ",valor)
                }else{
                    fmt.Println(valor)
                }
                
            }
        }


        //unicos
        figFaltantes := make([]int, 0 ,qtdTotalFig)

        for i := 1; i <= qtdTotalFig; i++ {
            
            // figuraFaltosa, existe := figurasUnicas[i]

            if(!figurasUnicas[i]){
                figFaltantes = append(figFaltantes, i)
            }

        }


        //imprimindo figuras faltantes

        if (len(figFaltantes) == 0) {
            fmt.Println("N")
        }else{
            for i, valor := range figFaltantes{
                if(i < len(figFaltantes) - 1){
                    fmt.Printf("%v ", valor)
                }else{
                    fmt.Println(valor)
                }
            }
        }

    
}





















/*


    listaDeCartas := make(map[int]int)

    for i := 0; i < qtdCartasBaruel; i++ {
        var carta int
        fmt.Scan(&carta)
        listaDeCartas[carta]++
        // if(listaDeCartas[carta] > 0){
        //     listaDeCartas[carta]++
        // }else{
        //     listaDeCartas[carta] = 1
        // }
    }

    var repetidas bool

    for i := 0; i < qtdCartasBaruel; i++ {
        
        if(){
            repetidas = true
        }
    }

    // cartasFaltosas := make([]int)

    // for i := 1; i <= qtdTotalFig; i++ {
    //     if(listaDeCartas[i] < 0){
    //         fmt.Println("N")
    //     }
    // }

    // fmt.Println(listaDeCartas)

*/