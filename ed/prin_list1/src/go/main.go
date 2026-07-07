package main

import (
	"fmt"
	"strings"
)

// mostra a lista com o elemento sword destacado
func ToStr(l *DList[int], sword *DNode[int]) string {

	var out []string

	for it := l.Front(); it != l.End(); it = it.Next() {
		if(it == sword){
			out = append(out, fmt.Sprintf(">%d",it.Value))
		}else{
			out = append(out, fmt.Sprintf("%d", it.Value))
		}

	}

	return "[" + strings.Join(out, " ") +"]"
}

// move para frente na lista circular
func Next(l *DList[int], it *DNode[int]) *DNode[int] {

	it = it.Next()

	if(it == l.End()){
		return l.Front()
	}

	return it

}

func main() {
	var qtd, chosen int
	fmt.Scan(&qtd, &chosen)
	fmt.Println(qtd, chosen)
	l := NewDList[int]()
	for i := 1; i <= qtd; i++ {
		l.PushBack(i)
	}
	sword := l.Front()
	for range chosen - 1 {
		sword = Next(l, sword)
	}
	for range qtd - 1 {
		fmt.Println(ToStr(l, sword))
		l.Erase(Next(l, sword))
		sword = Next(l, sword)
	}
	fmt.Println(ToStr(l, sword))
}
