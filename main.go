package main

import (
	"fmt"

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
		err := server.Init()
	}()
	err := <-end
	fmt.Println("ERROR:", err)
}
