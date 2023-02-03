package bootstrap

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	deliverHttp "github.com/krobus00/go-test-service/internal/delivery/http"
	"github.com/krobus00/go-test-service/internal/infrastructure"
	"github.com/krobus00/go-test-service/internal/usecase"
)

type ServerBootstrap struct{}

func NewServerBootstrap() *ServerBootstrap {
	return new(ServerBootstrap)
}

func (b *ServerBootstrap) Execute() {
	infra := infrastructure.NewInfrastructure()
	err := infra.LoadConfig()
	if err != nil {
		log.Fatal(err.Error())
	}
	infra.NewRouter()

	e := infra.Router

	// init usecases
	simpleMathUsecase := usecase.NewSimpleMathUsecase()

	// init handler
	simpleMathHandler := deliverHttp.NewSimpleMathHandler()
	simpleMathHandler.RegisterSimpleMathUecase(simpleMathUsecase)

	// init delivery
	httpHandler := deliverHttp.NewHttpHandler()
	httpHandler.RegisterEcho(infra.Router)
	httpHandler.RegisterSimpleMathHandler(simpleMathHandler)

	httpHandler.RegisterHttpRoute()

	// Start server
	go func() {
		log.Println("running server on port", infra.Config.App.Port)
		if err := e.Start(fmt.Sprintf(":%d", infra.Config.App.Port)); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

}
