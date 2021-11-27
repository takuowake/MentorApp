package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ShowHome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"name": "home"})
}
