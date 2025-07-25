package main

import (
	"fmt"
	"net/http"
	"order/api/configs"
	"order/api/internal/auth"
	"order/api/internal/jwt"
	"order/api/internal/product"
	"order/api/internal/user"
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
	jwtSecret := jwt.NewSecret(conf.JwtSecret)

	// Migrations
	if err := database.AutoMigrate(&product.Product{}, &user.User{}); err != nil {
		panic(fmt.Sprintf("Failed to migrate database: %v", err))
	}

	// Repositories
	productRepository := product.NewProductRepository(database.DB)
	userRepository := user.NewUserRepository(database.DB)

	// Handlers
	router := http.NewServeMux()
	product.NewProductHandler(router, product.ProductHandlerDeps{
		ProductRepository: productRepository,
		Config:            conf,
	})
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		UserRepository: userRepository,
		JWT:            jwtSecret,
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
