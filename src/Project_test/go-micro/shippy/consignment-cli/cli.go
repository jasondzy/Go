package main

import (
	pb "Project_test/go-micro/shippy/proto/consignment"
	"github.com/micro/go-micro/registry/consul"
	"github.com/hashicorp/consul/api"
	"io/ioutil"
	"encoding/json"
	"errors"
	// "google.golang.org/grpc"  //import for grpc
	"log"
	"os"
	"context"
	"github.com/micro/go-micro"
)

//这里定义常量的作用是尽量减少代码中的硬编码，这里若是对交互地址的更改会更加方便，避免了代码中多处引用而不能一致改动的问题
const (
	// ADDRESS				= "localhost:50051"
	DEFAULT_INFO_FILE	= "consignment.json"
)

//接下来的函数的作用是读取json中的数据

func paraseFile(fileName string) (*pb.Consignment, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	var consignment *pb.Consignment
	err = json.Unmarshal(data, &consignment) //调用json函数的解析函数，将上文读取到的二进制值解析到对应的结构体中
	if err != nil {
		return nil, errors.New("consignment.jso file content error")
	}

	return consignment, nil
}

//结下来正是实现客户端的程序

func main() {

/****************** grpc *****************************************/	
/*
	//连接调grpc服务器，从而创建一个连接到grpc服务器的对象
	conn, err := grpc.Dial(ADDRESS, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("connect error: %v", err)
	}

	defer conn.Close()
	//对上边的连接conn调用pb中的函数进行封装，这样client就具备了相应的调用函数，这里的作用就相当于一个闭包，将上边的连接对象进行了封装，使其具有了一定的方法
	client := pb.NewShippingServiceClient(conn)
*/
/*******************************************************************/

/********************** go-micro ***********************************/

	options := &api.Config{Address:"192.168.31.119:8500"}  //这里通过api中的Config生成一个options类型的结构
	registry := consul.NewRegistry(consul.Config(options))  //这里的registry是Registry接口的一个实例，由于是接口类型所以只具有方法，参数是传入的options这样就可以配置client的默认连接

	service := micro.NewService(
		micro.Registry(registry),
	)
	service.Init()  //这里和server端是一样的，同时通过micro的NewDervice函数来初始化一个service，只不过接下来对这个service的封装不一样，就决定了其是client还是server

	client := pb.NewShippingService("go.micro.srv.consignment", service.Client()) //这里调用pb中的函数将service封装成了一个client

	//指定货物信息
	infofile := DEFAULT_INFO_FILE

	if len(os.Args) >1 {
		infofile = os.Args[1]
	}
	//解析出json中的数据
	consignment, errFile := paraseFile(infofile)
	if errFile != nil {
		log.Fatalf("parase file info error:%v", errFile)
	}

	log.Printf(" the send data is %v",consignment)
	//将解析出的数据传入到调用函数中，从而获取一个返回值
	resp, errResp := client.CreateConsignment(context.Background(), consignment)
	// log.Printf(" the receive data is %v", resp)
	if errResp != nil {
		log.Fatalf("create consignemnt error: %v", errResp)
	}

	log.Printf("created: %t \n", resp.Created)
	log.Printf("infomation: %v \n\n", resp)

	log.Printf("Get all consignment:\n")

	respAll, errall := client.GetConsignment(context.Background(), &pb.Getmessage{}) //第二个参数传入的要是一个地址值才可以
	if errall != nil {
		log.Fatalf("failed to get all consignments")
	}

	for _, data := range respAll.Consignment {
		log.Printf("data:%v \n", data)
	}


}

