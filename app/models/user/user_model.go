// Package user 存放用户 Model 相关逻辑
package user

import (
	"gomi/app/models"
)

// User 用户模型
type User struct {
	models.BaseModel

	Name     string `json:"name,omitempty" gorm:"type:varchar(255)"`
	Email    string `json:"-" gorm:"type:varchar(255)"`
	Phone    string `json:"-" gorm:"type:varchar(255)"`
	Password string `json:"-" gorm:"type:varchar(255)"`

	models.CommonTimestampsField
}
