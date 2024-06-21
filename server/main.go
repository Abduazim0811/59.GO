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
var (
	Descriptions string
)

func getDescription() string {
	descriptions := []string{"Sunny", "Cloudy","Rainy","Windy"}
	Descriptions = descriptions[rand.Intn(len(descriptions))]
	return Descriptions
}

func getTemperatura()float32{
	if Descriptions == "Sunny"{
		return 30 + rand.Float32()*(45-30)
	}else if Descriptions == "Cloudy"{
		return 20 + rand.Float32()*(30-20)
	}else if Descriptions == "Rainy"{
		return 15 + rand.Float32()*(20-15)
	}else{
		return 5 + rand.Float32()*(15-5)
	}
}

func getHumidity()float32{
	if Descriptions == "Sunny"{
		return 5 + rand.Float32()*(10-5)
	}else if Descriptions == "Cloudy"{
		return 20 + rand.Float32()*(30-20)
	}else if Descriptions == "Rainy"{
		return 50 + rand.Float32()*(70-50)
	}else{
		return 30 + rand.Float32()*(40-30)
	}
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