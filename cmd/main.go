package main

// This is the entry point (main)

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/basicServer/internal/database"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("GET /users", foo)

	server := http.Server{
		Addr:      fmt.Sprintf(":%s", "8080"),
		Handler:   router,
		TLSConfig: nil,
	}

	fmt.Println("Server is running")
	database.InitDB()
	//database.InsertCars("Nissan", "Versa", 2016)
	database.GetCars()
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}

func foo(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Request body:", string(body))

	data := struct {
		Message string `json"message"`
	}{
		Message: "ping",
	}
	jsonResponse, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

type bar struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
