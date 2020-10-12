package service

import (
	"go-gin-system/entity"
	"go-gin-system/global"
	_ "gorm.io/gorm"
)

func GetUserInfoList() (err error, list interface{}, total int64) {
	db := global.GVA_DB.Model(&entity.SysUser{})
	var userList []entity.SysUser
	err = db.Count(&total).Error
	err = db.Limit(0).Offset(10).Preload("Authority").Find(&userList).Error
	return err, userList, total
}
