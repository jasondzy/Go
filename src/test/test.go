package main 
import (
	"fmt"
)

type  interger int

func (i interger) show() {
	fmt.Println("i:",i)
}

func (i *interger) add(j interger) {
	*i += j
	fmt.Println("i:",*i)
}
  
type inter interface { 
	show()
	add(j interger) 
}

func main(){
	var a interger

	var ii inter = &a 

	ii.add(3)

	ii.show()
}