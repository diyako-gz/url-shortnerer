package main

import (
	"log"
	"urlShortnerer/cmd/config"
	"urlShortnerer/cmd/server"
)

func main() {
	config.LoadedEnv()
	if err := server.Start(); err != nil {
		log.Fatal("faild to start server", "error:", err)
	}

}
