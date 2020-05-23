import service from '@/utils/request'

// @Tags Info
// @Summary 创建Info
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Info true "创建Info"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /info/createInfo [post]
export const createInfo = (data) => {
     return service({
         url: "/info/createInfo",
         method: 'post',
         data
     })
 }


// @Tags Info
// @Summary 删除Info
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Info true "删除Info"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /info/deleteInfo [delete]
 export const deleteInfo = (data) => {
     return service({
         url: "/info/deleteInfo",
         method: 'delete',
         data
     })
 }

// @Tags Info
// @Summary 更新Info
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Info true "更新Info"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /info/updateInfo [put]
 export const updateInfo = (data) => {
     return service({
         url: "/info/updateInfo",
         method: 'put',
         data
     })
 }


// @Tags Info
// @Summary 用id查询Info
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Info true "用id查询Info"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /info/findInfo [get]
 export const findInfo = (params) => {
     return service({
         url: "/info/findInfo",
         method: 'get',
         params
     })
 }


// @Tags Info
// @Summary 分页获取Info列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Info列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /info/getInfoList [get]
 export const getInfoList = (params) => {
     return service({
         url: "/info/getInfoList",
         method: 'get',
         params
     })
 }