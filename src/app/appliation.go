package app

import (
	"bookstore/bookstore_OAuth-api/src/domain/access_token"
	"bookstore/bookstore_OAuth-api/src/http"
	"bookstore/bookstore_OAuth-api/src/repository/db"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {

	atHandler := http.NewHandler(access_token.NewService(db.NewRepository()))

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)

	router.Run(":8080")
}
