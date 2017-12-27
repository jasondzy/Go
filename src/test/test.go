package main 
import (
	"fmt"
)

func main(){

	var array = []int{1,2,3}
	array = append(array,4)

	fmt.Printf("%T  %v  %d  %d",array,array,len(array),cap(array))
}