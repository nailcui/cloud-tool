package api

import (
	"fmt"
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

func GetHostname(c *gin.Context) {
	hostname, _ := os.Hostname()
	result := fmt.Sprintf("%s\n", hostname)
	c.String(http.StatusOK, result)
}

func PostInfo(c *gin.Context) {
	c.String(http.StatusOK, "ok\n")
}

func Exit(c *gin.Context) {
	os.Exit(1)
}
