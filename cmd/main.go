package main

import (
	"log"

	"github.com/azybk/mini-forum/internal/configs"
	"github.com/azybk/mini-forum/internal/handler/memberships"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	var(
		cfg *configs.Config
	)

	err := configs.Init(
		configs.WithConfigFolder([]string{"./internal/configs"}),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)
	if err != nil {
		log.Fatal("Gagal inisiasi Config", err)
	}

	cfg = configs.Get()
	log.Println("config", cfg)

	handlerMemberships := memberships.NewHandler(r)
	handlerMemberships.RegisterRoute()
	r.Run(cfg.Service.Port) // listen and serve on 0.0.0.0:8080
}
