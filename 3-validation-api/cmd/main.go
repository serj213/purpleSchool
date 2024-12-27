package main

import (
	"fmt"
	"net/http"
	"purpleschool/3-validation-api/configs"
	"purpleschool/3-validation-api/internal/verify"
)



func main(){
	config := configs.NewConfig()

	route := http.NewServeMux()

	verify.NewVerifyHandler(route, verify.VerifyHandlerDeps{
		Config: config,
	})

	serve := http.Server{
		Addr: ":8081",
		Handler: route,
	}


	fmt.Println("server starting")
	if err := serve.ListenAndServe(); err !=nil{
		fmt.Println("error starting server ", err)
	}
}

