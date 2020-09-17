package application

import (
	"github.com/gin-gonic/gin"
	"github.com/nguyenphucthienan/book-store-oauth-service/handler"
	"github.com/nguyenphucthienan/book-store-oauth-service/repository"
	"github.com/nguyenphucthienan/book-store-oauth-service/service"
)

const (
	apiPrefix = "/api"
)

func Start() {
	accessTokenHandler := handler.NewAccessTokenHandler(
		service.NewAccessTokenService(
			repository.NewAccessTokenRepository(),
			service.NewRestUserService(),
		),
	)

	router := gin.Default()
	router.GET(apiPrefix+"/oauth/access_tokens/:access_token_id", accessTokenHandler.GetById)
	router.POST(apiPrefix+"/oauth/access_tokens", accessTokenHandler.Create)

	if err := router.Run("localhost:8080"); err != nil {
		panic(err)
	}
}
