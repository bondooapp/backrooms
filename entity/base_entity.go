package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// BaseEntity
//
// BaseEntity of database model.
type BaseEntity struct {
	ID         string    `json:"id" gorm:"primaryKey"`
	CreateBy   string    `json:"createBy" gorm:"default:NULL"`
	CreateTime time.Time `json:"createTime" gorm:"default:NULL"`
	UpdateBy   string    `json:"updateBy" gorm:"default:NULL"`
	UpdateTime time.Time `json:"updateTime" gorm:"default:NULL"`
	Version    int       `json:"version" gorm:"default:1"`
	OwnerID    string    `json:"ownerId" gorm:"default:NULL"`
	EntryType  string    `json:"entryType" gorm:"default:NULL"`
}

// BeforeCreate
//
// Execute this method before create.
func (e *BaseEntity) BeforeCreate(*gorm.DB) error {
	u7 := uuid.Must(uuid.NewV7()).String()
	e.ID = u7
	e.CreateTime = time.Now().UTC()
	return nil
}

// BeforeUpdate
//
// Execute this method before update.
func (e *BaseEntity) BeforeUpdate(tx *gorm.DB) (err error) {
	tx.Statement.SetColumn("UpdateTime", time.Now().UTC())
	return nil
}
