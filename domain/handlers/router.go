package handlers

import (
	"github.com/akuera/go-template/domain"
	"github.com/gin-gonic/gin"
)

func AddDomainHandlers(rg *gin.RouterGroup, service domain.Int) {
	v1 := rg.Group("v1")

	v1.GET("/domains/:id", getDomainHandler(service))
	v1.GET("/domains", getDomainsHandler(service))
	v1.PATCH("/domains/:id", updateDomainHandler(service))
	v1.POST("/domains", createDomainHandler(service))
	v1.DELETE("/domains/:id", deleteDomainHandler(service))
}
