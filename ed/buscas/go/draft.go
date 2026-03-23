package main

import (
	"fmt"
	// "slices"
)
func main() {

    //----------------------------declaração de variaveis----------------------------
    
    var qtdConsultas, qtdBuscas int

    fmt.Scan(&qtdConsultas)

    //----------------------------inserção de elementos na lista de Consultas----------------------------
    listaDeConsultas := make([]string, qtdConsultas)

    for i :=0; i < qtdConsultas; i++{
        fmt.Scan(&listaDeConsultas[i])
    }

    // fmt.Println(listaDeConsultas)


  ////----------------------------inserção de elementos no mapa de Consultas----------------------------


    fmt.Scan(&qtdBuscas)

    listaDeBuscas := make([]string, qtdBuscas)
    
    mapaDeConsultas := make(map[string]int)

    for i :=0; i < qtdBuscas; i++{
        fmt.Scan(&listaDeBuscas[i])
        mapaDeConsultas[listaDeBuscas[i]] = 0
    }

    // fmt.Println(mapaDeConsultas)

    for i := 0; i < qtdConsultas; i++ {
        _, existe := mapaDeConsultas[listaDeConsultas[i]]
        if(existe){
            mapaDeConsultas[listaDeConsultas[i]]++
        }
    }

    // fmt.Println(mapaDeConsultas)

    //------------------------imprimindo

    for i := 0; i < len(listaDeBuscas); i++ {
         if(i == len(listaDeBuscas) - 1){
             fmt.Printf("%v\n", mapaDeConsultas[listaDeBuscas[i]] ) 
        }else{
            fmt.Printf("%v ", mapaDeConsultas[listaDeBuscas[i]] )
        }
    }

    // fmt.Println(len(listaDeBuscas))

}