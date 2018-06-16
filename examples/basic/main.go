package main

import (
	"net/http"

	"github.com/go-mego/cors"
	"github.com/go-mego/mego"
)

func main() {
	e := mego.Default()
	e.Use(cors.Default())
	e.GET("/", func(c *mego.Context) {
		c.String(http.StatusOK, "CORS is supported with the server")
	})
	e.Run()
}
