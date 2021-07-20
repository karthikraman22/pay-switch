package api

import (
	"net/http"

	"achuala.in/pay-switch/ep"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type EndpointResource struct {
	RouteEngine *gin.Engine
	epMgr       *ep.EndpointMgr
	logger      *zap.Logger
}

func NewEndpointResource(r *gin.Engine, epMgr *ep.EndpointMgr, logger *zap.Logger) *EndpointResource {
	epr := &EndpointResource{RouteEngine: r, epMgr: epMgr, logger: logger}
	epr.addV1Routes()
	return epr
}

func (er *EndpointResource) addV1Routes() {
	v1 := er.RouteEngine.Group("/v1")
	{
		v1.POST("/endpoint/create", er.createEndpoint)
	}
}
func (er *EndpointResource) createEndpoint(c *gin.Context) {
	var epCfg ep.EndpointCfg
	if err := c.ShouldBindJSON(&epCfg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if epCfg.Type == "server" {
		er.logger.Info("creating a server endpoint", zap.Any("cfg", epCfg))
		ep, err := er.epMgr.NewServerEndpoint(&epCfg)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, ep)
	} else if epCfg.Type == "client" {
		er.logger.Info("creating a client endpoint", zap.Any("cfg", epCfg))
		ep, err := er.epMgr.NewClientEndpoint(&epCfg)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, ep)
	}

}
