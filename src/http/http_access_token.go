package http

import (
	"bookstore/src/github.com/luckyparakh/bookstore_oauth-api/src/domain/access_token"
	"bookstore/src/github.com/luckyparakh/bookstore_oauth-api/src/domain/utils/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AccessTokenHandler interface {
	GetById(*gin.Context)
	Create(*gin.Context)
}
type accessTokenHandler struct {
	service access_token.Service
}

func (ath *accessTokenHandler) GetById(c *gin.Context) {
	accessToken, err := ath.service.GetById(c.Param("access_token_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, accessToken)
}
func (ath *accessTokenHandler) Create(c *gin.Context) {
	var at access_token.AccessToken
	if err := c.ShouldBindJSON(&at); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	if err := ath.service.Create(at); err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, at)
}
func NewHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}
