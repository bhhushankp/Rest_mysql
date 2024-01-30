package main

import (
	"log"
	"net/http"
	"rest_mysql/routers"
)

func main() {
	router := routers.InitializeRoute()

	log.Fatal(http.ListenAndServe(":8181", router))
}
