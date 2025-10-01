package hoge

import (
	"time"

	"github.com/google/uuid"
)

type hogeDto struct {
	ID    uuid.UUID `gorm:"type:uuid;primaryKey"`
	Value string    `gorm:"not null"`
}

func (hogeDto) TableName() string { return "hoge" }

type cancelledHogeDto struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	Reason      string    `gorm:"not null"`
	CancelledAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`
}

func (cancelledHogeDto) TableName() string { return "cancelled_hoge" }

type piyoDto struct {
	ID    uuid.UUID `gorm:"type:uuid;primaryKey"`
	Value string    `gorm:"not null"`
}

func (piyoDto) TableName() string { return "piyo" }

type processedHogeDto struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	PiyoID      uuid.UUID `gorm:"type:uuid;not null"`
	ProcessedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`
}

func (processedHogeDto) TableName() string { return "processed_hoge" }
