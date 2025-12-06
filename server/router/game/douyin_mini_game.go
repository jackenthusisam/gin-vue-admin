package game

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DouyinMiniGameRouter struct{}

// InitDouyinMiniGameRouter 初始化 抖音小游戏 路由信息
func (s *DouyinMiniGameRouter) InitDouyinMiniGameRouter(Router *gin.RouterGroup) {
	douyinMiniGameRouter := Router.Group("douyinMiniGame").Use(middleware.OperationRecord())
	douyinMiniGameRouterWithoutRecord := Router.Group("douyinMiniGame")
	var douyinMiniGameApi = v1.ApiGroupApp.GameApiGroup.DouyinMiniGameApi
	{
		douyinMiniGameRouter.POST("createDouyinMiniGame", douyinMiniGameApi.CreateDouyinMiniGame)             // 新建抖音小游戏
		douyinMiniGameRouter.DELETE("deleteDouyinMiniGame", douyinMiniGameApi.DeleteDouyinMiniGame)           // 删除抖音小游戏
		douyinMiniGameRouter.DELETE("deleteDouyinMiniGameByIds", douyinMiniGameApi.DeleteDouyinMiniGameByIds) // 批量删除抖音小游戏
		douyinMiniGameRouter.PUT("updateDouyinMiniGame", douyinMiniGameApi.UpdateDouyinMiniGame)              // 更新抖音小游戏
	}
	{
		douyinMiniGameRouterWithoutRecord.GET("findDouyinMiniGame", douyinMiniGameApi.FindDouyinMiniGame)       // 根据ID获取抖音小游戏
		douyinMiniGameRouterWithoutRecord.GET("getDouyinMiniGameList", douyinMiniGameApi.GetDouyinMiniGameList) // 获取抖音小游戏列表
	}
}
