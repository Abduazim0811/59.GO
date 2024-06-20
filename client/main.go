package main

import (
    "context"
    "log"
    "time"

    "google.golang.org/grpc"
    pb "Weather/genproto"
)

func main() {
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()

    client := pb.NewWeatherServiceClient(conn)

    req := &pb.WeatherRequest{Location: "New York"}
    stream, err := client.GetWeatherUpdates(context.Background(), req)
    if err != nil {
        log.Fatalf("could not get weather updates: %v", err)
    }

    for {
        res, err := stream.Recv()
        if err != nil {
            log.Fatalf("error receiving weather updates: %v", err)
        }
        log.Printf("Weather update for %s: %s, %.2f°C, %.2f%% humidity, at %s",
            res.Location, res.Description, res.Temperatura, res.Humidity, time.Unix(res.Timestamp, 0))
    }
}
