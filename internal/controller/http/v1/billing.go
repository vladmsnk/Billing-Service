package v1

import (
	"vladmsnk/billing/internal/middlewares"

	"github.com/gin-gonic/gin"

	"vladmsnk/billing/internal/usecase"
	"vladmsnk/billing/pkg/logger"
)

type BillingRoutes struct {
	t usecase.Billing
	l logger.Interface
}

func newBillingRoutes(api *gin.RouterGroup, t usecase.Billing, l logger.Interface) {
	r := &BillingRoutes{t, l}

	secured := api.Group("/api/v1").Use(middlewares.Auth())
	{
		secured.POST("/activity", r.makePurchase)
		secured.GET("/balance", r.getBalance)
	}
}

func (r *BillingRoutes) makePurchase(c *gin.Context) {

}

func (r *BillingRoutes) getBalance(c *gin.Context) {

}
