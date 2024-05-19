package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/admin/apis"
	"go-admin/common/actions"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerUserRouter2)
}

// registerUserRouter
func registerUserRouter2(v1 *gin.RouterGroup) {
	api := apis.User2{}
	r := v1.Group("/user2")
	{
		r.GET("", actions.PermissionAction(), api.GetPage)
		r.GET("/:id", actions.PermissionAction(), api.Get)
		r.PUT("/:id", actions.PermissionAction(), api.Update)
		r.DELETE("", api.Delete)
		r.POST("", api.Insert) //注册
		r.GET("/login", api.Login)

	}
	r.POST("upload", api.Upload)
}
