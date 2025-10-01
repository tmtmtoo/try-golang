package hoge

import (
	"repository_pattern/domain/hoge"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (g *Gateway) VisitUnprocessedHoge(h *hoge.UnprocessedHoge) error {
	hogeDto := hogeDto{
		ID:    uuid.UUID(h.Id),
		Value: string(h.Value),
	}
	return g.conn.Save(&hogeDto).Error
}

func (g *Gateway) VisitCanceledHoge(h *hoge.CanceledHoge) error {
	cancelledDto := cancelledHogeDto{
		ID:     uuid.UUID(h.Id),
		Reason: string(h.Reason),
	}
	return g.conn.Save(&cancelledDto).Error
}

func (g *Gateway) VisitProcessedHoge(h *hoge.ProcessedHoge) error {
	return g.conn.Transaction(func(tx *gorm.DB) error {
		piyoDto := piyoDto{
			ID:    uuid.UUID(h.Piyo.Id),
			Value: string(h.Piyo.Value),
		}
		if err := tx.Save(&piyoDto).Error; err != nil {
			return err
		}

		processedDto := processedHogeDto{
			ID:     uuid.UUID(h.Id),
			PiyoID: uuid.UUID(h.Piyo.Id),
		}
		return tx.Save(&processedDto).Error
	})
}

func (g *Gateway) PersistHoge(hoge hoge.Hoge) error {
	return hoge.Accept(g)
}
