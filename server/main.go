package main

import (
	pb "Weather/genproto"
	"log"
	"math/rand"
	"net"
	"time"

	"google.golang.org/grpc"
)

type WeatherServer struct{
	pb.UnimplementedWeatherServiceServer
}

func (s *WeatherServer) GetWeatherUpdates(req *pb.WeatherRequest, stream pb.WeatherService_GetWeatherUpdatesServer) error{
	location :=req.GetLocation()
	for {
		response := &pb.WeatherResponce{
			Location: location,
			Description: getDescription(),
			Temperatura: getTemperatura(),
			Humidity: getHumidity(),
			Timestamp: time.Now().Unix(),
		}

		if  err:=stream.Send(response); err!=nil{
			return err
		}

		time.Sleep(5 * time.Second)
	}
}


func getDescription() string {
	descriptions := []string{"Sunny", "Cloudy","Rainy","Windy"}
	return descriptions[rand.Intn(len(descriptions))]
}

func getTemperatura()float32{
	return 20 + rand.Float32()*(35-20)
}

func getHumidity()float32{
	return  40 + rand.Float32()*(80-40)
}

func main(){
	lis, err :=net.Listen("tcp", ":1108")
	if err!=nil{
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterWeatherServiceServer(s, &WeatherServer{})
	log.Printf("server listening at %v", lis.Addr())
	if  err:=s.Serve(lis); err!=nil{
		log.Fatal(err)
	}
}