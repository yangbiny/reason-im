package utils

import (
	"context"
	"github.com/gin-gonic/gin"
)

func Call[A, C any, B struct{}](
	rpc func(client C, ctx context.Context, req *A) (*B, error),
	client C,
	c *gin.Context,
) *B {
	var req A
	if err := c.BindJSON(&req); err != nil {
		Warn(c, "bind req has failed", "req", req)
		return nil
	}
	data, err := rpc(client, c, &req)
	if err != nil {
		return nil
	}
	return data
}
