package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ShowUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"name": "hello"})
}
