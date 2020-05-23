// 自动生成模板Category
package model

import (
    "github.com/jinzhu/gorm"
)

type Category struct {
      gorm.Model 
      Name  string `json:"name" gorm:"column:name"`
}