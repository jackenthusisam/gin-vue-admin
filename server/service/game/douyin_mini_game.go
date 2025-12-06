package game

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/game"
)

type DouyinMiniGameService struct{}

// CreateDouyinMiniGame 创建抖音小游戏记录
func (s *DouyinMiniGameService) CreateDouyinMiniGame(d game.DouyinMiniGame) (err error) {
	err = global.GVA_DB.Create(&d).Error
	return err
}

// DeleteDouyinMiniGame 删除抖音小游戏记录
func (s *DouyinMiniGameService) DeleteDouyinMiniGame(d game.DouyinMiniGame) (err error) {
	err = global.GVA_DB.Delete(&d).Error
	return err
}

// DeleteDouyinMiniGameByIds 批量删除抖音小游戏记录
func (s *DouyinMiniGameService) DeleteDouyinMiniGameByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]game.DouyinMiniGame{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateDouyinMiniGame 更新抖音小游戏记录
func (s *DouyinMiniGameService) UpdateDouyinMiniGame(d game.DouyinMiniGame) (err error) {
	err = global.GVA_DB.Save(&d).Error
	return err
}

// GetDouyinMiniGame 根据id获取抖音小游戏记录
func (s *DouyinMiniGameService) GetDouyinMiniGame(id uint) (d game.DouyinMiniGame, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&d).Error
	return
}

// GetDouyinMiniGameInfoList 分页获取抖音小游戏记录
func (s *DouyinMiniGameService) GetDouyinMiniGameInfoList(info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&game.DouyinMiniGame{})
	var douyinMiniGames []game.DouyinMiniGame
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&douyinMiniGames).Error
	return douyinMiniGames, total, err
}
