package types

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	gorm.Model
	ID uuid.UUID `gorm:"type:char(36); primary key" json:"id"`
	CreatedAt string `gorm:"not null" json:"created_at"`
	UpdatedAt string `gorm:"not null" json:"updated_at"`
}

func (base *Base) BeforeCreate(tx *gorm.DB) error {
	base.ID = uuid.New()
	base.CreatedAt = time.Now().String()
	base.UpdatedAt = time.Now().String()
	return nil
}

type Post struct {
	Base
	Title string `gorm:"not null" json:"title"`
	Slug string `gorm:"not null" json:"slug"`
	Content string `gorm:"type:text; not null" json:"content"`
	Summary string `gorm:"not null; default: ''" json:"summary"`
}

type Home struct {
	Base
	Message string `gorm:"type:text; not null" json:"message"`
}