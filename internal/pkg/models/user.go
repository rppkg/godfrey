package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID       string  `gorm:"column:id;type:string;primaryKey;comment:主键ID" json:"id"`
	Username string  `gorm:"column:username;type:string;not null;unique:idx_username;comment:用户名" json:"username"`
	Nickname *string `gorm:"column:nickname;type:string;index:idx_nickname;comment:昵称" json:"nickname"`
	Password string  `gorm:"column:password;type:string;comment:密码" json:"password"`
	Salt     string  `gorm:"column:salt;type:string;comment:混淆" json:"salt"`
	Avatar   string  `gorm:"column:avatar;type:string;comment:头像" json:"avatar"`
	Email    string  `gorm:"column:email;type:string;index:idx_email;comment:邮箱" json:"email"`
	Phone    string  `gorm:"column:phone;type:string;index:idx_phone;comment:电话" json:"phone"`
	RoleID   string  `gorm:"column:role_id;type:string;index:idx_role_id;not null;comment:角色" json:"role_id"`

	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime:true;index:idx_created_at;comment:创建时间" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime:true;comment:更新时间" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index:idx_deleted_at;comment:删除时间" json:"deleted_at"`

	Role *Role `gorm:"foreignKey:RoleID;references:ID"`
}

func (*User) TableName() string {
	return "users"
}
