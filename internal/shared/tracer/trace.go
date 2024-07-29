package tracer

import (
	"context"
	"github.com/martirosharutyunyan/fraudnetic-test-task/internal/shared/request"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc/metadata"
)

type Constructor func(ctx *request.Context, name string) trace.Span

func NewProvider(ctx context.Context, collectorURL string) (*tracesdk.TracerProvider, error) {
	exporter, err := otlptrace.New(
		ctx,
		otlptracegrpc.NewClient(
			otlptracegrpc.WithInsecure(),
			otlptracegrpc.WithEndpoint(collectorURL),
		),
	)
	if err != nil {
		return nil, err
	}

	resources, err := resource.New(
		ctx,
		resource.WithAttributes(
			attribute.String("service.name", "main-service"),
			attribute.String("library.language", "go"),
		),
	)
	if err != nil {
		return nil, err
	}

	tp := tracesdk.NewTracerProvider(
		tracesdk.WithSampler(tracesdk.AlwaysSample()),
		tracesdk.WithBatcher(exporter),
		tracesdk.WithResource(resources),
	)

	return tp, nil
}

func NewOtel(otelName string) Constructor {
	return func(ctx *request.Context, name string) trace.Span {
		_, span := otel.Tracer(otelName).Start(ctx, name)
		md := metadata.New(map[string]string{"x-trace-id": span.SpanContext().TraceID().String()})
		ctx.Context = metadata.NewOutgoingContext(ctx.Context, md)

		return span
	}
}
