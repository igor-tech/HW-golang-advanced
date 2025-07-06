package main

import (
	"fmt"
	"order/api/configs"
	"order/api/internal/product"
	"order/api/pkg/db"
)

func main() {
	conf := configs.LoadConfig()
	fmt.Println(conf.DbConfig.Dsn)
	db := db.NewDb(conf)

	if err := db.AutoMigrate(&product.Product{}); err != nil {
		panic(err)
	}
}
