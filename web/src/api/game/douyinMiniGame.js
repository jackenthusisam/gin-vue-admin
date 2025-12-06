import service from '@/utils/request'

// @Tags DouyinMiniGame
// @Summary 创建抖音小游戏
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DouyinMiniGame true "创建抖音小游戏"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /douyinMiniGame/createDouyinMiniGame [post]
export const createDouyinMiniGame = (data) => {
    return service({
        url: '/douyinMiniGame/createDouyinMiniGame',
        method: 'post',
        data
    })
}

// @Tags DouyinMiniGame
// @Summary 删除抖音小游戏
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DouyinMiniGame true "删除抖音小游戏"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /douyinMiniGame/deleteDouyinMiniGame [delete]
export const deleteDouyinMiniGame = (data) => {
    return service({
        url: '/douyinMiniGame/deleteDouyinMiniGame',
        method: 'delete',
        data
    })
}

// @Tags DouyinMiniGame
// @Summary 批量删除抖音小游戏
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除抖音小游戏"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /douyinMiniGame/deleteDouyinMiniGameByIds [delete]
export const deleteDouyinMiniGameByIds = (data) => {
    return service({
        url: '/douyinMiniGame/deleteDouyinMiniGameByIds',
        method: 'delete',
        data
    })
}

// @Tags DouyinMiniGame
// @Summary 更新抖音小游戏
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DouyinMiniGame true "更新抖音小游戏"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /douyinMiniGame/updateDouyinMiniGame [put]
export const updateDouyinMiniGame = (data) => {
    return service({
        url: '/douyinMiniGame/updateDouyinMiniGame',
        method: 'put',
        data
    })
}

// @Tags DouyinMiniGame
// @Summary 用id查询抖音小游戏
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DouyinMiniGame true "用id查询抖音小游戏"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /douyinMiniGame/findDouyinMiniGame [get]
export const findDouyinMiniGame = (params) => {
    return service({
        url: '/douyinMiniGame/findDouyinMiniGame',
        method: 'get',
        params
    })
}

// @Tags DouyinMiniGame
// @Summary 分页获取抖音小游戏列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取抖音小游戏列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /douyinMiniGame/getDouyinMiniGameList [get]
export const getDouyinMiniGameList = (params) => {
    return service({
        url: '/douyinMiniGame/getDouyinMiniGameList',
        method: 'get',
        params
    })
}
