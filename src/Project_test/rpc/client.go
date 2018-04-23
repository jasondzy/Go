package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}


func main() {
	//这里通过HTTP的方式来连接远程rpc服务器
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("dialing", err)
	}

	//这里传入的参数的类型要确保和服务器中的变量类型保持一致
	args := &Args{7, 8}

	var reply int

	//这里通过client.call的方式来调用远程服务器，reply中保存的是返回值，该函数是不会创建一个异步进程来进行处理的
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error;", err)
	}

	fmt.Printf("Arith: %d*%d = %d", args.A, args.B, reply)

	//如下是通过异步的方式来调用server
	// 通过new的方式来创建变量，这样系统便会为该变量分配一个地址空间，同时变量获得的是一个地址。直接var quotient Quotient的方式变量是不会被分配地址空间的
	quotient := new(Quotient)
	//这里测client.Go是一个异步进程，会通过go的方式来派生一个gorounie
	//divacall是一个 *call类型，divcall.Done中保存了GO处理后返回的一个call类型，该返回的call类型中包含了返回值
	divcall := client.Go("Arith.Divide", args, quotient, nil)
	replyCall := <-divcall.Done

	//replycall这个变量接受的是一个通道中的值，该值是一个*rpc.call的指针类型，call在rpc中是一个结构体里边包含了返回值reply，该reply是一个空接口类型
	//这里(*Quotient)的作用是将replycall.Reply这个空接口类型转换为Quotine类型，这是空接口的一种类型断言方式，转换为Quotient类型后就可以调用QUO和rem这个变量了
	//这是使用*Quotient是因为replycall是一个指针类型，所以再进行类型断言的时候也要用Quotient的指针类型
	fmt.Println("divide value :", replyCall.Reply.(*Quotient).Quo, replyCall.Reply.(*Quotient).Rem)

}