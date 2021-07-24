package main

import (
	"context"
	"errors"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"achuala.in/payswitch/api"
	v1 "achuala.in/payswitch/api/v1"
	"achuala.in/payswitch/core"
	"achuala.in/payswitch/ep"
	"github.com/gin-gonic/gin"
	_ "go.uber.org/automaxprocs"
	"go.uber.org/zap"
)

func main() {

	logger, _ := zap.NewDevelopment()

	r := gin.Default()

	epMgr := ep.NewEndpointMgr(logger)
	router := core.NewRouter("localhost:29092")
	api.NewEndpointResource(r, epMgr, logger)
	v1.NewPaymentResource(r, router, logger)

	srv := &http.Server{
		Addr:    ":9090",
		Handler: r,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Fatal("listen", zap.Error(err))
		}
	}()

	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Listen for the interrupt signal.
	<-ctx.Done()

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Notify the endpoint manager to stop all the endpoints
	epMgr.Shutdown(ctx)

	if err := srv.Shutdown(ctx); err != nil {
		logger.Info("server forced to shutdown : ", zap.Error(err))
	}

	logger.Info("server shutdown")
}
