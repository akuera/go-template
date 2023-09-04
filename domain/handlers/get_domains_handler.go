package handlers

import (
	"github.com/akuera/go-template/domain"
	"github.com/gin-gonic/gin"
)

func getDomainsHandler(service domain.Int) func(*gin.Context) {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		domainResp, err := service.GetDomains(ctx)
		c.JSON(httpStatus(err), domainResp)
	}
}
