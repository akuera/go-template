package handlers

import (
	"github.com/akuera/go-template/domain"
	"github.com/akuera/go-template/domain/handlers/requests"
	"github.com/akuera/go-template/domain/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func updateDomainHandler(service domain.Int) func(*gin.Context) {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		id := c.Param("id")
		if len(id) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid id passed; must be hex",
			})
			return
		}
		var r requests.UpdateDomain
		err := c.ShouldBind(&r)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "failed to bind to domain.Domain",
			})
			return
		}
		domainResp, err := service.UpdateDomain(ctx, id, &models.Domain{
			Name: r.Name,
		})
		c.JSON(httpStatus(err), domainResp)
	}
}
