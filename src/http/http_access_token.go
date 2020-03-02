package http

import (
	"bookstore/bookstore_OAuth-api/src/domain/access_token"
	"bookstore/bookstore_OAuth-api/src/utils/errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AccessTokenHandler interface {
	GetById(*gin.Context)
	Create(*gin.Context)
}

type accessTokenHandler struct {
	service access_token.Service
}

func NewHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (handler *accessTokenHandler) GetById(c *gin.Context) {
	accessToken, err := handler.service.GetById(c.Param("access_token_id"))

	if err != nil {
		c.JSON(err.Status, err.Message)
		return
	}
	c.JSON(http.StatusOK, accessToken)
}

func (handler *accessTokenHandler) Create(c *gin.Context) {
	var at access_token.AccessToken
	fmt.Println("data", at.AccessToken, at.ClientId, at.UserId, at.Expires)
	if err := c.ShouldBindJSON(&at); err != nil {
		restErr := errors.NewBadRequestError("Invalid json body :(")
		fmt.Println("errormessage", err.Error())
		c.JSON(restErr.Status, restErr.Message)
		return
	}

	if err := handler.service.Create(at); err != nil {
		c.JSON(err.Status, err.Message)
		return
	}
	c.JSON(http.StatusCreated, at)
}
