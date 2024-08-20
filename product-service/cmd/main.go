package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"vision/product-service/internal/handler"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	h := handler.NewHandler()

	r.GET("/products", h.GetProducts)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
