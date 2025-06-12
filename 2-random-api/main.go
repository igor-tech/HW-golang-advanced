package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func random(w http.ResponseWriter, _ *http.Request) {
	randomNumber := rand.Intn(6) + 1
	w.Header().Set("Content-Type", "text/plain")
	_, err := w.Write([]byte(strconv.Itoa(randomNumber)))
	if err != nil {
		fmt.Printf("произошла ошибка при формировании запроса %v\n", err)
	}
}

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	router := http.NewServeMux()
	router.HandleFunc("/random", random)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Println("Server started on port 8080")
	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
