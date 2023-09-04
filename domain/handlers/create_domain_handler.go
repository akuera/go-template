package handlers

import (
	"github.com/akuera/go-template/domain"
	"github.com/akuera/go-template/domain/handlers/requests"
	"github.com/akuera/go-template/domain/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func createDomainHandler(service domain.Int) func(*gin.Context) {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		var r requests.CreateDomain
		err := c.ShouldBind(&r)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "failed to bind to Domain.Domain",
			})
			return
		}
		if err := r.Validate(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		domainResp, err := service.CreateDomain(ctx, &models.Domain{
			Name: r.Name,
		})
		c.JSON(httpStatus(err), domainResp)
	}
}
