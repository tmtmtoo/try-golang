package foo

import (
	"github.com/google/uuid"
)

type fooDto struct {
	ID    uuid.UUID `gorm:"type:uuid;primaryKey"`
	Value string    `gorm:"not null"`
}

func (fooDto) TableName() string { return "foo" }

type barDto struct {
	ID    uuid.UUID `gorm:"type:uuid;primaryKey"`
	Value string    `gorm:"not null"`
	FooID uuid.UUID `gorm:"type:uuid;not null"`
}

func (barDto) TableName() string { return "bar" }

type bazDto struct {
	ID    uuid.UUID `gorm:"type:uuid;primaryKey"`
	Value string    `gorm:"not null"`
}

func (bazDto) TableName() string { return "baz" }

type fooBazDto struct {
	FooID uuid.UUID `gorm:"type:uuid;not null"`
	BazID uuid.UUID `gorm:"type:uuid;not null"`
}

func (fooBazDto) TableName() string { return "foo_baz" }
