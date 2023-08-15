package utils

import "github.com/gin-gonic/gin"

func Call[A, B any](
	function func(req *A) *B,
	c *gin.Context,
	req *A,
) {
	if err := c.BindJSON(&req); err != nil {
		Warn(c, "bind req has failed", "req", req)
		return
	}
	data := function(req)
	c.JSON(200, data)
}
