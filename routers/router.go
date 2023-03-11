package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/nailcui/cloud-tool/routers/api"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	{
		r.GET("/", api.Root)
		r.GET("/hostname", api.GetHostname)
		r.GET("/get", api.GetInfo)
		r.POST("/post", api.PostInfo)
		r.GET("/exit", api.Exit)
	}

	{
		r.GET("/health", api.Health)
		r.GET("/health/status", api.Status)
		r.GET("/health/setStatus", api.SetStatus)
	}

	{
		r.GET("/memory/alloc", api.MemoryAlloc)
		r.GET("/memory/release", api.MemoryRelease)
	}

	{
		r.GET("/cpu/alloc", api.CpuAlloc)
		r.GET("/cpu/release", api.CpuRelease)
	}

	return r
}
