package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var (
	status = 200
)

func Health(c *gin.Context) {
	c.String(status, strconv.Itoa(status))
}

func Status(c *gin.Context) {
	paramStatus := c.Query("status")
	newStatus, err := strconv.Atoi(paramStatus)
	if err != nil {
		c.String(http.StatusOK, err.Error())
	} else {
		c.String(http.StatusOK, "status: %d", newStatus)
	}
}

func SetStatus(c *gin.Context) {
	paramStatus := c.Query("status")
	newStatus, err := strconv.Atoi(paramStatus)
	if err != nil {
		c.String(http.StatusOK, "param status unknown: %s\n", paramStatus)
	} else {
		status = newStatus
		c.String(http.StatusOK, "health status set to: %d\n", newStatus)
	}
}
