package models

import "time"

// Migration representa uma migration executada no banco
type Migration struct {
	ID        uint      `gorm:"primaryKey"`
	Version   string    `gorm:"size:255;unique;not null"`
	Name      string    `gorm:"size:255;not null"`
	ExecutedAt time.Time `gorm:"autoCreateTime"`
} 