package main

import (
	"github.com/ammiranda/connect_rn/pkg/rest_api"
)

func main() {

	router := rest_api.NewRouter()

	router.Run()
}
