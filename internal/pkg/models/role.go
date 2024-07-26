package models

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	ID       string `gorm:"column:id;type:string;primaryKey;comment:主键ID" json:"id"`
	Name     string `gorm:"column:name;type:string;not null;comment:角色名" json:"name"`
	Identify string `gorm:"column:identify;type:string;not null;unique:idx_identify;default:normal;comment:身份标识" json:"identify"`

	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime:true;index:idx_created_at;comment:创建时间" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime:true;comment:更新时间" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index:idx_deleted_at;comment:删除时间" json:"deleted_at"`

	Users []User `gorm:"foreignKey:RoleID"`
}

func (*Role) TableName() string {
	return "roles"
}
