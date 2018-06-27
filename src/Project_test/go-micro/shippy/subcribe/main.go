package main

import (
	"github.com/micro/go-micro"
	// "github.com/micro/go-micro/broker"
	pb "Project_test/go-micro/shippy/proto/consignment"
	// "encoding/json"
	"context"
	"log"
)

const topic = "user.created"
type Subscriber struct {}  //这里任意定义了一个对象，主要是应用该对象的方法来实现订阅函数的回调功能

func main() {
	srv := micro.NewService(
		micro.Name("go.micro.srv.publish"),//这个服务是作为订阅消息的服务，所以这里的服务名称不要和发布消息的服务相同
		micro.Version("latest"),
	)

	srv.Init()

	//如下使用pubsub层来实现发布和订阅功能
	micro.RegisterSubscriber(topic, srv.Server(), new(Subscriber))


/*	//如下使用的是broekr来实现发布和订阅功能
	pubsub := srv.Server().Options().Broker
	//连接到broker中，这样才能正常的通过broker接收到publish发布的消息，同时publish端的服务也必须要连接到broker中
	if err := pubsub.Connect(); err != nil {
		log.Fatalf("broker connect error: %v\n", err)
	}
	//订阅消息，调用订阅函数来进行订阅，该函数中同时包含了回调函数
	_, err := pubsub.Subscribe(topic, func(pub broker.Publication) error {
		var req  *pb.Consignment
		if err := json.Unmarshal(pub.Message().Body, &req); err != nil {
			return err
		}
		log.Printf("[created user]:%v", req)
		
		return nil
	})

	if err != nil {
		log.Printf("subscribe error %v", err)
	}
*/

	//注意：由于这里的服务只是用来测试订阅功能的，所以这里的服务不需要要调用protobuf文件中的RegisterShippingServiceHandler来进行注册
	//同时从这里也可以验证protobuf中的RegisterShippingServiceHandler函数功能就是注册一个使用protobuf协议的功能函数，真正的go-micro的通信机制并不需要该注册函数
	//真正的go-micro中的通信直接运行如下的Run（）函数即可
	//运行该订阅了相关消息的服务
	if err := srv.Run(); err != nil {
		log.Fatalf(" run server failed %v", err)
	}

}

func (s *Subscriber) Process(ctx context.Context, data *pb.Consignment) error {
	log.Printf("Pick up a neww publish message: %v", data)
	return nil
}