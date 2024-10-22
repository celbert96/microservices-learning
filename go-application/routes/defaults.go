package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddDefaultRoutes(rg *gin.RouterGroup) {
	defaultGroup := rg.Group("/")

	defaultGroup.GET("/", getGreeting)
}

func getGreeting(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, "Hello!")
}
