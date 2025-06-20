package main

import (
	"log"

	pb "protobufgrpc/api"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	grpcClient := pb.NewHelloServiceClient(conn)

	r := gin.Default()
	r.GET("/:name", func(c *gin.Context) {
		ctx := c.Request.Context()
		name := c.Param("name")
		res, err := grpcClient.SayHello(ctx, &pb.HelloRequest{Name: name})
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": res.Message})
	})

	r.Run(":3000")
}

