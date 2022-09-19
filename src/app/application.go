package app

import (
	"bookstore/src/github.com/luckyparakh/bookstore_oauth-api/src/domain/access_token"
	"bookstore/src/github.com/luckyparakh/bookstore_oauth-api/src/http"
	"bookstore/src/github.com/luckyparakh/bookstore_oauth-api/src/repository/db"

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
