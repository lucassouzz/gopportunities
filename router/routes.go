package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lucassouzz/gopportunities/handler"
)

func initializeRoutes(r *gin.Engine) {
	handler.Init()

	v1 := r.Group("/api/v1")
	{
		v1.GET("/opening", handler.ShowOpeningHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		v1.GET("/openings", handler.ListOpeningsHandler)
	}
}
