package hoge

import (
	"context"
	"repository_pattern/domain/hoge"
	"repository_pattern/domain/primitives"
	"repository_pattern/infrastructure/pg"
	"testing"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestLoadHoge(t *testing.T) {
	tests := []struct {
		name   string
		id     string
		assert func(hoge.Hoge, *testing.T)
	}{
		{
			"unprocessed",
			"00000000-0000-0000-0000-000000000001",
			func(h hoge.Hoge, t *testing.T) {
				if v, ok := h.(*hoge.UnprocessedHoge); ok {
					if v.Id.String() != "00000000-0000-0000-0000-000000000001" {
						t.Error("unprocessed hoge, but invalid id")
					}
					if v.Value.String() != "unprocessed" {
						t.Error("unprocessed hoge, but invalid value")
					}
				} else {
					t.Error("not unprocessed hoge")
				}
			},
		},
		{
			"canncelled",
			"00000000-0000-0000-0000-000000000002",
			func(h hoge.Hoge, t *testing.T) {
				if v, ok := h.(*hoge.CanceledHoge); ok {
					if v.Id.String() != "00000000-0000-0000-0000-000000000002" {
						t.Errorf("cancelled hoge, but invalid id")
					}
					if v.Reason.String() != "reason_text" {
						t.Errorf("cancelled hoge, but invalid reason")
					}
				} else {
					t.Errorf("not cancelled hoge")
				}
			},
		},
		{
			"processed",
			"00000000-0000-0000-0000-000000000003",
			func(h hoge.Hoge, t *testing.T) {
				if v, ok := h.(*hoge.ProcessedHoge); ok {
					if v.Id.String() != "00000000-0000-0000-0000-000000000003" {
						t.Errorf("processed hoge, but invalid id")
					}
					if v.Piyo.Id.String() != "00000000-0000-0000-0000-000000000001" {
						t.Errorf("processed hoge, but invalid piyo id")
					}
					if v.Piyo.Value.String() != "piyopiyo" {
						t.Errorf("processed hoge, but invalid piyo value")
					}
				} else {
					t.Errorf("not processed hoge")
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()

			container, err := pg.WithPostgresContainer(ctx, "./load.fixture.sql", t)
			if err != nil {
				t.Fatalf("failed to start postgres container: %v", err)
			}
			defer container.Terminate()

			conn, err := gorm.Open(postgres.Open(container.DSN), &gorm.Config{})
			if err != nil {
				t.Fatalf("failed to connect to postgres: %v", err)
			}

			repository := NewGateway(conn)

			id, err := primitives.ParseIdString(tt.id)
			if err != nil {
				t.Fatalf("failed to parse id: %v", err)
			}

			hoge, err := repository.LoadHoge(id)
			if err != nil {
				t.Fatalf("failed to load hoge: %v", err)
			}

			tt.assert(hoge, t)
		})
	}
}
