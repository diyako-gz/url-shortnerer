package main

import (
	"fmt"
	"os"
	"urlShortnerer/cmd/server"
)

func main() {
	err := server.Run()
	if err != nil {
		fmt.Println("failed to start server or proplem in mongo or redis", "error:", err)
		os.Exit(1)
	}

}
