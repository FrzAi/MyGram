package app

import (
	"Final_Project/config"
	"Final_Project/repository"
	"Final_Project/route"
	"Final_Project/service"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

var router = gin.New()

func StartApplication() {
	repo := repository.NewRepo(config.PSQL.DB)
	app := service.NewService(repo)
	route.RegisterApi(router, app)

	port := os.Getenv("PORT")
	router.Run(fmt.Sprintf(":%s", port))
}
