package main

import(
	"net/rpc"
	"net/http"
	"errors"
	"net"
	"log"
)


type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

//这里定义了一个类型，接下来便是针对这个类型创建的方法
type Arith int

//如下是针对该类型创建的方法，注意该函数需要满足rpc要求的格式，即首字母大写，并且参数都是大写
func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}

	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func main() {
	//调用new方法来创建一个变量，这样便会获得所创建变量的地址，这里arith保存的是变量Arith的一个地址
	arith := new(Arith)
	//调用register来注册该变量类型
	rpc.Register(arith)
	//通过HTTP的方式来处理远程连接
	rpc.HandleHTTP()
	//开启端口监听，这个和HTTP的注册一致
	l, e := net.Listen("tcp", "127.0.0.1:1234")

	if e != nil {
		log.Fatal("listen error;", e)
	}
	//启动HTTP服务器，这里和启动HTTP一致
	go http.Serve(l, nil)

	//这里是一个阻塞，保证server一致处在运行中
	select {

	}

}