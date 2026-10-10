package main

import (
	"fmt"
	"os"
	"urlShortnerer/cmd/server"
	"urlShortnerer/cmd/shortcode"
)

func main() {
	code1 := shortcode.CreateShortCode()
	code2 := shortcode.CreateShortCode()
	code3 := shortcode.CreateShortCode()
	println(code1, code2, code3)
	err := server.Run()
	if err != nil {
		fmt.Println("failed to start server or problem in mongo or redis", "error:", err)
		os.Exit(1)
	}

}
