package api

import (
	"github.com/Insperias/l2_golang/develop/dev_11/go-calendar/api/handler"
	"github.com/Insperias/l2_golang/develop/dev_11/go-calendar/api/middleware"
	"github.com/Insperias/l2_golang/develop/dev_11/go-calendar/internal/domain"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"net"
	"net/http"
	"time"
)

// ServerConfig Struct contains general otions for HTTP server
type ServerConfig struct {
	Host           string
	Port           string
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	MaxHeaderBytes int
}

func initRouter(h *handler.Handler) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", h.Handle(h.MainHandler))
	router.HandleFunc("/create_event", h.Handle(h.EventCreateHandler)).Methods(http.MethodPost)
	router.HandleFunc("/update_event", h.Handle(h.EventUpdateHandler)).Methods(http.MethodPost)
	router.HandleFunc("/delete_event", h.Handle(h.EventDeleteHandler)).Methods(http.MethodPost)
	router.HandleFunc("/events_for_day", h.Handle(h.EventsForDayHandler)).Methods(http.MethodGet)
	router.HandleFunc("/events_for_week", h.Handle(h.EventsForWeekHandler)).Methods(http.MethodGet)
	router.HandleFunc("/events_for_month", h.Handle(h.EventsForMonthHandler)).Methods(http.MethodGet)

	return router
}

//StartServer Start HTTTP server for API
func StartServer(storage domain.StorageInterface, config *ServerConfig) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	h := &handler.Handler{
		Storage: storage,
		Logger:  logger,
	}
	router := initRouter(h)

	//Apply middleware
	router.Use(middleware.RecoverMiddleware(h.Logger))
	router.Use(middleware.LoggerMiddleware(h.Logger))

	server := &http.Server{
		Addr:           net.JoinHostPort(config.Host, config.Port),
		Handler:        router,
		ReadTimeout:    config.ReadTimeout,
		WriteTimeout:   config.WriteTimeout,
		MaxHeaderBytes: config.MaxHeaderBytes,
	}

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		logger.Fatalf("HTTP server error: %s", err)
	}
}
