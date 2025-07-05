package main

import (
	"net/http"
	"verify/email/configs"
	"verify/email/internal/verify"
)

func main() {
	conf := configs.LoadConfig()
	router := http.NewServeMux()
	verify.NewVerifyHandler(router, verify.VerifyHandlerDeps{Config: conf})

	server := http.Server{
		Addr:    conf.Address,
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
