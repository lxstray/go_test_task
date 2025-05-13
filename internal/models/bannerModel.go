package models

import "github.com/google/uuid"

type Banner struct {
	ID      uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name    string    `gorm:"type:varchar(255);not null"`
	Image   string    `gorm:"type:text;not null"`
	CPM     float64   `gorm:"type:numeric(10,2);not null;check:cpm >= 0 AND cpm <= 1000"`
	Geo     string    `gorm:"type:varchar(2);not null"`
	Feature int       `gorm:"type:integer;not null;check:feature >= 0 AND feature <= 100"`
}
