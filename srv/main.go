package main

import (
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	if err := runServer(); err != nil {
		log.Fatal(err)
	}
}

func runServer() error {
	e := echo.New()
	e.Static("/", "../web/dist")

	return e.Start(":8086")
}
