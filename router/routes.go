package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lucassouzz/gopportunities/handler"
)

func initializeRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		v1.GET("/opening", handler.GetOpeningHandler)
		v1.POST("/opening", handler.PostOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
		v1.PUT("/opening", handler.PutOpeningHandler)
		v1.GET("/openings", handler.ListOpeningsHandler)
	}
}
