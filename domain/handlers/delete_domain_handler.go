package handlers

import (
	"github.com/akuera/go-template/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

func deleteDomainHandler(service domain.Int) func(*gin.Context) {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		id := c.Param("id")
		if len(id) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid id passed; must be hex",
			})
			return
		}
		err := service.DeleteDomain(ctx, id)
		c.JSON(httpStatus(err), nil)
	}
}
