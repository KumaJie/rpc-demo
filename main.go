package main

import (
	"context"
	"fmt"
	"github.com/openzipkin/zipkin-go/propagation/b3"
	"google.golang.org/grpc/metadata"
	"rpc-douyin/src/util/tracer"
)

func main() {
	tracer.InitTracer("aa", "127.0.0.1:20")
	span, _ := tracer.NewSpanFromContext(context.Background(), "asd")
	md := metadata.MD{}
	i := b3.InjectGRPC(&md)
	i(span.Context())
	fmt.Println(md)
}
