PROJECTNAME=$(shell basename "$(PWD)")

proto-file:
	#"Build proto file"
	protoc -I. --go_out=plugins=micro:. proto/vessel/vessel.proto

build:
	#"Build locally our app"
	env CGO_ENABLED=0  GOOS=linux go build -a -installsuffix cgo -o ./builds/$(PROJECTNAME)

	#"Building container"
	#docker build -t $(PROJECTNAME) .

run:
	#"Running docker container"
	docker run -p 50052:50051 -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=mdns $(PROJECTNAME)