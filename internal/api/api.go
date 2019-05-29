/*
Copyright 2019 Vladislav Dmitriyev.
*/

package api

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/eeonevision/avito-pro-test/internal/api/endpoint"
	"github.com/eeonevision/avito-pro-test/internal/pkg/idempotent"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func buildMuxRouter() *mux.Router {
	router := mux.NewRouter()
	router.NotFoundHandler = http.HandlerFunc(endpoint.NotFoundHandler)

	api := router.PathPrefix("/api").Subrouter()

	v1 := api.PathPrefix("/v1").Subrouter()

	v1.HandleFunc("/generate", endpoint.PostGenerateHandler).Methods(http.MethodPost)
	v1.HandleFunc("/retrieve/{requestID}", endpoint.GetRetrieveHandler).Methods(http.MethodGet)

	return router
}

// Serve mthod starts new http server with setted params in HTTPServer struct.
// The graceful shutdown code was brought from https://github.com/gorilla/mux#graceful-shutdown
func Serve(address string, logger *logrus.Logger) {
	logger.WithFields(logrus.Fields{
		"address": address,
	}).Info("starting server ...")

	srv := &http.Server{
		Addr:         address,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      buildMuxRouter(),
	}

	// Create db for idempotent behaviour of API
	endpoint.IdempotentDB = idempotent.NewDB()

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			logger.WithFields(logrus.Fields{
				"address": address,
				"error":   err.Error(),
			}).Panic("error in starting server ...")
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), srv.ReadTimeout)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	logger.WithFields(logrus.Fields{
		"address": address,
		"timeout": srv.ReadTimeout,
	}).Info("shutting down ...")
	os.Exit(0)
}
