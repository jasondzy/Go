package main

import (
	pb "Project_test/go-micro/shippy/proto/vessel"
	"github.com/micro/go-micro"
	"context"
	"errors"
	"log"
)

type Respority interface {
	FindAvilable(*pb.Specification) (*pb.Vessel, error) 
}

type VesselRespority struct {
	vessels []*pb.Vessel
}

func (repo *VesselRespority) FindAvilable(spec *pb.Specification) (*pb.Vessel, error) {
	
	for _, v := range repo.vessels {
		if v.Capicity >= spec.Capicity && v.MaxWeight >= spec.MaxWeight {
			return v, nil
		}
	}

	return nil, errors.New("no vessel can bu used")
}

type service struct {  //这里的结构体可以任意进行定义，但是必须要实现FindAvailable方法
	repo VesselRespority
}

func (s *service)  (ctx context.Context, spec *pb.Specification, resp *pb.Response) error {
	v, err := s.repo.FindAvilable(spec)
	if err != nil {
		return err
	}

	resp.Vessel = v
	return nil
}

func main() {

	vessels := []*pb.Vessel{
		{Id: "vessel001", Name: "Boaty McBoatface", MaxWeight: 200000, Capicity: 500},
	}

	repo := VesselRespority{vessels}

	server := micro.NewService(
		micro.Name("go.micro.srv.vessel"),
		micro.Version("latest"),
	)

	server.Init()

	pb.RegisterVesselServiceHandler(server.Server(), &service{repo})

	if err := server.Run(); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}
}