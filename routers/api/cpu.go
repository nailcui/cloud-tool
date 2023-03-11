package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

var (
	limiter = rate.NewLimiter(1, 1)
	result = math.Pow(rand.Float64(), rand.Float64())
)

func task() {
	math.Pow(rand.Float64(), rand.Float64())
}

func init() {
	for i := 0; i < 10; i++ {
		go func() {
			for {
				_ = limiter.Wait(context.TODO())
				task()
			}
		}()
	}
}

func CpuAlloc(c *gin.Context) {
	limitString := c.Query("limit")
	limit, err := strconv.Atoi(limitString)
	if err != nil || limit <= 0 {
		c.String(http.StatusOK, "param of limit error: %s\n", limitString)
	} else {
		limiter = rate.NewLimiter(rate.Limit(limit), limit)
		c.String(http.StatusOK, "success, total limit: %dM\n", limitString)
	}
}

func CpuRelease(c *gin.Context) {
	limiter = rate.NewLimiter(1, 1)
	fmt.Printf("result: %s\n", result)
}
