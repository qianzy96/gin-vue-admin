package v1

import (
	"fmt"
    "gin-vue-admin/global/response"
    "gin-vue-admin/model"
    "gin-vue-admin/model/request"
    resp "gin-vue-admin/model/response"
    "gin-vue-admin/service"
    "github.com/gin-gonic/gin"
)


// @Tags Info
// @Summary 创建Info
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Info true "创建Info"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /info/createInfo [post]
func CreateInfo(c *gin.Context) {
	var info model.Info
	_ = c.ShouldBindJSON(&info)
	err := service.CreateInfo(info)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}


// @Tags Info
// @Summary 删除Info
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Info true "删除Info"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /info/deleteInfo [delete]
func DeleteInfo(c *gin.Context) {
	var info model.Info
	_ = c.ShouldBindJSON(&info)
	err := service.DeleteInfo(info)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}


// @Tags Info
// @Summary 更新Info
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Info true "更新Info"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /info/updateInfo [put]
func UpdateInfo(c *gin.Context) {
	var info model.Info
	_ = c.ShouldBindJSON(&info)
	err := service.UpdateInfo(&info)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}


// @Tags Info
// @Summary 用id查询Info
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Info true "用id查询Info"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /info/findInfo [get]
func FindInfo(c *gin.Context) {
	var info model.Info
	_ = c.ShouldBindQuery(&info)
	err,reinfo := service.GetInfo(info.ID)
	if err != nil {
	response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData( gin.H{"reinfo":reinfo,}, c)
	}
}


// @Tags Info
// @Summary 分页获取Info列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Info列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /info/getInfoList [get]
func GetInfoList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetInfoInfoList(pageInfo)
	if err != nil {
	    response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
        response.OkWithData(resp.PageResult{
                    List:     list,
                    Total:    total,
                    Page:     pageInfo.Page,
                    PageSize: pageInfo.PageSize,
                }, c)
	}
}