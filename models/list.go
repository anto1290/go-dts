package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type List struct {
	ID         string `gorm:"primary_key " json:"id"`
	Title      string `gorm:"type:varchar(30)" json:"title"`
	Desciption string `gorm:"type:varchar(100)" json:"desciption"`
	Is_done    bool   `gorm:"type:boolean;default:false" json:"is_done"`
	// createAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"create_at"`
	// updateAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"update_at"`
}

func (list *List) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	list.ID = uuid.NewV4().String()
	return
}
