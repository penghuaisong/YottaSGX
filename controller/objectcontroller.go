package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//GetInfo test router
func GetInfo(g *gin.Context) {
	var msg string
	msg = "Test golang router....."
	g.JSON(http.StatusOK, gin.H{"msg": msg})
}
