package main

import (
	pb "Project_test/go-micro/shippy/proto/consignment"
	vesselpb "Project_test/go-micro/shippy/proto/vessel"
	"context"
	// "net"   //import for grpc 
	"log"
	// "google.golang.org/grpc" //import for grpc
	"errors"
	"github.com/micro/go-micro"
	// "github.com/micro/go-micro/broker"  
	// "encoding/json"
	// "time"
)

const (
	PORT = ":50051"
	topic = "user.created"
)
//这里的作用是定义仓库接口，这里之所以定一个仓库接口是为了统一下边定义的仓库对象
type IRepository interface {
	Create(consignment *pb.Consignment) (*pb.Consignment, error) //定义仓库接口的函数，这样满足该函数的对象也就满足了该接口
}

//接下来定义的是一个仓库对象
type Repository struct {
	consignments []*pb.Consignment //仓库对象中定义了所要存放的元素
}

func(repo *Repository) Create(consignment *pb.Consignment) ([]*pb.Consignment, error) { //定义了该对象的方法，这样便就实现了上边定义的接口
	repo.consignments = append(repo.consignments, consignment)
	return repo.consignments, nil
}

func(repo *Repository) GetAll() ([]*pb.Consignment, error) {
	return repo.consignments, nil
}


//接下来回到该函数的重点，即定义一个server对象，由于server的接口和方法在pb中已经定义了，所以这里需要定义一个具体的对象，并实现该对象的接口方法即可
//serve对象可以任意定义，即使是一个空的结构体都行，主要是要实现该对象的接口方法
//由于为了要实现创建仓库的功能，所以这里的结构体中包含了一个仓库结构体

type service struct {
	repo Repository
	client vesselpb.VesselService  //set for vessel service client
	// pubsub broker.Broker  //这里定义的broker用来存放的是server端获得的broker对象，接下来server端便是通过该对象进行过publish操作
	pubsub micro.Publisher //这里使用的是pubsub层来实现发布订阅功能
}

//这里定义了该service结构体的方法，该方法也是client和server之间可以调用的方法
//grpc调用方式
// func(s *service) CreateConsignment(ctx context.Context, req *pb.Consignment) (resp *pb.Response, error) {
//go-micro调用方式
func(s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, resp *pb.Response) error {
	log.Printf("call the createconsignment service\n")
	
	/**************************调用vessel中的服务********************************/
	vreq := &vesselpb.Specification{
		Capicity:int32(len(req.Containers)),
		MaxWeight: req.Weight,
	}

	vResp, err := s.client.FindAvailable(context.Background(), vreq)
	if err != nil {
		return err
	}

	log.Printf(" FInd a vessel %v \n", vResp.Vessel)
	/****************************end***********************************/
	consignment, err := s.repo.Create(req)
	if err != nil {
		return err
	}

	// resp = &pb.Response{Created: true, Consignment: consignment} //通过这种方式来向client返回值是不可以的，这样client是无法接受到返回值的
	//通过如下的赋值方式我们也可以清楚的知道，client将返回值的地址进行了固定，若采用如上的方式则返回值的地址发生了变化，这样client在之前的地址中是无法取到值的
	resp.Created = true
	resp.Consignment = consignment

	/***************************这里调用publish函数将信息发送到订阅了该消息的客户端中去**********/
/*
	if err := s.publishEvent(req); err != nil {  //这里使用的是broker作为ie发布和订阅功能
		log.Fatalf(" publishEvent function error %v", err)
	}
*/
	if err := s.pubsub.Publish(ctx, req); err != nil {
		log.Fatalf(" publishEvent function error %v", err)
	}
	/*************************** end *****************************************************/

	// log.Printf(" the return data is %v", resp)

	return nil
}

//grpc使用方式
// func(s *service) GetConsignment(ctx context.Context, message *pb.Getmessage） （resp *pb.Response, error) {

//go-micro使用方式
func(s *service) GetConsignment(ctx context.Context, message *pb.Getmessage, resp *pb.Response)  error {
	
	allconsignments, err := s.repo.GetAll()
	if err != nil {
		return errors.New("No consignment ")
	}
	// resp = &pb.Response{Created:true, Consignment: allconsignments}//这里的返回值方式也是错误的，同上

	resp.Created = true
	resp.Consignment = allconsignments


	return nil
}

/*
//如下定义的是broker的publish功能
func(s *service) publishEvent(message *pb.Consignment) error {
	body, err := json.Marshal(message)  //由于broker使用的是http的机制，所以这里使用的是json进行数据的通信
	if err != nil {
		return err
	}
	msg := &broker.Message{  //这是对要publsih的数据进行相关格式的封装，该具体的格式在broker中进行了定义，所以所传递的数据要按照这样的要求来
		Header: map[string]string{
			"id": message.Id,
		},
		Body: body,
	}
	if err := s.pubsub.Publish(topic, msg); err != nil {  //这里使用的是broker的发布订阅动能
		
		log.Fatalf("pub fatal %v", err)
	}
	return nil
}

*/

func main() {

/**************** grpc 使用方式******************************************/
/*
	listener, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("failed to listen :%v", err)
	}

	log.Printf("listen on: %s", PORT)

	//这里创建一个grpc的服务器
	server := grpc.NewServer()

	//这里创建一个仓库对象
	repo := Repository{}

	//这里使用pb中的注册函数将上边创建的server和service 对象进行绑定,这样也就绑定了一个服务方法
	pb.RegisterShippingServiceServer(server, &service{repo})

	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)

	}
*/

/*****************go-micro 调用方式*****************************************/
	//这里对server的参数进行配置，Newervice调用的是service.go文件下的newervice函数，该函数返回一个service结构体对象，该结构体对象实现了Service这个接口中的所有方法,也就实现了该Serveice接口
	server := micro.NewService(
		micro.Name("go.micro.srv.consignment"),//里边的这些参数会在调用newService函数的时候被初始化
		micro.Version("latest"),
		// micro.RegisterTTL(time.Second*30),
		// micro.RegisterInterval(time.Second*10),
	)
	//将上边配置的参数进行init的操作,进行了该操作后一个服务器便产生了，其中client和server均是通过该server产生的
	server.Init()

	/*****************这里增加一个 broker实现的publish-subscribe的功能**************************/
/*
	pubsub := server.Server().Options().Broker   //这里获取的是server的默认 broker，broekr是一个代理用来使用服务之间的异步通信
	//上边是通过server对象来获取的一个broker代理
	if err := pubsub.Connect(); err != nil {  //这里也要连接到broker，这样才能正常的通过broker来publish消息
		log.Fatalf("publish server connect failed %v", err)
	}
*/
	/******************end*******************************************************/

	/******************这里用go-micro的pubsub层来实现发布和订阅功能*******************/

	publisher := micro.NewPublisher(topic, server.Client())
	
	/************************end ************************************************/


	repo := Repository{}

	/***********************这里增加对vessel服务的调用************************************/
	vClient := vesselpb.NewVesselService("go.micro.srv.vessel", server.Client());

	/*************************end**************************************/

	//将上边init的server和service对象进行注册到service discovery中
	// pb.RegisterShippingServiceHandler(server.Server(),&service{repo, vClient, pubsub}) //这里使用的是broker实现的发布和订阅功能
	pb.RegisterShippingServiceHandler(server.Server(),&service{repo, vClient, publisher}) //这里使用的是pubsub层来实现

	//开启服务器
	if err := server.Run(); err != nil {
		log.Fatalf("failed to serve:%v", err)
	}


}
