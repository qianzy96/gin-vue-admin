// 自动生成模板Info
package model

import (
    "github.com/jinzhu/gorm"
)

type Info struct {
      gorm.Model 
      Desc  string `json:"desc" gorm:"column:desc"`
}