package main
import "fmt"
func main() {

	//resolvendo primeiro com slice
	//depois irei resolver com map

	//captura

	var qtdPessoas, qtdImpacientes int

	fmt.Scan(&qtdPessoas)

	listaDePessoas := make([]int, qtdPessoas)


	for i := 0; i < len(listaDePessoas); i++ {
		fmt.Scan(&listaDePessoas[i])
	}

	fmt.Scan(&qtdImpacientes)

	listaDeImpacientes :=  make([]int, qtdImpacientes)

	for i := 0; i < len(listaDeImpacientes); i++ {
		fmt.Scan(&listaDeImpacientes[i])
	}

	//remocão
	// filaFinal := make([]int, (qtdPessoas - qtdImpacientes))
	for i := 0; i < len(listaDePessoas); i++ {
		
		for  j := 0; j < len(listaDeImpacientes); j++{
			if(listaDePessoas[i] == listaDeImpacientes[j]){
				// filaFinal = append(filaFinal, listaDePessoas[i])
				
				listaDePessoas = append(listaDePessoas[:i],listaDePessoas[i+1:]...)
				i--
				listaDeImpacientes = append(listaDeImpacientes[:j],listaDeImpacientes[j+1:]...)
				j--
				break
			}
		}
	}


	// filaFinal := make([]int, (qtdPessoas - qtdImpacientes))
	// for i, valor := range listaDePessoas {
		
	// 	for  j, _ := range listaDeImpacientes {
	// 		if(listaDePessoas[i] != listaDeImpacientes[j]){
	// 			// filaFinal[i] = valor
	// 			filaFinal = append(filaFinal, valor)
	// 			break
	// 		}
	// 	}
	// }

	//exibição

	for i := 0; i < len(listaDePessoas); i++ {
		
		if(i == len(listaDePessoas) - 1){
			fmt.Printf("%v \n", listaDePessoas[i])
		}else{
			fmt.Printf("%v ", listaDePessoas[i])
		}
	}

}
/*






*/