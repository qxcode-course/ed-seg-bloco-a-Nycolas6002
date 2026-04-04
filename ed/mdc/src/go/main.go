package main

import (
	"fmt"
)

func mdc(a, b int) int {

	if(b % a == 0){
		return a
	}

	return mdc(b,a%b)

}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(mdc(a, b))
}

/* fiz com for no Mycompiler depois implementei a ideia com recursão e melhorei a solução

    a := 276
    b := 192

    resultado := 0
    mdc := 0


    for{
        if(b % a == 0){
            mdc = a
            break
        }

        
        resultado = b % a // 42%30 = 12
        b = a // 42 passa a ser 30
        a = resultado // 30 / 12
        
}
    fmt.Println(mdc)



*/

/*
	if(a == 0){
		return b
	}

	if(b == 0){
		return a
	}

	maiorNumero := 0
	menorNumero := 0
	if(a > b){
		maiorNumero = a
		menorNumero = b
	}else{
		maiorNumero = b
		menorNumero = a
	}

	q := maiorNumero / menorNumero
	r := a % b
	valorA := b* r


	return mdc(a,b)

*/
