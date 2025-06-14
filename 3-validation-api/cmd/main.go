package main

import (
	"net/http"
	"verify/email/configs"
	"verify/email/internal/verify"
)

func main() {
	conf := configs.NewConfig()
	router := http.NewServeMux()
	verify.NewVerifyHandler(router, verify.VerifyHandlerDeps{conf})

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
