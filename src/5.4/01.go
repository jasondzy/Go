package main 

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){

	channel := make([]chan bool, 6) // 这里创建了6个通道的切片类型，即切片中有6个bool类型的通道，这里只是申明了6个chan类型，但是并没有创建

	for i := range channel{  //这里才是真正的给这个切片创建chan类型的值，类似是初始化的效果，不进行该操作的话，会报出奇怪的错误
		channel[i] = make(chan bool) 
	}

	defer func(){
		for i := range channel{
		close(channel[i])
		}
		}()

	var x int

	go func(){
		for {
			channel[rand.Intn(6)] <- true
			time.Sleep(3000*time.Millisecond)
		}

		}()

	for  {
	select { //这里是select语句，用法和switch比较相似，这里的select是阻塞的，加上一个default选项后就变成非阻塞类型
		case <-channel[0]:
			x=0
		case <-channel[1]:
			x=1
		case <-channel[2]:
			x=2
		case <-channel[3]:
			x=3
		case <-channel[4]:
			x=4
		case <-channel[5]:
			x=5
	}
	fmt.Printf("%d ",x)
}

	
 fmt.Println()
}