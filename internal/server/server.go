package server

import (
	"fmt"
	"log"
	"net/http"
)

func Start() {
	fmt.Println("Server started on :8080")
	routerConfig()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
