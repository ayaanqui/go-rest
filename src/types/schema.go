package types

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	gorm.Model
	ID uuid.UUID `gorm:"type:char(36); primary key" json:"id,omitempty"`
	CreatedAt time.Time `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at,omitempty"`
}

func (base *Base) BeforeCreate(tx *gorm.DB) error {
	base.ID = uuid.New()
	base.CreatedAt = time.Now()
	base.UpdatedAt = time.Now()
	return nil
}

type Post struct {
	Base
	Title string `gorm:"not null" json:"title,omitempty"`
	Slug string `gorm:"not null" json:"slug,omitempty"`
	Content string `gorm:"type:text; not null" json:"content,omitempty"`
	Summary string `gorm:"not null; default: ''" json:"summary,omitempty"`
}

type Home struct {
	Base
	Message string `gorm:"type:text; not null" json:"message,omitempty"`
}