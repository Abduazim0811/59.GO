package main

import (
	"context"
	"log"
	"time"

	pb "Weather/genproto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:1108", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewWeatherServiceClient(conn)

	req := &pb.WeatherRequest{Location: "Samarqand"}
	stream, err := client.GetWeatherUpdates(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	for {
		res, err := stream.Recv()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Weather update for %s: %s, %.2f°C, %.2f%% humidity, at %s",
			res.Location, res.Description, res.Temperatura, res.Humidity, time.Unix(res.Timestamp, 0))
	}
}
