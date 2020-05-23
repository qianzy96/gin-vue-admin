package router

import (
	"gin-vue-admin/api/v1"
    "gin-vue-admin/middleware"
    "github.com/gin-gonic/gin"
)

func InitInfoRouter(Router *gin.RouterGroup) {
	InfoRouter := Router.Group("info").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		InfoRouter.POST("createInfo", v1.CreateInfo)     // 新建Info
		InfoRouter.DELETE("deleteInfo", v1.DeleteInfo)   //删除Info
		InfoRouter.PUT("updateInfo", v1.UpdateInfo)   //更新Info
		InfoRouter.GET("findInfo", v1.FindInfo)           // 根据ID获取Info
		InfoRouter.GET("getInfoList", v1.GetInfoList) //获取Info列表
	}
}
