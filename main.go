package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"achuala.in/pay-switch/api"
	"github.com/gin-gonic/gin"
	_ "go.uber.org/automaxprocs"
	"go.uber.org/zap"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	logger, _ := zap.NewDevelopment()

	r := gin.Default()

	api.NewEndpointResource(r, logger)

	srv := &http.Server{
		Addr:    ":9090",
		Handler: r,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			logger.Fatal("listen", zap.Error(err))
		}
	}()

	/*
		iso8583Server := iso8583.NewIso8583Server("tcp://:9000", logger)

		go func() {
			if err := gnet.Serve(iso8583Server, iso8583Server.Addr,
				gnet.WithMulticore(true),
				gnet.WithCodec(iso8583Server.Codec),
				gnet.WithTCPKeepAlive(time.Second*30),
				gnet.WithSocketRecvBuffer(8*1024),
				gnet.WithSocketSendBuffer(8*1024),
				gnet.WithReusePort(true),
				gnet.WithLogger(logger.Sugar()),
				gnet.WithLogLevel(zapcore.DebugLevel)); err != nil {
				logger.Fatal("server start failed", zap.Error(err))
			}
		}()
	*/
	// Register for shutdown events and handle it gracefully
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT)
	<-done
	logger.Info("shutdown initiated...")
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("server forced to shutdown : ", err)
	}
	logger.Info("server shutdown")
}
