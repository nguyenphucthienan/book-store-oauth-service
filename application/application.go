package application

import (
	"github.com/gin-gonic/gin"
	"github.com/nguyenphucthienan/book-store-oauth-service/handler"
	"github.com/nguyenphucthienan/book-store-oauth-service/repository"
	"github.com/nguyenphucthienan/book-store-oauth-service/service"
)

var (
	router = gin.Default()
)

func Start() {
	accessTokenHandler := handler.NewHandler(service.NewService(repository.NewRepository()))

	router.GET("/oauth/access_token/:access_token_id", accessTokenHandler.GetById)

	if err := router.Run("localhost:8080"); err != nil {
		panic(err)
	}
}
