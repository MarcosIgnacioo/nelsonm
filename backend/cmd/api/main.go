package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/julienschmidt/httprouter"
)

const version = "1.0.0"

type Config struct {
	port       int
	enviroment string
}

type Application struct {
	config Config
	logger *slog.Logger
}

func (app *Application) routes() *httprouter.Router {

	router := httprouter.New()
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheck_handler)
	router.HandlerFunc(http.MethodPost, "/v1/image", app.image)
	router.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // Permite cualquier origen (*)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.WriteHeader(http.StatusNoContent)
	})
	return router
}

func main() {
	var config Config
	var app *Application

	flag.IntVar(&config.port, "port", 4076, "Super cool API server port") // would be cool and sexy to implement this in c
	flag.StringVar(&config.enviroment, "env", "dev", "Enviroment is pretty")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app = &Application{
		config: config,
		logger: logger,
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.config.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	logger.Info("starting a new server girl", "address", server.Addr, "env", app.config.enviroment)

	err := server.ListenAndServe()

	logger.Error(err.Error())
	os.Exit(1)

}
