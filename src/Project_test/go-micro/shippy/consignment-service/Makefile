build:
	#buidl proto for grpc
	# protoc -I. --go_out=plugins=grpc:$(GOPATH)/src/Project_test/go-micro/shippy/consignment-service proto/consignment/consignment.proto

	#build proto for go-micro
	protoc --proto_path=$(GOPATH)/src/Project_test/go-micro/shippy/proto/consignment:. --micro_out=$(GOPATH)/src/Project_test/go-micro/shippy/proto/consignment --go_out=$(GOPATH)/src/Project_test/go-micro/shippy/proto/consignment $(GOPATH)/src/Project_test/go-micro/shippy/proto/consignment/consignment.proto

	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o consignment-service .
	cp ./consignment-service ./../docker/consignment-server/
