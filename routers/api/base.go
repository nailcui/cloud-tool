package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func Root(c *gin.Context) {
	c.String(http.StatusOK, "ok\n")
}

func GetInfo(c *gin.Context) {
	c.String(http.StatusOK, "ok\n")
}

func GetIp(c *gin.Context) {
	c.String(http.StatusOK, "ok\n")
}

func PostInfo(c *gin.Context) {
	c.String(http.StatusOK, "ok\n")
}

func Exit(c *gin.Context) {
	os.Exit(1)
}
