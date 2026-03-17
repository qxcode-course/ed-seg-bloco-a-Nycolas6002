package main
import "fmt"
func main() {
    var h,p,f,d int

    fmt.Scan(&h,&p,&f,&d)
    // fmt.Println(h,p,f,d)

    if(d == -1){
        // modulo := f % 16
        // posicao := f - modulo
        // if(posicao == h){
        //     fmt.Println("S",posicao)
        // }else{
        //     fmt.Println("N",posicao)
        // }
    }else{
                // 7
        posicao := (f + h) % 16
        fmt.Println(posicao)
    }
    
    // if(d == -1){

    //     for i := f; i < 15; i-- {

    //         if f == h {
    //             fmt.Println("S")
    //         }

    //     }

    // }
}
