package main

import "fmt"

/*
#include <stdlib.h>
*/
import "C"

//注意以上的 import“C”和 #include<> 之间是不能存在任何空格的，必须紧紧的挨着，否则会报错。也可以使用//代替/*注释
func Random() int {
	return int(C.random()) //注意这里的random和下边的srandom都是小写
}

func Seed(i int) {
	C.srandom(C.uint(i)) //
}

func main() {
	Seed(100)
	fmt.Println("Random:", Random())
}
