package foo

import (
	"repository_pattern/domain/foo"

	"github.com/google/uuid"
)

func (g *Gateway) FindFooById(id foo.Id) (*foo.Foo, error) {
	var fooDto fooDto
	if err := g.conn.First(&fooDto, "id = ?", uuid.UUID(id)).Error; err != nil {
		return nil, err
	}

	var barDtos []barDto
	if err := g.conn.Find(&barDtos, "foo_id = ?", uuid.UUID(id)).Error; err != nil {
		return nil, err
	}

	bars := make([]foo.Bar, len(barDtos))
	for i, barDto := range barDtos {
		bar, err := foo.NewBar(barDto.ID[:], barDto.Value)
		if err != nil {
			return nil, err
		}
		bars[i] = *bar
	}

	var fooBazDtos []fooBazDto
	if err := g.conn.Find(&fooBazDtos, "foo_id = ?", uuid.UUID(id)).Error; err != nil {
		return nil, err
	}

	bazIds := make([]uuid.UUID, len(fooBazDtos))
	for i, fooBazDto := range fooBazDtos {
		bazIds[i] = fooBazDto.BazID
	}

	var bazDtos []bazDto
	if len(bazIds) > 0 {
		if err := g.conn.Find(&bazDtos, "id IN ?", bazIds).Error; err != nil {
			return nil, err
		}
	}

	bazs := make([]foo.Baz, len(bazDtos))
	for i, bazDto := range bazDtos {
		baz, err := foo.NewBaz(bazDto.ID[:], bazDto.Value)
		if err != nil {
			return nil, err
		}
		bazs[i] = *baz
	}

	return foo.NewFoo(fooDto.ID[:], fooDto.Value, bars, bazs)
}
