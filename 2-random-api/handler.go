package main

import (
	"math/rand"
	"net/http"
	"strconv"
	"time"
)


type RandomNumHandler struct {
}

func NewRandomHandler(router *http.ServeMux){
	handler := &RandomNumHandler{}
	router.HandleFunc("/random", handler.GetRandomNumber())
}

func (r *RandomNumHandler) GetRandomNumber() http.HandlerFunc{
	return func(w http.ResponseWriter, req *http.Request){
		randomNum := getGenerateNum()
		w.Write([]byte(strconv.Itoa(randomNum)))
	}
}


func getGenerateNum()int{
	rand.NewSource(time.Now().UnixNano())
	return rand.Intn(6) 
}