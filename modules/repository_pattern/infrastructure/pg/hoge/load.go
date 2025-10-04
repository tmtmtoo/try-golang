package hoge

import (
	"errors"
	"repository_pattern/domain/hoge"
	"repository_pattern/domain/primitives"
	"time"

	"github.com/google/uuid"
)

func (g *Gateway) LoadHoge(id primitives.Id) (hoge.Hoge, error) {
	type result struct {
		ID              uuid.UUID  `gorm:"column:id"`
		Value           string     `gorm:"column:value"`
		CancelledReason *string    `gorm:"column:cancelled_reason"`
		CancelledAt     *time.Time `gorm:"column:cancelled_at"`
		ProcessedAt     *time.Time `gorm:"column:processed_at"`
		PiyoID          *uuid.UUID `gorm:"column:piyo_id"`
		PiyoValue       *string    `gorm:"column:piyo_value"`
	}

	var res result
	query := `
		SELECT
			h.id, 
			h.value,
			ch.reason as cancelled_reason,
			ch.cancelled_at,
			ph.processed_at,
			ph.piyo_id,
			p.value as piyo_value
		FROM public.hoge as h
		LEFT JOIN cancelled_hoge as ch ON h.id = ch.id
		LEFT JOIN processed_hoge as ph ON h.id = ph.id
		LEFT JOIN piyo as p ON ph.piyo_id = p.id
		WHERE h.id = ?
		LIMIT 1
	`
	err := g.conn.Raw(query, id.Bytes()).Scan(&res).Error

	if err != nil {
		return nil, err
	}
	if res.ID == uuid.Nil {
		return nil, errors.New("hoge not found")
	}

	if res.CancelledReason != nil && res.ProcessedAt != nil {
		if res.CancelledAt.After(*res.ProcessedAt) {
			return hoge.NewCanceledHoge(res.ID[:], *res.CancelledReason)
		} else {
			piyoId := primitives.NewId(*res.PiyoID)
			piyo, err := hoge.NewPiyo(piyoId[:], *res.PiyoValue)
			if err != nil {
				return nil, err
			}
			return hoge.NewProcessedHoge(res.ID[:], piyo)
		}
	}

	if res.CancelledReason != nil {
		return hoge.NewCanceledHoge(res.ID[:], *res.CancelledReason)
	}

	if res.ProcessedAt != nil {
		piyoId := primitives.NewId(*res.PiyoID)
		piyo, err := hoge.NewPiyo(piyoId[:], *res.PiyoValue)
		if err != nil {
			return nil, err
		}
		return hoge.NewProcessedHoge(res.ID[:], piyo)
	}

	return hoge.NewUnprocessedHoge(res.ID[:], res.Value)
}
