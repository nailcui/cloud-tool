package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime/debug"
	"strconv"
)

var (
	b         [][]byte
	totalSize = 0
)

func MemoryAlloc(c *gin.Context) {
	sizeString := c.Query("size")
	size, err := strconv.Atoi(sizeString)
	if err != nil {
		c.String(http.StatusOK, "param of size error: %s\n", sizeString)
	} else {
		a := make([]byte, 1024 * 1024 * size)
		for i := 0; i < len(a); i += 4096 {
			a[i] = 'x'
		}
		b = append(b, a)
		totalSize += size
		c.String(http.StatusOK, "success, total size: %dM\n", totalSize)
	}
}

func MemoryRelease(c *gin.Context) {
	b = [][]byte{}
	totalSize = 0
	debug.FreeOSMemory()
}
