package main

import "fmt"


func eh_primo(x int, div int) bool {

	if(x <= 1){
		return false
	}

	//condição de parada
	//outra opção é colocar div*div > x
	if(x == div){
		return true
	}

	//se dividir não é primo
			
	if(x % div == 0){
		return false
	}

	//chamada recursiva
	return eh_primo(x, div+1)
	
}


func main() {
	var x int
	fmt.Scan(&x)
	if eh_primo(x, 2) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
