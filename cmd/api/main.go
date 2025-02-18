package main

import (
	"fmt"

	"github.com/8soat-grupo35/fastfood-payment/internal/api/server"
	"github.com/8soat-grupo35/fastfood-payment/internal/external"
)

func main() {
	fmt.Println("Starting Payment Microservice")
	cfg := external.GetConfig()
	server.Start(cfg)
}
