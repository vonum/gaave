package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Handler1(ctx *gin.Context) {
  ctx.JSON(http.StatusOK, "Hello")
}

func Handler2(ctx *gin.Context) {
  ctx.JSON(http.StatusOK, "World")
}
