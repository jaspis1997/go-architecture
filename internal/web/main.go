package web

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"playground/internal/app"
	"playground/internal/log"
	"playground/internal/repository"
	"sync"
	"syscall"

	"github.com/gin-gonic/gin"
)

type (
	HandlerFunc = gin.HandlerFunc
	Context     = *gin.Context
)

type Engine interface {
	gin.IRouter
	Run(host string, port uint16) error
}

type engine struct {
	*gin.Engine
}

func (e *engine) Run(host string, port uint16) error {
	addr := func() string {
		return fmt.Sprintf("%s:%d", host, port)
	}
	server := &http.Server{
		Addr:    addr(),
		Handler: e.Engine.Handler(),
	}
	var wg sync.WaitGroup
	go gracefulShutdown(server, &wg)
	err := server.ListenAndServe()
	if err != http.ErrServerClosed {
		return err
	}
	wg.Wait()
	return nil
}

func New() Engine {
	return &engine{gin.New()}
}

func gracefulShutdown(server *http.Server, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	<-sig

	log.Infof(LogFormatStartShutdown, TimeoutShutdown)

	ctx, cancel := context.WithTimeout(context.Background(), TimeoutShutdown)
	defer cancel()
	go forceShutdown(ctx)
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf(LogFormatFailedShutdown, err)
	}
	repository.Done()
	log.Info(LogMessageRepositoryProcessingCompleted)

	app.Done()
	log.Info(LogMessageApplicationProcessingCompleted)

	log.Info(LogMessageShutdownCompleted)
}

func forceShutdown(ctx context.Context) {
	<-ctx.Done()
	log.Fatal(LogMessageForceShutdown)
}
