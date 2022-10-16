package main

import (
	"aoisoft/gateway/auth"
	"aoisoft/gateway/config"
	"aoisoft/gateway/order"
	"aoisoft/gateway/product"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	authSvc := auth.RegisterRoutes(r, &c)
	product.RegisterRoutes(r, &c, authSvc)
	order.RegisterRoutes(r, &c, authSvc)

	r.Run(c.Port)

}
