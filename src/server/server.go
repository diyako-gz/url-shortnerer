package server

import (
	"fmt"
	"net/http"
	"time"
)

type ServerConnection struct {
}

func ServerInit() {
	server := http.Server{
		Addr:         ":8080",
		ReadTimeout:  time.Second * 20,
		WriteTimeout: time.Second * 20,
		Handler:      &ServerConnection{},
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func (c *ServerConnection) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "server custom handler response")
}
