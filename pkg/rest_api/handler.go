package rest_api

import (
	"github.com/ammiranda/connectRN/pkg/user_service"
	"github.com/gin-gonic/gin"
)

func NewRouter(u user_service.Service) *gin.Engine {
	r := gin.New()

	r.Use(gin.Recovery())

	r.POST("/api/users", postUsersHandler(u))

	return r
}

func postUsersHandler(u user_service.Service) func(c *gin.Context) {
	return func(c *gin.Context) {
		var users request.UserRequestBody
		if err := c.ShouldBindJSON(&users); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	}
}
