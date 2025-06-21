package main

import (
	"log"

	pb "protobufgrpc/api"

	"strconv"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conSh, err := grpc.NewClient(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conSh.Close()

	conBw, err := grpc.NewClient(":50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conBw.Close()
	
	router := gin.Default()
	router.GET("/hello/:name", func(c *gin.Context) {
		grpcHelloClient := pb.NewHelloServiceClient(conSh)
		ctx := c.Request.Context()
		name := c.Param("name")
		res, err := grpcHelloClient.SayHello(ctx, &pb.HelloRequest{Name: name})
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": res.Message})
	})
	router.GET("/badword/:name/:jmlkt", func(c *gin.Context)  {
		grpcBadwordClient := pb.NewBadwordServiceClient(conBw)
		ctx := c.Request.Context()
		name := c.Param("name")
		jmlktStr := c.Param("jmlkt")
		jmlkt, err := strconv.Atoi(jmlktStr)
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid jmlkt parameter"})
			return
		}
		res, err := grpcBadwordClient.SayBadword(ctx, &pb.BadwordReq{Nama: name, JmlKata: int32(jmlkt)})
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": res.Kalimat})
	})

	router.Run(":3000")
}

