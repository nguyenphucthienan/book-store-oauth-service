package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nguyenphucthienan/book-store-oauth-service/service"
	"net/http"
)

func NewHandler(service service.AccessTokenService) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

type AccessTokenHandler interface {
	GetById(c *gin.Context)
}

type accessTokenHandler struct {
	service service.AccessTokenService
}

func (h *accessTokenHandler) GetById(c *gin.Context) {
	accessToken, err := h.service.GetById(c.Param("access_token_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": accessToken})
}
