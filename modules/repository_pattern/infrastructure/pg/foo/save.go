package foo

import (
	"errors"
	"repository_pattern/domain/foo"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (g *Gateway) SaveFoo(f *foo.Foo) error {
	if g == nil || g.conn == nil {
		return errors.New("gateway or connection is nil")
	}

	fooID := uuid.UUID(f.Id)

	return g.conn.Transaction(func(tx *gorm.DB) error {
		var existingFoo fooDto
		err := tx.First(&existingFoo, "id = ?", fooID).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				if err := tx.Create(&fooDto{ID: fooID, Value: string(f.Value)}).Error; err != nil {
					return err
				}
			} else {
				return err
			}
		} else {
			if err := tx.Model(&fooDto{}).Where("id = ?", fooID).Update("value", string(f.Value)).Error; err != nil {
				return err
			}
		}

		var currentBarDtos []barDto
		if err := tx.Find(&currentBarDtos, "foo_id = ?", fooID).Error; err != nil {
			return err
		}
		currentBarMap := make(map[uuid.UUID]barDto, len(currentBarDtos))
		for _, b := range currentBarDtos {
			currentBarMap[b.ID] = b
		}
		newBarIds := make(map[uuid.UUID]struct{}, len(f.Bars))
		for _, b := range f.Bars {
			bid := uuid.UUID(b.Id)
			newBarIds[bid] = struct{}{}
			if _, ok := currentBarMap[bid]; ok {
				if err := tx.Model(&barDto{}).Where("id = ?", bid).Updates(map[string]any{"value": string(b.Value)}).Error; err != nil {
					return err
				}
			} else {
				if err := tx.Create(&barDto{ID: bid, Value: string(b.Value), FooID: fooID}).Error; err != nil {
					return err
				}
			}
		}

		var toDeleteBars []uuid.UUID
		for id := range currentBarMap {
			if _, still := newBarIds[id]; !still {
				toDeleteBars = append(toDeleteBars, id)
			}
		}
		if len(toDeleteBars) > 0 {
			if err := tx.Delete(&barDto{}, "id IN ?", toDeleteBars).Error; err != nil {
				return err
			}
		}

		var currentLinks []fooBazDto
		if err := tx.Find(&currentLinks, "foo_id = ?", fooID).Error; err != nil {
			return err
		}
		currentBazSet := make(map[uuid.UUID]struct{}, len(currentLinks))
		for _, l := range currentLinks {
			currentBazSet[l.BazID] = struct{}{}
		}

		desiredBazSet := make(map[uuid.UUID]foo.Baz, len(f.Bazs))
		for _, bz := range f.Bazs {
			bzID := uuid.UUID(bz.Id)
			desiredBazSet[bzID] = bz
			var count int64
			if err := tx.Model(&bazDto{}).Where("id = ?", bzID).Count(&count).Error; err != nil {
				return err
			}
			if count == 0 {
				if err := tx.Create(&bazDto{ID: bzID, Value: string(bz.Value)}).Error; err != nil {
					return err
				}
			} else {
				if err := tx.Model(&bazDto{}).Where("id = ?", bzID).Update("value", string(bz.Value)).Error; err != nil {
					return err
				}
			}
			if _, linked := currentBazSet[bzID]; !linked {
				if err := tx.Create(&fooBazDto{FooID: fooID, BazID: bzID}).Error; err != nil {
					return err
				}
			}
		}

		var removeLinks []uuid.UUID
		for bzID := range currentBazSet {
			if _, want := desiredBazSet[bzID]; !want {
				removeLinks = append(removeLinks, bzID)
			}
		}
		if len(removeLinks) > 0 {
			if err := tx.Delete(&fooBazDto{}, "foo_id = ? AND baz_id IN ?", fooID, removeLinks).Error; err != nil {
				return err
			}
		}

		return nil
	})
}
