package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/admin/apis"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerBlogsRouter)
}

// registerBlogsRouter
func registerBlogsRouter(v1 *gin.RouterGroup) {
	api := apis.Blogs{}
	r := v1.Group("/blogs")
	{
		r.GET("", api.GetPage)
		r.GET("/:id", api.Get)
		r.POST("", api.Insert)
		r.PUT("/:id", api.Update)
		r.DELETE("", api.Delete)
	}
}
