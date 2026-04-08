package main
import "fmt"

func fibonnaci(n,k int) int{

    if(n == 1){
        return 1
    }

    if(n == 2){
        return 1
    }

    resultado := fibonnaci(n-1,k) + k  * fibonnaci(n-2,k) 

    return resultado
}

func main() {
    var n, k int
    fmt.Scan(&n,&k)
    fmt.Println(fibonnaci(n,k))
}
