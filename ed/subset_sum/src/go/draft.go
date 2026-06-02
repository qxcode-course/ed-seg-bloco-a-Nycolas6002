package main
import "fmt"

func backtracking(nums []int, index int, soughtValue int, currentSum int) bool{

    //soma chegou ao valor certo
    if(soughtValue == currentSum){
        return true
    }

    //soma maior que o procurado
    if(currentSum > soughtValue){
        return false
    }
    
    //verificou todos
    if(index == len(nums)){
        return false
    }
    

    //soma
    if(backtracking(nums, index+1, soughtValue, currentSum+nums[index])){
        return true
    }

    //não soma
    if(backtracking(nums, index+1, soughtValue, currentSum)){
        return true
    }

    return false

}

func main() {
    
                            //radução:valor procurado em inglês
    var quantityElements, soughtValue int

    fmt.Scan(&quantityElements,&soughtValue)

    nums := []int{}

    for i := 0; i < quantityElements; i++ {
        var varAux int
         fmt.Scan(&varAux)
        nums = append(nums, varAux )
    }

    // fmt.Println(nums, quantityElements, soughtValue)

    if( backtracking(nums, 0, soughtValue, 0)){
        fmt.Println("true")
    }else{
        fmt.Println("false")
    }



}