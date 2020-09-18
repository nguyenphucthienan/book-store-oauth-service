package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nguyenphucthienan/book-store-oauth-service/domain/access_token"
	"github.com/nguyenphucthienan/book-store-oauth-service/service"
	"github.com/nguyenphucthienan/book-store-utils-go/errors"
	"net/http"
)

func NewAccessTokenHandler(accessTokenService service.AccessTokenService) AccessTokenHandler {
	return &accessTokenHandler{
		accessTokenService: accessTokenService,
	}
}

type AccessTokenHandler interface {
	GetById(*gin.Context)
	Create(*gin.Context)
}

type accessTokenHandler struct {
	accessTokenService service.AccessTokenService
}

func (h *accessTokenHandler) GetById(c *gin.Context) {
	accessToken, err := h.accessTokenService.GetById(c.Param("access_token_id"))
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, accessToken)
}

func (h *accessTokenHandler) Create(c *gin.Context) {
	var request access_token.AccessTokenRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := errors.NewBadRequestError("Invalid JSON body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	accessToken, err := h.accessTokenService.Create(request)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusCreated, accessToken)
}
