package main
import (
    "fmt"
    "bufio"
    "os"
)

func main() {
    
    scanner := bufio.NewReader(os.Stdin)

    var entrada string
    fmt.Fscan(scanner, &entrada)

    palavra := []rune{}
    cursor := 0

    for _, caractere := range entrada {

        // [abc<<<]

        switch caractere {

        case 'R':
            palavra = append(palavra[:cursor], append([]rune{'\n'}, palavra[cursor:]...)...)
            cursor++
        case 'B':
            if(cursor > 0){
                palavra = append(palavra[:cursor-1], palavra[cursor:]...)
                cursor--
            }
        case 'D':

            if(cursor < len(palavra)){
                palavra = append(palavra[:cursor], palavra[cursor+1:]...)
                // cursor--
            }
    
    
        case '>':
            if(cursor < len(palavra)){
                cursor++
            }
        case '<':
            if(cursor > 0){
                cursor--
            }
            // fmt.Println(cursor)
        default:
            palavra = append(palavra[:cursor], append([]rune{caractere}, palavra[cursor:]...)...)
            cursor++
            // fmt.Printf("caractere:%c cursor:%d,\n",caractere,cursor)
        }

    }

    // fmt.Println("-----------")

    // fmt.Printf("%c\n",palavra)
    // fmt.Printf("cursor: %d\n",cursor)

    for i := 0; i <= len(palavra); i++{

        if i == cursor{
            fmt.Print("|")
        }

        if(i < len(palavra)){
            fmt.Printf("%c", palavra[i])
        }

    }

    fmt.Println()

}
