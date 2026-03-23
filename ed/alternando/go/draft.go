package main
import "fmt"

func imprimir(listaNumeros []int, posAtual int){

    fmt.Print("[ ")
    for i,valor := range listaNumeros {
        
        if(valor < 0){
            if(posAtual == i ){
                    fmt.Printf("<%v ", valor)
            }else{
                    fmt.Printf("%v ", valor)
            }
        }else{
            if(posAtual == i ){
                fmt.Printf("%v> ", valor)
            }else{
                fmt.Printf("%v ", valor)
            }    
        }       
    
    }
    fmt.Printf("]\n")


}

func main() {

    var qtdNumeros, posicao, sinal int

    fmt.Scan(&qtdNumeros, &posicao, &sinal)

    listaNumeros := make([]int, qtdNumeros)

    for i := 0; i < len(listaNumeros); i++ {

        if(sinal == -1){
            if((i+1) % 2 == 1){
                listaNumeros[i] = -(i+1)
            }else{
                listaNumeros[i] = i+1
            }
        }else{
            if((i+1) % 2 == 0){
                listaNumeros[i] = -(i+1)
            }else{
                listaNumeros[i] = i+1
            }
        }
    }

    posicaoAtual := posicao - 1
    // morto := -1
    // fmt.Println(listaNumeros)
    // imprimir(listaNumeros, 1)

    for len(listaNumeros) > 0 {
        imprimir(listaNumeros, posicaoAtual)


            if(len(listaNumeros) == 1){
                break
            }

            if(listaNumeros[posicaoAtual] < 0){
                //morto
                morto := (posicaoAtual - 1 + len(listaNumeros)) % len(listaNumeros)

                //slice
                listaNumeros = append(listaNumeros[:morto],listaNumeros[morto+1:]...)
                
                //atualizar posicão
                posicaoAtual = (morto - 1 + len(listaNumeros)) % len(listaNumeros)

            }else{
                //morto
                morto := (posicaoAtual + 1) % len(listaNumeros)
                
                //slice
                listaNumeros = append(listaNumeros[:morto],listaNumeros[morto+1:]...)

                //atualizar posicão
                posicaoAtual = morto % len(listaNumeros)
            }

    }


}



//solução da V1


// func imprimir (slice []int, pos int){

//     fmt.Print("[ ")
//     for i,valor := range slice {
            
//         if(pos == i ){
//             fmt.Printf("%v> ", valor)
//         }else{
//             fmt.Printf("%v ", valor)
//         }        
//     }
//     fmt.Printf("]\n")


// }


    // var qtdPessoas, posicao, sinal int

    // fmt.Scan(&qtdPessoas, &posicao, &sinal)

    // listaParticipantes := make([]int, qtdPessoas)

    // for i := 0; i < qtdPessoas; i++ {
    //     listaParticipantes[i] = i+1
    // }

    // posAtual := posicao -1

    // for len(listaParticipantes) > 0 {
    //     imprimir(listaParticipantes, posAtual)

    //     morto := (posAtual + 1) % len(listaParticipantes)

    //     if(len(listaParticipantes) == 1){
    //         break
    //     }

    //     listaParticipantes = append(listaParticipantes[:morto], listaParticipantes[morto+1:]...)

    //     posAtual = morto % len(listaParticipantes)
    // }

    // // fmt.Println(listaParticipantes)
