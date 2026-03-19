package main
import "fmt"
func main() {
    var h,p,f,d int

    fmt.Scan(&h,&p,&f,&d)

    var posicao = f

    for{

        posicao = (posicao + d + 16) %16
        
        if(posicao == p){
            fmt.Println("N")
            break
        }
        if(posicao == h){
            fmt.Println("S")
            break
        }

    }
}
