package main

import (
	"fmt"
	"go-dep-inj/libs"

	"github.com/aws/aws-lambda-go/lambda"
)

// Handler holds dependencies for the lambda handler.
type Handler struct {
	RedisClient libs.RedisClient
}

func (h *Handler) HandleRequest(request map[string]interface{}) {
	fmt.Println("Handle request > ", request)
	conn, err := h.RedisClient.GetConnection("host.docker.internal:6379")
	fmt.Println("Err > ", err)
	fmt.Println("Conn > ", conn)
	test, err := libs.CheckKeyInRedis(conn, "TEST")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("test > ", test)
}

func main() {
	handler := &Handler{
		RedisClient: &libs.RedisClientImpl{},
	}
	lambda.Start(handler.HandleRequest)
}
