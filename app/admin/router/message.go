package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/admin/apis"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerMessageRouter)
}

// registerBlogsRouter
func registerMessageRouter(v1 *gin.RouterGroup) {
	api := apis.Message{}
	r := v1.Group("/message")
	{
		r.GET("", api.GetPage)
		r.GET("/:id", api.Get)
		r.PUT("/:id", api.Update)
		r.PUT("/all", api.AllRead)
		r.DELETE("", api.Delete)
		r.DELETE("delete_all", api.DeleteAll)
	}
}
