package router

import (
	"github.com/gin-gonic/gin"
	"github.com/leetpy/daisy/router/api"
)

func Load(r *gin.Engine) *gin.Engine {
	r.Use(gin.Recovery())

	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/add_event", api.AddEventHandler)
	}
	return r
}
