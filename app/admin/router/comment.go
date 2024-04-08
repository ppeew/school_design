package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/admin/apis"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerCommentsRouter)
}

// registerBlogsRouter
func registerCommentsRouter(v1 *gin.RouterGroup) {
	api := apis.Comments{}
	r := v1.Group("/comments")
	{
		r.GET("", api.GetPage)
		r.GET("/:id", api.Get)
		r.POST("", api.Insert)
		r.PUT("/:id", api.Update)
		r.DELETE("", api.Delete)
	}
}
