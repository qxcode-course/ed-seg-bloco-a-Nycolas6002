package main
import "fmt"

func combinacao(n,k int) int{

    // casos base
    if(n == k){
        return 1
    }

    if(k == 0){
        return 1
    }

    if(k == 1){
        return n
    }

    resultado := combinacao(n-1,k-1) + combinacao(n-1,k)

    return resultado

}

func main() {
    var n,k int
    fmt.Scan(&n, &k)

    fmt.Println(combinacao(n,k))
    
    // fmt.Println("Hello, World!")
}
