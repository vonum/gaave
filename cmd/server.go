package main

import (
	"log"

  "github.com/vonum/gaave/server"

	"github.com/gin-gonic/gin"
)

func pls() {
  r := gin.Default()

  r.GET("/hello", server.Handler1)
  r.GET("/world", server.Handler2)

  err := r.Run(":8080")

  if err != nil {
    log.Fatal("Can't run on port 8080")
  }
}
