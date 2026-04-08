package main
import "fmt"

func fibonacci(mes int) int{
    if( mes == 1){
        return 1
    }

    if(mes == 2){
        return 1
    }

    if(mes == 3){
        return 2
    }

    resultado := fibonacci(mes-2) + fibonacci(mes-3)

    return resultado

}

func main() {
   var mes int
   fmt.Scan(&mes);
    fmt.Println(fibonacci(mes))
}
