package hoge

import (
	"errors"
	"repository_pattern/domain/hoge"
	"repository_pattern/domain/primitives"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (g *Gateway) LoadHoge(id primitives.Id) (hoge.Hoge, error) {
	id_uuid := uuid.UUID(id)

	type result struct {
		hogeDto
		CancelledReason *string    `gorm:"column:cancelled_reason"`
		CancelledAt     *time.Time `gorm:"column:cancelled_at"`
		ProcessedAt     *time.Time `gorm:"column:processed_at"`
		PiyoID          *uuid.UUID `gorm:"column:piyo_id"`
		PiyoValue       *string    `gorm:"column:piyo_value"`
	}

	var res result
	err := g.conn.Table("hoge h").
		Select("h.*, ch.reason as cancelled_reason, ch.cancelled_at, ph.processed_at, ph.piyo_id, p.value as piyo_value").
		Joins("LEFT JOIN cancelled_hoge ch ON h.id = ch.id").
		Joins("LEFT JOIN processed_hoge ph ON h.id = ph.id").
		Joins("LEFT JOIN piyo p ON ph.piyo_id = p.id").
		Where("h.id = ?", id_uuid).
		First(&res).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("hoge not found")
		}
		return nil, err
	}

	if res.CancelledReason != nil && res.ProcessedAt != nil {
		if res.CancelledAt.After(*res.ProcessedAt) {
			return hoge.NewCanceledHoge(id_uuid[:], *res.CancelledReason)
		} else {
			piyoId := primitives.NewId(*res.PiyoID)
			piyo, err := hoge.NewPiyo(piyoId[:], *res.PiyoValue)
			if err != nil {
				return nil, err
			}
			return hoge.NewProcessedHoge(id_uuid[:], piyo)
		}
	}

	if res.CancelledReason != nil {
		return hoge.NewCanceledHoge(id_uuid[:], *res.CancelledReason)
	}

	if res.ProcessedAt != nil {
		piyoId := primitives.NewId(*res.PiyoID)
		piyo, err := hoge.NewPiyo(piyoId[:], *res.PiyoValue)
		if err != nil {
			return nil, err
		}
		return hoge.NewProcessedHoge(id_uuid[:], piyo)
	}

	return hoge.NewUnprocessedHoge(id_uuid[:], res.Value)
}
