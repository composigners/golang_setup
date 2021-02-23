package main

import (
	"example/env"
	"example/routers"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("start")
	err := env.SetupEnv()
	if err != nil {
		return
	}

	r := routers.InitRouter()
	server := &http.Server{
		Addr: "",
		Handler: r,
	}
	server.SetKeepAlivesEnabled(false)
	fmt.Println(server.ListenAndServe())
}
