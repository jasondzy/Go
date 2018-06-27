build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o consignment-cli .
	cp ./consignment-cli ./../docker/consignment-cli/
	cp ./consignment.json ./../docker/consignment-cli/
