//Vessel-service/Makefile

echo "Build proto file"
protoc -I. --go_out=plugins=micro:. proto/vessel/vessel.proto

echo "Build locally our app"
env CGO_ENABLED=0  GOOS=linux go build -a -installsuffix cgo -o ./builds/shippy-service-vessel

echo "Building container"
docker build -t shippy-service-vessel .


echo "Running docker container"
docker run -p 50052:50051 -e MICRO_SERVER_ADDRESS=50051 shippy-service-vessel