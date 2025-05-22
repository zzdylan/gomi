// Package user 存放用户 Model 相关逻辑
package user

import (
	"gomi/app/models"
	"gomi/pkg/database"
	"gomi/pkg/hash"
)

// User 用户模型
type User struct {
	models.BaseModel
	Username string `json:"username" gorm:"type:varchar(255);not null;unique"`
	Name     string `json:"name,omitempty" gorm:"type:varchar(255)"`
	Email    string `json:"-" gorm:"type:varchar(255)"`
	Phone    string `json:"-" gorm:"type:varchar(255)"`
	Password string `json:"-" gorm:"type:varchar(255)"`
	Avatar   string `json:"avatar,omitempty" gorm:"type:varchar(255)"`

	models.CommonTimestampsField
}

// Create 创建用户，通过 User.ID 来判断是否创建成功
func (userModel *User) Create() {
	database.DB.Create(&userModel)
}

// ComparePassword 密码是否正确
func (userModel *User) ComparePassword(_password string) bool {
	return hash.BcryptCheck(_password, userModel.Password)
}

func (userModel *User) Save() (rowsAffected int64) {
	result := database.DB.Save(&userModel)
	return result.RowsAffected
}
