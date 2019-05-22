package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/taller-go/src/api/controllers/myml"
	"github.com/mercadolibre/taller-go/src/api/controllers/ping"
)

const (
	port = ":8083"
)

var (
	router = gin.Default()
)

func main() {
	router.GET("/ping", ping.Ping)
	router.GET("/myml/:userID", myml.GetUser)
	router.Run(port)
}
