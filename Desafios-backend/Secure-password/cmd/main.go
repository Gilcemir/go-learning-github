package main

import (
	"net/http"
	"secure-password/internal/web"
)

func main() {
	var handler web.HttpHandler = &web.HttpHandlerImpl{}
	router := http.NewServeMux()
	router.HandleFunc("POST /secure-password", handler.CheckPassword)
	http.ListenAndServe(":8080", router)
}
