package main

import (
	"log"
	"context"
	"github.com/smallnest/rpcx/client"
)

type Args struct { 
	A int   
	B int
}

type Reply struct { 
	C int  
}

func main() {
	d := client.NewPeer2PeerDiscovery("tcp@127.0.0.1:8972", "") //defines a service discovery implementation. In this example we use the simplest discovery: Peer2PeerDiscovery. Client get the server address by accessing this discovery and connect the server directly.
	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close() //creates a XClient, and passes FailMode, SelectMode and default option. FailMode indicates the client how to handle call failures: retry, return fast or retry another server? SelectMode indicates the client how to select a server if there are multiple servers for one service.

	args := &Args {  //这里定义要传入到server的参数，注意这里的参数要定义的和server中保持一致
		A: 10,
		B: 20,
	}

	reply := &Reply{} //这里定义一个空的reply数据，用来接收数据

	err := xclient.Call(context.Background(), "Mul", args, reply) //执行调用server中的函数功能，call the remote service and gets the result synchronously
	if err != nil {
		log.Fatalf("failed to call:%v", err)
	}

	log.Printf("%d * %d = %d", args.A, args.B, reply.C)

	//也可使用如下的异步方式来调用服务器函数，通过go的方式新创建一个gorounie来运行client函数
	// call, err := xclient.Go(context.Background(), "Mul", args, reply, nil)
    // if err != nil {
    //     log.Fatalf("failed to call: %v", err)
    // }

    // replyCall := <-call.Done
    // if replyCall.Error != nil {
    //     log.Fatalf("failed to call: %v", replyCall.Error)
    // } else {
    //     log.Printf("%d * %d = %d", args.A, args.B, reply.C)
    // }
}