package main

import (
	"github.com/ammiranda/connectRN/pkg/image_service"
	"github.com/ammiranda/connectRN/pkg/rest_api"
	"github.com/ammiranda/connectRN/pkg/user_service"
)

func main() {
	u := user_service.NewService()
	i := image_service.NewService()

	router := rest_api.NewRouter(u, i)

	err := router.Run()
	if err != nil {
		panic(1)
	}
}
