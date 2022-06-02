package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/braintree/manners"
	"github.com/fakovacic/users-service/cmd/users/config"
	"github.com/fakovacic/users-service/cmd/users/config/integrations"
	"github.com/fakovacic/users-service/internal/users"
	"github.com/fakovacic/users-service/internal/users/handlers/health"
	"github.com/fakovacic/users-service/internal/users/handlers/http/middleware"
	svcMiddleware "github.com/fakovacic/users-service/internal/users/middleware"
	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
)

const errorChan int = 10

func main() {
	c, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	dbStore, err := config.NewStore(c)
	if err != nil {
		log.Fatal(err)
	}

	notifier := integrations.NewNotifier(c)

	service := users.New(c, dbStore, time.Now, uuid.New)
	service = svcMiddleware.NewNotificationMiddleware(service, notifier)
	service = svcMiddleware.NewLoggingMiddleware(service, c)

	h := config.NewHandlers(c, service)

	router := httprouter.New()

	router.GET("/users", h.List)
	router.POST("/users", h.Create)
	router.PATCH("/users/:id", h.Update)
	router.DELETE("/users/:id", h.Delete)

	var (
		httpAddr   = "0.0.0.0:8080"
		healthAddr = "0.0.0.0:8081"
	)

	httpServer := manners.NewServer()
	httpServer.Addr = httpAddr
	httpServer.Handler = middleware.ReqID(
		middleware.Logger(
			c, router,
		),
	)

	errChan := make(chan error, errorChan)

	healthServer := health.StartServer(healthAddr)

	go func() {
		c.Log.Debug().Msgf("Health service listening on %s", healthAddr)
		errChan <- healthServer.ListenAndServe()
	}()
	go func() {
		c.Log.Debug().Msgf("HTTP service listening on %s", httpAddr)
		errChan <- httpServer.ListenAndServe()
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case err := <-errChan:
			if err != nil {
				c.Log.Fatal().Msg(err.Error())
			}
		case s := <-signalChan:
			c.Log.Debug().Msgf(fmt.Sprintf("Captured %v. Exiting...", s))
			health.SetHealthStatus(http.StatusServiceUnavailable)
			httpServer.BlockingClose()
			os.Exit(0)
		}
	}
}
