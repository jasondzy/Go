build:
	protoc --proto_path=$(GOPATH)/src/Project_test/go-micro/shippy/proto/vessel:. --micro_out=$(GOPATH)/src/Project_test/go-micro/shippy/proto/vessel --go_out=$(GOPATH)/src/Project_test/go-micro/shippy/proto/vessel $(GOPATH)/src/Project_test/go-micro/shippy/proto/vessel/vessel.proto
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o vessel-service .