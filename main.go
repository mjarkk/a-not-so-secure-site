package main

import (
	"fmt"
	"os"

	"github.com/mjarkk/a-not-so-secure-site/database"
	"github.com/mjarkk/a-not-so-secure-site/server"
)

func main() {
	end := make(chan error)
	go func() {
		err := database.Init()
		if err != nil {
			end <- err
		}
	}()

	go func() {
		end <- server.Init()
	}()
	err := <-end
	if err == nil {
		fmt.Println("server stopped working")
		os.Exit(1)
	}
	fmt.Println("ERROR:", err)
}
