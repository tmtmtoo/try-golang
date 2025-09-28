package foo

import (
	"gorm.io/gorm"
)

type Gateway struct {
	conn *gorm.DB
}

func NewGateway(conn *gorm.DB) *Gateway {
	return &Gateway{
		conn,
	}
}
