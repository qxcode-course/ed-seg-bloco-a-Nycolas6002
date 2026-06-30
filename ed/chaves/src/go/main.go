package main

import (
	"fmt"
)

func main() {

	fila := NewQueue[string]()

	for letra := 'A'; letra <= 'P'; letra++ {
		fila.Enqueue(string(letra))
		// fmt.Println(string(letra))
	}

	for i := 0; i < 15; i++ {

		time1 := fila.Dequeue()
		time2 := fila.Dequeue()

		var golsTime1, golsTime2 int

		fmt.Scan(&golsTime1, &golsTime2)

		if(golsTime1 > golsTime2){
			fila.Enqueue(time1)
		}else{
			fila.Enqueue(time2)
		}
		
	}

	fmt.Println(fila.Dequeue())

}
