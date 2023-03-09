package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func DefaultHandler(c *gin.Context) {
	res := "Working fine"
	c.JSON(http.StatusOK, gin.H{
		"body": res,
	})
}
