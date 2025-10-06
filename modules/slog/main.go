package main

import (
	"context"
	"log/slog"
	"os"
)

type CustomHandler struct {
	handler slog.Handler
}

func (c *CustomHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return c.handler.Enabled(ctx, level)
}

func (c *CustomHandler) Handle(ctx context.Context, r slog.Record) error {
	if v, ok := ctx.Value("customKey").(string); ok {
		r.AddAttrs(slog.String("customKey", v))
	}
	return c.handler.Handle(ctx, r)
}

func (c *CustomHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return c.handler.WithAttrs(attrs)
}

func (c *CustomHandler) WithGroup(name string) slog.Handler {
	return c.handler.WithGroup(name)
}

type Masked string

func (s Masked) LogValue() slog.Value {
	return slog.StringValue("XXXXXXXXXXXXXXX")
}

func main() {
	handler := &CustomHandler{slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == "foo" {
				return slog.String("foo", "masked")
			}
			return a
		},
	})}
	jsonLogger := slog.New(handler)
	slog.SetDefault(jsonLogger)

	ctx := context.Background()
	ctx = context.WithValue(ctx, "customKey", "what")
	slog.InfoContext(ctx, "wow", "aaa", 10)
	slog.Info("bbb", slog.String("hoge", "piyo"), slog.Int("foo", 100), slog.Any("struct", &struct {
		Value int
	}{
		10,
	}))
	slog.Info("masked", slog.Any("dounaru", Masked("hogehoge")))
}
