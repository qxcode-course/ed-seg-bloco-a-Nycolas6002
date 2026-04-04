package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


/*escrevendo como eu acho que deve ser o passo a passo para resolver o problema.

se o vetor não tiver nada eu retorno []

caso contrario eu tenho que usar a return tostr(vet[1;]) para ficar chamando sempre a função

aí eu tenho que tratar a impressão que no final tem que me exibir nesse formato ["numero", "numero2"]

tenho que usar a função printar que criei para fazer a formatação que eu quero passando como argumento o numero que estou percorrendo agora
em seguida converto ele para string e retorno ele para a função.

a partir daqui não sei muito bem como desenvolver mas valide esse meu raciocinio e diga se está coerente fale sobre ele, mas não me dê dicas de como resolver o problema


*/

func tostr(vet []int) string {

	if(len(vet) == 0){
		return "[]"
	}

	if(len(vet) == 1){
		return "[" + strconv.Itoa(vet[0]) + "]"
	}

	return "[" + strconv.Itoa(vet[0]) + ", "+ tostr(vet[1:])[1:]
	
}

func tostrrev(vet []int) string {

	
	if(len(vet) == 0){
		return "[]"
	}

	if(len(vet) == 1){
		return "[" + strconv.Itoa(vet[0]) + "]"
	}

	// "[" + strconv.Itoa(vet[0]) + ", "+ tostr(vet[1:])[1:]
	return "[" + strconv.Itoa(vet[len(vet)-1]) + ", "+ tostrrev(vet[:len(vet)-1])[1:]

	// _ = vet
}

// reverse: inverte os elementos do slice
func reverse(vet []int) {
	    if len(vet) <= 1 {
        return
    }

    // troca primeiro com último
    vet[0], vet[len(vet)-1] = vet[len(vet)-1], vet[0]

    // recursão no meio
    reverse(vet[1 : len(vet)-1])
	// _ = vet
}

// sum: soma dos elementos do slice
func sum(vet []int) int {

	if(len(vet) == 0){
		return 0
	}

	return vet[0] + sum(vet[1:]) 

	// _ = vet
	// return 0
}

// mult: produto dos elementos do slice
func mult(vet []int) int {

	if(len(vet) == 0){
		return 1
	}

	return vet[0] * mult(vet[1:]) 
	// _ = vet
	// return 0
}

// min: retorna o índice e valor do menor valor
// crie uma função recursiva interna do modelo
// var rec func(v []int) (int, int)
// para fazer uma recursão que retorna valor e índice
func min(vet []int) int {
	if(len(vet) == 0){
		return -1
	}
	_ = vet
	return 0
}

func main() {
	var vet []int
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Fields(line)
		fmt.Println("$" + line)

		switch args[0] {
		case "end":
			return
		case "read":
			vet = nil
			for _, arg := range args[1:] {
				if val, err := strconv.Atoi(arg); err == nil {
					vet = append(vet, val)
				}
			}
		case "tostr":
			fmt.Println(tostr(vet))
		case "torev":
			fmt.Println(tostrrev(vet))
		case "reverse":
			reverse(vet)
		case "sum":
			fmt.Println(sum(vet))
		case "mult":
			fmt.Println(mult(vet))
		case "min":
			fmt.Println(min(vet))
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
