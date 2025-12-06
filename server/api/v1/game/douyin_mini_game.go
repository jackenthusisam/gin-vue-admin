package game

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/game"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DouyinMiniGameApi struct{}

var douyinMiniGameService = service.ServiceGroupApp.GameServiceGroup.DouyinMiniGameService

// CreateDouyinMiniGame 创建抖音小游戏
func (a *DouyinMiniGameApi) CreateDouyinMiniGame(c *gin.Context) {
	var d game.DouyinMiniGame
	err := c.ShouldBindJSON(&d)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := douyinMiniGameService.CreateDouyinMiniGame(d); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteDouyinMiniGame 删除抖音小游戏
func (a *DouyinMiniGameApi) DeleteDouyinMiniGame(c *gin.Context) {
	var d game.DouyinMiniGame
	err := c.ShouldBindJSON(&d)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := douyinMiniGameService.DeleteDouyinMiniGame(d); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteDouyinMiniGameByIds 批量删除抖音小游戏
func (a *DouyinMiniGameApi) DeleteDouyinMiniGameByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := douyinMiniGameService.DeleteDouyinMiniGameByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateDouyinMiniGame 更新抖音小游戏
func (a *DouyinMiniGameApi) UpdateDouyinMiniGame(c *gin.Context) {
	var d game.DouyinMiniGame
	err := c.ShouldBindJSON(&d)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := douyinMiniGameService.UpdateDouyinMiniGame(d); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindDouyinMiniGame 用id查询抖音小游戏
func (a *DouyinMiniGameApi) FindDouyinMiniGame(c *gin.Context) {
	var d game.DouyinMiniGame
	err := c.ShouldBindQuery(&d)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reDouyinMiniGame, err := douyinMiniGameService.GetDouyinMiniGame(d.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reDouyinMiniGame": reDouyinMiniGame}, c)
	}
}

// GetDouyinMiniGameList 分页获取抖音小游戏列表
func (a *DouyinMiniGameApi) GetDouyinMiniGameList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := douyinMiniGameService.GetDouyinMiniGameInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
