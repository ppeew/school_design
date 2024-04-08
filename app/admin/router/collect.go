package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/admin/apis"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerCollectsRouter)
}

// registerBlogsRouter
func registerCollectsRouter(v1 *gin.RouterGroup) {
	api := apis.Collects{}
	r := v1.Group("/collect")
	{
		r.GET("", api.GetPage)
		r.GET("/:id", api.Get)
		r.POST("", api.Insert)
		r.PUT("/:id", api.Update)
		r.DELETE("", api.Delete)
	}
	v1.POST("/removeCollect", api.RemoveCollect)
}
