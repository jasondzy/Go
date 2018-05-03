package main

import (
	"log"
    "context"
    "github.com/smallnest/rpcx/server"
)

type Args struct { //这里定义的是传入的参数的结构
    A int  //注意这些参数都是需要在client端进行传入的，所以这里的变量需要大写
    B int
}

type Reply struct { //这里定义的是返回值参数
    C int //这里同上需要在client端进行使用，所以这里的参数需要大写
}

type Arith int //这里定义一个int类型

func (t *Arith) Mul(ctx context.Context, args *Args, reply *Reply) error { //定义该类型的一个方法，该方法会在client端被调用
   
    reply.C = args.A * args.B
    log.Printf(" %d * %d = %d", args.A, args.B, reply.C)
    return nil
}

func main() {
    s := server.NewServer()  //这里调用rpcx/server package中的函数来初始化一个server
    s.RegisterName("Arith", new(Arith), "") //将Arith类型注册到该server中
    s.Serve("tcp", ":8972") //开启服务器
}

