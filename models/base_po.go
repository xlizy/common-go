package models

import (
	"gorm.io/gorm/schema"
	"time"
)

type ControlBy struct {
	CreatedBy int64 `json:"createdBy" gorm:"column:created_by;index;comment:创建者"`
	UpdatedBy int64 `json:"updatedBy" gorm:"column:updated_by;index;comment:更新者"`
}

// SetCreatedBy 设置创建人id
func (e *ControlBy) SetCreatedBy(createdBy int64) {
	e.CreatedBy = createdBy
}

// SetUpdatedBy 设置修改人id
func (e *ControlBy) SetUpdatedBy(updatedBy int64) {
	e.UpdatedBy = updatedBy
}

type PrimaryKey struct {
	Id int64 `json:"id" gorm:"primaryKey;comment:主键编码"`
}

// SetPrimaryKey 主键
func (e *PrimaryKey) SetPrimaryKey(primaryKey int64) {
	e.Id = primaryKey
}

type ModelTime struct {
	CreatedAt time.Time `json:"created_time" gorm:"column:created_time;comment:创建时间"`
	UpdatedAt time.Time `json:"updated_time" gorm:"column:updated_time;comment:最后更新时间"`
}

type ActiveRecord interface {
	schema.Tabler
	SetCreatedBy(createBy int64)
	SetUpdatedBy(updateBy int64)
	Generate() ActiveRecord
	GetId() interface{}
}
