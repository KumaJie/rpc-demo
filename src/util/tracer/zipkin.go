package tracer

import (
	"context"
	"github.com/openzipkin/zipkin-go"
	"github.com/openzipkin/zipkin-go/model"
	"github.com/openzipkin/zipkin-go/propagation/b3"
	zipkinhttp "github.com/openzipkin/zipkin-go/reporter/http"
	"google.golang.org/grpc/metadata"
)

var tracer *zipkin.Tracer

func InitTracer(serviceName, hostPort string) {
	reporter := zipkinhttp.NewReporter("http://localhost:9411/api/v2/spans")
	endpoint, err := zipkin.NewEndpoint(serviceName, hostPort)
	if err != nil {
		panic(err)
	}
	sampler := zipkin.NewModuloSampler(1)
	tracer, err = zipkin.NewTracer(
		reporter,
		zipkin.WithSampler(sampler),
		zipkin.WithLocalEndpoint(endpoint))
	if err != nil {
		panic(err)
	}
}

func NewSpan(spanName string) zipkin.Span {
	return tracer.StartSpan(spanName)
}

func NewSpanFromContext(ctx context.Context, spanName string) (zipkin.Span, context.Context) {
	return tracer.StartSpanFromContext(ctx, spanName)
}

func NewSpanFromSpanCtx(sc model.SpanContext, spanName string) zipkin.Span {
	return tracer.StartSpan(spanName, zipkin.Parent(sc))
}

func InjectGRPC(ctx model.SpanContext) context.Context {
	md := metadata.MD{}
	b3.InjectGRPC(&md)(ctx)
	return metadata.NewOutgoingContext(context.Background(), md)
}

func ExtractGRPC(md metadata.MD) (*model.SpanContext, error) {
	return b3.ExtractGRPC(&md)()
}
