package handlers

import (
	"github.com/akuera/go-template/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getDomainHandler(service domain.Int) func(*gin.Context) {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		id := c.Param("id")
		if len(id) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "missing/invalid id passed",
			})
			return
		}
		domainResp, err := service.GetDomain(ctx, id)
		c.JSON(httpStatus(err), domainResp)
	}
}
