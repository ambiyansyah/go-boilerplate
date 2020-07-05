package server

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/ambiyansyah/go-boilerplate/pkg/config"
	"github.com/ambiyansyah/go-boilerplate/pkg/ping/handler"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

// HTTPServer ...
type HTTPServer interface {
	Start()
}

type httpServer struct {
	config *config.AppConfig
}

// NewHTTPServer ...
func NewHTTPServer() HTTPServer {
	appConfig, err := config.SetAppConfig()
	if err != nil {
		panic(err.Error())
	}

	return &httpServer{
		config: appConfig,
	}
}

func (h *httpServer) Start() {
	e := echo.New()
	e.HideBanner = true
	e.Server.Addr = h.config.Port

	// Logging
	e.Logger.SetLevel(log.INFO)

	// Routing
	r := e.Group("/v1")
	handler.NewHTTPHandler(h.config, r).Mount()

	// Start server
	go func() {
		if err := e.Start(e.Server.Addr); err != nil {
			e.Logger.Info("Shutting down the server")
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
