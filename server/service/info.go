package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateInfo
// @description   create a Info
// @param     info               model.Info
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateInfo(info model.Info) (err error) {
	err = global.GVA_DB.Create(&info).Error
	return err
}

// @title    DeleteInfo
// @description   delete a Info
// @auth                     （2020/04/05  20:22）
// @param     info               model.Info
// @return                    error

func DeleteInfo(info model.Info) (err error) {
	err = global.GVA_DB.Delete(info).Error
	return err
}

// @title    UpdateInfo
// @description   update a Info
// @param     info          *model.Info
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateInfo(info *model.Info) (err error) {
	err = global.GVA_DB.Save(info).Error
	return err
}

// @title    GetInfo
// @description   get the info of a Info
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    Info        Info

func GetInfo(id uint) (err error, info model.Info) {
	err = global.GVA_DB.Where("id = ?", id).First(&info).Error
	return
}

// @title    GetInfoInfoList
// @description   get Info list by pagination, 分页获取用户列表
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetInfoInfoList(info request.PageInfo) (err error, list interface{}, total int) {
	limit := info.PageSize
    	offset := info.PageSize * (info.Page - 1)
    	db := global.GVA_DB
    	var infos []model.Info
    	err = db.Find(&infos).Count(&total).Error
    	err = db.Limit(limit).Offset(offset).Find(&infos).Error
    	return err,infos, total
}
