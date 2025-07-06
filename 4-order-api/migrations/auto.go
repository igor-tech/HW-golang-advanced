package main

import (
	"order/api/configs"
	"order/api/internal/product"
	"order/api/pkg/db"
)

func main() {
	conf := configs.LoadConfig()
	db := db.NewDb(conf)

	if err := db.AutoMigrate(&product.Product{}); err != nil {
		panic(err)
	}
}
