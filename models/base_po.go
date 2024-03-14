package models

import (
	"gorm.io/gorm/schema"
	"time"
)

type ControlBy struct {
	CreatedBy string `json:"createdBy" gorm:"column:created_by;index;comment:创建者"`
	UpdatedBy string `json:"updatedBy" gorm:"column:updated_by;index;comment:更新者"`
}

// SetCreatedBy 设置创建人id
func (e *ControlBy) SetCreatedBy(createdBy string) {
	e.CreatedBy = createdBy
}

// SetUpdatedBy 设置修改人id
func (e *ControlBy) SetUpdatedBy(updatedBy string) {
	e.UpdatedBy = updatedBy
}

type PrimaryKey struct {
	Id string `json:"id" gorm:"column:id;primaryKey;comment:主键编码" url:"id"`
}

// SetPrimaryKey 主键
func (e *PrimaryKey) SetPrimaryKey(primaryKey string) {
	e.Id = primaryKey
}

type ModelTime struct {
	CreatedAt time.Time `json:"createdTime" gorm:"column:created_time;comment:创建时间" url:"createdTime"`
	UpdatedAt time.Time `json:"updatedTime" gorm:"column:updated_time;comment:最后更新时间" url:"updatedTime"`
}

type ActiveRecord interface {
	schema.Tabler
	SetCreatedBy(createBy int64)
	SetUpdatedBy(updateBy int64)
	Generate() ActiveRecord
	GetId() interface{}
}
