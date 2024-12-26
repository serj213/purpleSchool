package main

import (
	"fmt"
	"net/http"
)



func main(){
	route := http.NewServeMux()

	NewRandomHandler(route)

	server := http.Server{
		Addr: ":8080",
		Handler: route,
	}

	fmt.Println("server starting")
	if err := server.ListenAndServe(); err != nil{
		fmt.Println("Ошибка при старте сервера ", err)
	}
}