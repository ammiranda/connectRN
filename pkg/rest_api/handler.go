package rest_api

import (
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Recovery())

	r.POST("/api/users")

	return r
}

func PostUsersHandler 
