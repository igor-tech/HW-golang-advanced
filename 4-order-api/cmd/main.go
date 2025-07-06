package main

import (
	"fmt"
	"net/http"
	"order/api/configs"
	"order/api/internal/product"
	"order/api/pkg/db"
)

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
		Handler: router,
	}

	fmt.Printf("Server starting on %s\n", conf.Address)
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
