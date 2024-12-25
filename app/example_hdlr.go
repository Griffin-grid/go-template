package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a App) ExampleHdlr(c *gin.Context) {

	c.JSON(http.StatusOK, "Welcome new user ;)")
}
