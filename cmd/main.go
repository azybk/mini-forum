package main

import (
	"log"

	"github.com/azybk/mini-forum/internal/configs"
	"github.com/azybk/mini-forum/internal/handler/memberships"
	"github.com/azybk/mini-forum/internal/handler/posts"
	membershipRepo "github.com/azybk/mini-forum/internal/repository/memberships"
	postRepo "github.com/azybk/mini-forum/internal/repository/posts"
	membershipSvc "github.com/azybk/mini-forum/internal/service/memberships"
	postSvc "github.com/azybk/mini-forum/internal/service/posts"
	"github.com/azybk/mini-forum/pkg/internalsql"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	var (
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

	db, err := internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatal("Gagal inisiasi database", err)
	}

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	membershipRepo := membershipRepo.NewRepository(db)
	postRepo := postRepo.NewRepository(db)

	membershipService := membershipSvc.NewService(cfg, membershipRepo)
	postService := postSvc.NewService(cfg, postRepo)

	handlerMemberships := memberships.NewHandler(r, membershipService)
	handlerMemberships.RegisterRoute()

	handlerPost := posts.NewHandler(r, postService)
	handlerPost.RegisterRoute()

	r.Run(cfg.Service.Port) // listen and serve on 0.0.0.0:8080
}
