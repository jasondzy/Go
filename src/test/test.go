package main 
import (
	"fmt"
)

func add(args ...int){
	sum := 0
	for _, item := range args{
		sum += item
	}

	fmt.Println("sum:",sum)
}

// func show(args ...interface{}){
// 	for _, item := range args{
// 		switch item.(type) {
// 		case int:
// 			fmt.Println("int value:", item)
// 		case string:
// 			fmt.Println("string value:", item)
// 		default:
// 			fmt.Println("unknow type")
// 		}
// 	}

// 	// add(args...)
// }

func main(){
	show := func (args ...interface{}){
	for _, item := range args{
		switch item.(type) {
		case int:
			fmt.Println("int value:", item)
		case string:
			fmt.Println("string value:", item)
		default:
			fmt.Println("unknow type")
		}
	}

	// add(args...)
}
	show(1,2,"hello",4,5)
}