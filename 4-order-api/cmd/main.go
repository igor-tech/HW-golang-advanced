package main

import (
	"fmt"
	"net/http"
	"order/api/configs"
	"order/api/internal/product"
	"order/api/pkg/db"
	"order/api/pkg/middleware"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.InfoLevel)
}

func main() {
	conf := configs.LoadConfig()
	database := db.NewDb(conf)
	productRepository := product.NewProductRepository(database.DB)

	if err := database.AutoMigrate(&product.Product{}); err != nil {
		panic(fmt.Sprintf("Failed to migrate database: %v", err))
	}

	router := http.NewServeMux()

	product.NewProductHandler(router, product.ProductHandlerDeps{
		ProductRepository: productRepository,
	})

	server := http.Server{
		Addr:    conf.Address,
		Handler: middleware.Logger(router),
	}

	fmt.Printf("Server starting on %s\n", conf.Address)
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
