package rest_api

import (
	"io/ioutil"
	"github.com/ammiranda/connectRN/pkg/rest_api/models/request"
	"github.com/ammiranda/connectRN/pkg/user_service"
	"github.com/ammiranda/connectRN/pkg/image_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter(u user_service.UserService, i image_service.ImageService) *gin.Engine {
	r := gin.New()

	r.Use(gin.Recovery())

	r.POST("/api/users", postUsersHandler(u))
	r.POST("/api/images", postImageHandler(i))

	return r
}

func postUsersHandler(u user_service.UserService) func(c *gin.Context) {
	return func(c *gin.Context) {
		var users request.UserRequestBody
		if err := c.BindJSON(&users); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		resp, err := u.ParseUsers(users)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, resp)
	}
}

func postImageHandler(i image_service.ImageService) func(c *gin.Context) {
	return func(c *gin.Context) {
		f, image, err := c.Request.FormFile("image")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		defer f.Close()

		fileBytes, err := ioutil.ReadAll(f)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		fileType := http.DetectContentType(fileBytes)

		if fileType != "image/jpeg" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "not in jpg format"})
			return
		}

		p, err := i.GenerateImage(fileBytes, image.Filename)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.Data(200, "image/png", p)
	}
}
