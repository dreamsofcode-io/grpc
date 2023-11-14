package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/dreamsofcode-io/grpc/grpcalc/pb"
)

func main() {
	serverAddr := flag.String(
		"server", "localhost:8080",
		"The server address in the format of host:port",
	)
	flag.Parse()

	creds := credentials.NewTLS(&tls.Config{InsecureSkipVerify: false})

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, *serverAddr, opts...)
	if err != nil {
		log.Fatalln("fail to dial:", err)
	}
	defer conn.Close()

	client := pb.NewCalculatorClient(conn)

	res, err := client.Sum(ctx, &pb.NumbersRequest{
		Numbers: []int64{10, 10, 10, 10, 10},
	})
	if err != nil {
		log.Fatalln("error sending request:", err)
	}

	fmt.Println("result:", res.Result)
}
