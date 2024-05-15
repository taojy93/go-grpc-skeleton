package main

import (
	"context"
	"log"
	"time"

	pb "go-grpc-skeleton/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewProductServiceClient(conn)

	// 调用 GetProduct 方法
	getProduct(client)

	// 调用 ListProducts 方法
	listProducts(client)

	// 调用 StreamProducts 方法
	streamProducts(client)
}

func getProduct(client pb.ProductServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.GetProductRequest{Id: 1}
	res, err := client.GetProduct(ctx, req)
	if err != nil {
		log.Fatalf("could not get product: %v", err)
	}

	log.Printf("Product: %v", res.GetProduct())
}

func listProducts(client pb.ProductServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.ListProductsRequest{}
	res, err := client.ListProducts(ctx, req)
	if err != nil {
		log.Fatalf("could not list products: %v", err)
	}

	for _, product := range res.GetProducts() {
		log.Printf("Product: %v", product)
	}
}

func streamProducts(client pb.ProductServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	req := &pb.StreamProductsRequest{}
	stream, err := client.StreamProducts(ctx, req)
	if err != nil {
		log.Fatalf("could not stream products: %v", err)
	}

	for {
		res, err := stream.Recv()
		if err != nil {
			if err == context.Canceled {
				log.Println("Stream canceled")
				break
			}
			log.Fatalf("error receiving streamed product: %v", err)
		}

		log.Printf("Product: %v", res.GetProduct())
	}
}
