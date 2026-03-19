package main

import "fmt"

type coordenadas struct {
	x int
	y int
}

func main() {

	var qtdGomos int
	var dir string

	fmt.Scan(&qtdGomos, &dir)

	pontosNoPlano := make([]coordenadas, 0, qtdGomos)

	for i := 0; i < qtdGomos; i++ {
		var coordenadaX, coordenadaY int
		fmt.Scan(&coordenadaX, &coordenadaY)

		pontosNoPlano = append(pontosNoPlano, coordenadas{x: coordenadaX, y: coordenadaY})
	}

    for i := len(pontosNoPlano) - 1; i > 0; i-- {
			pontosNoPlano[i] = pontosNoPlano[i-1]
	}

	switch dir {
	case "U":

		pontosNoPlano[0].y--

	case "D":

		pontosNoPlano[0].y++

	case "L":

		pontosNoPlano[0].x--

	case "R":
		pontosNoPlano[0].x++
	}

	for _, valor := range pontosNoPlano {
		fmt.Println(valor.x, valor.y)
	}

}