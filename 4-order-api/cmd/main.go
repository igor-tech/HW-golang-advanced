package main

import (
	"fmt"
	"net/http"
	"order/api/configs"
	"order/api/pkg/db"
)

func main() {
	conf := configs.LoadConfig()
	_ = db.NewDb(conf)
	router := http.NewServeMux()

	server := http.Server{
		Addr:    conf.Address,
		Handler: router,
	}

	fmt.Printf("Server starting on %s\n", conf.Address)
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
