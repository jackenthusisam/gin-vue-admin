package game

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// DouyinMiniGame 抖音小游戏 结构体
type DouyinMiniGame struct {
	global.GVA_MODEL
	Name        string `json:"name" form:"name" gorm:"column:name;comment:游戏名称;type:varchar(100);"`
	AppID       string `json:"appId" form:"appId" gorm:"column:app_id;comment:AppID;type:varchar(100);"`
	AppSecret   string `json:"appSecret" form:"appSecret" gorm:"column:app_secret;comment:AppSecret;type:varchar(100);"`
	Status      *int   `json:"status" form:"status" gorm:"column:status;comment:状态 1开启 2关闭;type:int;"`
	Description string `json:"description" form:"description" gorm:"column:description;comment:描述;type:text;"`
}

// TableName 抖音小游戏 DouyinMiniGame自定义表名 douyin_mini_games
func (DouyinMiniGame) TableName() string {
	return "douyin_mini_games"
}
