package main 

import (
	"fmt"
)

type Item struct{
	value int
	num   int
}

func (item Item) cost() {
	fmt.Printf("the cost of item is %d \n",item.value*item.num)
}

type student struct{

	Item //这里嵌套了一个Item的类型数据，那么student中就会集成Item中的数据和方法
	name string
}

func main(){
	var st = student{Item{12,3}, "xiaoming"} //这里是对student的一个定义，初始化的时要注意Item类型的数据是如何初始化的

	fmt.Printf("the name is %s  value is %d \n",st.name, st.value)
	st.cost()

}