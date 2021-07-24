package v1

import (
	"net/http"

	"achuala.in/payswitch/core"
	"achuala.in/payswitch/ep"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PaymentResource struct {
	ginEngine *gin.Engine
	router    *core.Router
	logger    *zap.Logger
}

func NewPaymentResource(re *gin.Engine, router *core.Router, logger *zap.Logger) *PaymentResource {
	pr := &PaymentResource{ginEngine: re, router: router, logger: logger}
	pr.addV1Routes()
	return pr
}

func (pr *PaymentResource) addV1Routes() {
	v1 := pr.ginEngine.Group("/v1")
	{
		v1.POST("/payment/collect", pr.collectPayment)
	}
}
func (pr *PaymentResource) collectPayment(c *gin.Context) {
	var epCfg ep.EndpointCfg
	if err := c.ShouldBindJSON(&epCfg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, nil)
}
