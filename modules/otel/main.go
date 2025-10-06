package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
)

func initTracer() (func(), error) {
	ctx := context.Background()

	exporter, err := otlptracegrpc.New(ctx,
		otlptracegrpc.WithEndpoint("otel-lgtm:4317"),
		otlptracegrpc.WithInsecure(),
	)
	if err != nil {
		slog.Error("failed to create OLTP trace exporter", "error", err)
		return nil, err
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("gin-app"),
		)),
	)
	otel.SetTracerProvider(tp)

	return func() {
		_ = tp.Shutdown(ctx)
	}, nil
}

func initMetrics() (func(), error) {
	ctx := context.Background()

	exporter, err := otlpmetricgrpc.New(ctx,
		otlpmetricgrpc.WithEndpoint("otel-lgtm:4317"),
		otlpmetricgrpc.WithInsecure(),
	)
	if err != nil {
		slog.Error("failed to create OTLP metric exporter", "error", err)
		return nil, err
	}

	provider := sdkmetric.NewMeterProvider(
		sdkmetric.WithReader(sdkmetric.NewPeriodicReader(exporter, sdkmetric.WithInterval(time.Second*10))),
		sdkmetric.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("gin-app"),
		)),
	)
	otel.SetMeterProvider(provider)

	return func() {
		_ = provider.Shutdown(ctx)
	}, nil
}

func initLogger() {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: true})
	jsonLogger := slog.New(handler)
	slog.SetDefault(jsonLogger)
}

func main() {
	initLogger()

	traceShutdown, err := initTracer()
	if err != nil {
		slog.Error("fuck", "error", err)
		return
	}
	defer traceShutdown()

	metricShutdown, err := initMetrics()
	if err != nil {
		slog.Error("fuck", "error", err)
		return
	}
	defer metricShutdown()

	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(otelgin.Middleware("gin-app"))

	r.GET("/ping", func(c *gin.Context) {
		slog.Info("ping")
		ctx := c.Request.Context()
		time.Sleep(100 * time.Millisecond)
		res := nest(ctx)
		c.JSON(200, res)
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	slog.Info("Start server :8000")
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		slog.Error("server failed", "error", err)
	}
}

func nest(ctx context.Context) map[string]any {
	tr := otel.Tracer("gin-app/service")
	_, s := tr.Start(ctx, "nest")
	defer s.End()

	slog.Info("nest")
	s.SetAttributes(attribute.String("hoge", "piyo"), attribute.Int64("fuga", 10))
	s.AddEvent("foo")
	time.Sleep(200 * time.Millisecond)

	return gin.H{"message": "pong"}
}
