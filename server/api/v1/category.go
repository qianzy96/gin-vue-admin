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


// @Tags Category
// @Summary 创建Category
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Category true "创建Category"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /category/createCategory [post]
func CreateCategory(c *gin.Context) {
	var category model.Category
	_ = c.ShouldBindJSON(&category)
	err := service.CreateCategory(category)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}


// @Tags Category
// @Summary 删除Category
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Category true "删除Category"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /category/deleteCategory [delete]
func DeleteCategory(c *gin.Context) {
	var category model.Category
	_ = c.ShouldBindJSON(&category)
	err := service.DeleteCategory(category)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}


// @Tags Category
// @Summary 更新Category
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Category true "更新Category"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /category/updateCategory [put]
func UpdateCategory(c *gin.Context) {
	var category model.Category
	_ = c.ShouldBindJSON(&category)
	err := service.UpdateCategory(&category)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}


// @Tags Category
// @Summary 用id查询Category
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Category true "用id查询Category"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /category/findCategory [get]
func FindCategory(c *gin.Context) {
	var category model.Category
	_ = c.ShouldBindQuery(&category)
	err,recategory := service.GetCategory(category.ID)
	if err != nil {
	response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData( gin.H{"recategory":recategory,}, c)
	}
}


// @Tags Category
// @Summary 分页获取Category列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Category列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /category/getCategoryList [get]
func GetCategoryList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetCategoryInfoList(pageInfo)
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