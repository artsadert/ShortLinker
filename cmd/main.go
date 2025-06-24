package main

import (
	"fmt"
	service "github.com/artsadert/ShortLinker/internal/application/serivce"
	"github.com/artsadert/ShortLinker/internal/infrastructure/db/redis"
	"github.com/artsadert/ShortLinker/internal/interface/api/rest"
	"github.com/gin-gonic/gin"
)

func main() {

	conn := redis.NewConnection()
	repo := redis.NewRedisLinkReepository(conn)

	service := service.NewLinkService(repo)

	r := gin.Default()
	rest.NewLinkController(r, service)

	fmt.Println("starting on http://localhost:8080")
	r.Run("localhost:8000")
}
