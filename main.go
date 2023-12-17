package main

import (
	"exfillguard/internal/route"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()

	// ROTAS
	router.Use(gin.Logger())
	router.POST("/", route.UploadFile())
	router.GET(":id/:filename", route.DownloadFile())
	router.StaticFS("/portal", http.Dir("internal/static"))

	router.Run(":" + port)
}
