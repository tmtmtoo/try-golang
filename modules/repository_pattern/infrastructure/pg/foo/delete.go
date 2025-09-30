package foo

import (
	"repository_pattern/domain/primitives"

	"gorm.io/gorm"
)

func (g *Gateway) DeleteFoo(id primitives.Id) error {
	return g.conn.Transaction(func(tx *gorm.DB) error {
		fooID := id.Bytes()

		if err := tx.Where("foo_id = ?", fooID).Delete(&fooBazDto{}).Error; err != nil {
			return err
		}

		if err := tx.Where("foo_id = ?", fooID).Delete(&barDto{}).Error; err != nil {
			return err
		}

		if err := tx.Where("id = ?", fooID).Delete(&fooDto{}).Error; err != nil {
			return err
		}

		return nil
	})
}
