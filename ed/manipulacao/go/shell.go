package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getMen(vet []int) []int {

	var listaHomens []int
	
	for i :=0 ; i < len(vet); i++{
		if vet[i] > 0 {
			listaHomens = append(listaHomens, vet[i])
		}
	}

	return listaHomens

}

func getCalmWomen(vet []int) []int {

	var listaMulheresEstressadas []int

	for i := 0; i < len(vet); i++{
		if(vet[i] > -10 && vet[i] < 0){
			listaMulheresEstressadas = append(listaMulheresEstressadas,vet[i])
		}
	}

	return listaMulheresEstressadas

}

func sortVet(vet []int) []int {

	listaOrdenada := make([]int, len(vet))
	copy(listaOrdenada,vet)
	n := len(listaOrdenada)

	for i:= 0; i < n-1; i++{
		for j:= 0; j < n-1-i; j++{
			if(listaOrdenada[j] > listaOrdenada[j+1]){
				temp := listaOrdenada[j]
				listaOrdenada[j] = listaOrdenada[j+1]
				listaOrdenada[j+1] = temp
			}
		}
	}

	return listaOrdenada

}

func sortStress(vet []int) []int {
	
	listaEstressados:= make([]int, len(vet))
	copy(listaEstressados,vet)
	n := len(listaEstressados)

	 abs := func(x int) int{
		if x < 0{
			return -x
		}

		return x
	}

	for i:= 0; i < n-1; i++{
		for j:= 0; j < n-1-i; j++{
			if(abs(listaEstressados[j]) > abs(listaEstressados[j+1])){
				temp := listaEstressados[j]
				listaEstressados[j] = listaEstressados[j+1]
				listaEstressados[j+1] = temp
			}
		}
	}

	return listaEstressados

}

func reverse(vet []int) []int {

	var listaReversa []int

	for i := len(vet) -1; i >= 0 ; i-- {
		listaReversa = append(listaReversa, vet[i])
	}

	return listaReversa

}

func unique(vet []int) []int {

	mapaUnicos := make(map[int]int, len(vet))
	var listaUnicos[]int

	for _,elementoVet := range vet{
		_, existe := mapaUnicos[elementoVet]
		if(!existe){
			mapaUnicos[elementoVet] = 1
		}else{
			mapaUnicos[elementoVet]++
		}
	}

	for chave,_ := range mapaUnicos{
		listaUnicos = append(listaUnicos, chave)
	}

	return listaUnicos

}

func repeated(vet []int) []int {

	mapaRepetidos := make(map[int]int, len(vet))
	var listaRepetidos []int

	for _,elementoVet := range vet{
		_, existe := mapaRepetidos[elementoVet]
		if(!existe){
			mapaRepetidos[elementoVet] = 1
		}else{
			mapaRepetidos[elementoVet]++
		}
	}

	for chave, valor := range mapaRepetidos{
		if(valor >= 2){
			for i :=0; i < valor-1 ;i++{
				listaRepetidos = append(listaRepetidos, chave)
			}
		}
	}

	tamanhoLista := len(listaRepetidos)

	for i:=0; i < tamanhoLista-1; i++{
		for j :=0 ; j < tamanhoLista-1-i; j++ {
			if(listaRepetidos[j] > listaRepetidos[j+1]){
				temp := listaRepetidos[j]
				listaRepetidos[j] = listaRepetidos[j+1]
				listaRepetidos[j+1] = temp
			}
		}
	}

	return listaRepetidos
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if !scanner.Scan() {
			break
		}
		fmt.Print("$")
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "get_men":
			printVec(getMen(str2vet(args[1])))
		case "get_calm_women":
			printVec(getCalmWomen(str2vet(args[1])))
		case "sort":
			printVec(sortVet(str2vet(args[1])))
		case "sort_stress":
			printVec(sortStress(str2vet(args[1])))
		case "reverse":
			array := str2vet(args[1])
			other := reverse(array)
			printVec(array)
			printVec(other)
		case "unique":
			printVec(unique(str2vet(args[1])))
		case "repeated":
			printVec(repeated(str2vet(args[1])))
		case "end":
			return
		}
	}
}

func printVec(vet []int) {
	fmt.Print("[")
	for i, val := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(val)
	}
	fmt.Println("]")
}

func str2vet(s string) []int {
	if s == "[]" {
		return nil
	}
	s = s[1 : len(s)-1]
	parts := strings.Split(s, ",")
	var vet []int
	for _, part := range parts {
		n, _ := strconv.Atoi(part)
		vet = append(vet, n)
	}
	return vet
}

