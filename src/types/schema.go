package types

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	ID uuid.UUID `gorm:"type:char(36); primary key" json:"id"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at"`
}

func (base *Base) BeforeCreate(tx *gorm.DB) error {
	base.ID = uuid.New()
	base.CreatedAt = time.Now()
	base.UpdatedAt = time.Now()
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

type User struct {
	Base
	Username string `gorm:"varchar(30); not null; index; unique" json:"username"`
	Email string `gorm:"varchar(255); not null; index; unique" json:"email"`
	// Ignore field from json output
	Password string `gorm:"varchar(255); not null; index" json:"-"`
	IsAdmin bool `gorm:"default: false" json:"is_admin"`
	IsActive bool `gorm:"default: false" json:"is_active"`
}