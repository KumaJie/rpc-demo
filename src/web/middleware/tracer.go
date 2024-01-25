package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"rpc-douyin/src/util/tracer"
)

const SpanCtx = "span-ctx"

func TracerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		span, ctx := tracer.NewSpanFromContext(context.Background(), c.FullPath())
		defer span.Finish()
		span.Tag("http.url", c.FullPath())
		span.Tag("http.method", c.Request.Method)
		c.Set(SpanCtx, ctx)
		c.Next()
	}
}
