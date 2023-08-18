package caller

import (
	"github.com/gin-gonic/gin"
	"reason-im/internal/utils/logger"
)

func Call[A, B any](
	function func(req *A) *B,
	c *gin.Context,
	req *A,
) {
	if err := c.BindJSON(&req); err != nil {
		logger.Warn(c, "bind req has failed", "req", req)
		return
	}
	data := function(req)
	c.JSON(200, data)
}