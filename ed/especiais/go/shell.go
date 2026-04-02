package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sort"
)

type Pair struct {
	One int
	Two int
}
/*
func abs(valor int) int{
    if(valor > 0){
        return valor
    }

    return -valor
}

func occurr(vet []int) []Pair {

    mapa := make(map[int]Pair, len(vet))
	contagem := []Pair{}
	// ordem := []int{}
	
    for _,valor := range vet{
		elemento := abs(valor)
		
        _,existe := mapa[elemento]
        info := mapa[elemento]
        
        if(existe){
			info.Two++
			}else{
				info.One = elemento
				info.Two = 1   
				// ordem = append(ordem, elemento) 
				contagem = append(contagem, Pair{elemento, 1})
			}
			mapa[elemento] = info
			
		}
		
	// 
	
	// for _,valor := range ordem{
	// 	contagem = append(contagem, mapa[valor])
	// }

	sort.Ints(contagem)
    
	return contagem

    // for _, valor := range mapa{
    //     contagem = append(contagem, valor)
    // }

	//teste 01

	// for i := 0; i < len(mapa);i++{
	// 	//vetor
	// 	for chave,valor := range mapa{
	// 		//mapa
	// 		if(abs(vet[i]) == chave){
	// 			//add contagem na pos i
	// 			contagem = append(contagem, valor)
	// 		}
	// 	}
	// }

	// teste 2
	// for i := 0; i < len(vet);i++{
		
	// 	for chave, valor := range mapa{
	// 		if(vet[i] == chave){
	// 			contagem = append(contagem, valor)
	// 		}
	// 	}
	// }


}*/

func abs(valor int) int {
    if valor >= 0 { // corrigido aqui
        return valor
    }
    return -valor
}

func occurr(vet []int) []Pair {

	mapa := make(map[int]Pair)
    chave := []int{}
	
    for _, valor := range vet {
		elemento := abs(valor)
		
        info, existe := mapa[elemento] // 1 acesso ao map
		
        if existe {
			info.Two++
			} else {
				info = Pair{One: elemento, Two: 1}
				chave = append(chave, elemento)
			}
			
			mapa[elemento] = info
		}
		
	sort.Ints(chave)
    contagem := []Pair{}

    for _, valor := range chave{
        contagem = append(contagem, mapa[valor])
    }

    return contagem
}

func teams(vet []int) []Pair {


	if(len(vet) == 0){	
		return []Pair{}
	}

	times := []Pair{}
	contador := 1
	atual  := vet[0]

	for i := 1 ; i < len(vet); i++{

			// 0       1
		if(vet[i] == atual){
			contador++
		}else{
			times = append(times, Pair{One:atual, Two:contador})
			atual = vet[i]
			contador = 1
		}
	}

	times = append(times, Pair{One:atual, Two:contador})
	
	return times

}

func mnext(vet []int) []int {
	//revidsar depois
	listaDePosicoes := make([]int, len(vet))

	if(len(vet) == 1){
		return listaDePosicoes
	}

	for i := 1; i < len(vet); i++{

		if(vet[i-1] < 0 && vet[i] > 0){
			listaDePosicoes[i] = 1
		}
		
		if(i < len(vet)-1){
			if( vet[i] > 0 && vet[i+1] < 0){
				listaDePosicoes[i] = 1
				
			}

		}
	}

	return listaDePosicoes

}

func alone(vet []int) []int {


	listaDePosicoes := make([]int, len(vet))

	if(len(vet) == 1 && vet[0] > 0){
		
		return []int{0}
	}

		for i := 1; i < len(vet); i++{

		if(vet[i-1] > 0 && vet[i] > 0){
			listaDePosicoes[i] = 1
		}
		
		if(i < len(vet)-1){
			if( vet[i] > 0 && vet[i+1] > 0){
				listaDePosicoes[i] = 1
				
			}

		}
	}

	return listaDePosicoes

	// _ = vet
	// return nil
}

func couple(vet []int) int {

	if(len(vet) == 1){
		return 0
	}

	// contador := 0

	return 0
}

func hasSubseq(vet []int, seq []int, pos int) bool {
	_ = vet
	_ = seq
	_ = pos
	return false
}

func subseq(vet []int, seq []int) int {
	_ = vet
	_ = seq
	return -1
}

func erase(vet []int, posList []int) []int {

	posicoes := make([]int, len(vet))

    for _, valor := range posList{
      posicoes[valor] = 1
	}

    novaLista := []int{}

    for i := 0; i < len(posicoes); i++{
        if(posicoes[i] == 0){
            novaLista = append(novaLista, vet[i])
        }
    }

	return novaLista

}

func clear(vet []int, value int) []int {

	for i := 0; i < len(vet); i++{
		if(vet[i] == value){
			vet = append(vet[:i], vet[i+1:]...)
			i--
		}
	}

	return vet
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "occurr":
			printSlice(occurr(str2vet(args[1])))
		case "teams":
			printSlice(teams(str2vet(args[1])))
		case "mnext":
			printSlice(mnext(str2vet(args[1])))
		case "alone":
			printSlice(alone(str2vet(args[1])))
		case "erase":
			printSlice(erase(str2vet(args[1]), str2vet(args[2])))
		case "clear":
			val, _ := strconv.Atoi(args[2])
			printSlice(clear(str2vet(args[1]), val))
		case "subseq":
			fmt.Println(subseq(str2vet(args[1]), str2vet(args[2])))
		case "couple":
			fmt.Println(couple(str2vet(args[1])))
		case "end":
			return
		default:
			fmt.Println("Invalid command")
		}
	}
}

// Funções auxiliares

func str2vet(str string) []int {
	if str == "[]" {
		return nil
	}
	str = str[1 : len(str)-1]
	parts := strings.Split(str, ",")
	var vet []int
	for _, part := range parts {
		num, _ := strconv.Atoi(strings.TrimSpace(part))
		vet = append(vet, num)
	}
	return vet
}

func printSlice[T any](vet []T) {
	fmt.Print("[")
	for i, x := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(x)
	}
	fmt.Println("]")
}

func (p Pair) String() string {
	return fmt.Sprintf("(%v, %v)", p.One, p.Two)
}
