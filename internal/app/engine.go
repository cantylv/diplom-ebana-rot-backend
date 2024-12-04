package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/cantylv/online-music-lib/internal/delivery/route"
	"github.com/cantylv/online-music-lib/services/postgresql"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// Run инициализирует все проектные зависимости и определяет обработчики
func Run(logger *zap.Logger) {
	// инициализация клиента постгреса
	psqlConn := postgresql.Init(logger)
	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// закрытие соединения с postgresql
		err := psqlConn.Close(ctx)
		if err != nil {
			logger.Error(errors.Wrapf(err, "failed to close postgresql connection").Error())
		} else {
			logger.Info("postgresql connection closed successfully")
		}
	}()

	// define handlers
	r := mux.NewRouter()
	// run server
	handler := route.InitHTTPHandlers(r, psqlConn, logger)
	srv := &http.Server{
		Handler:      handler,
		Addr:         viper.GetString("SERVER_ADDRESS"),
		WriteTimeout: viper.GetDuration("SERVER_WRITE_TIMEOUT"),
		ReadTimeout:  viper.GetDuration("SERVER_READ_TIMEOUT"),
		IdleTimeout:  viper.GetDuration("SERVER_IDLE_TIMEOUT"),
	}

	go func() {
		logger.Info(fmt.Sprintf("server has started at the address %s", srv.Addr))
		if err := srv.ListenAndServe(); err != nil {
			logger.Warn(errors.Wrapf(err, "error after end of receiving requests").Error())
		}
	}()

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), viper.GetDuration("SERVER_SHUTDOWN_DURATION"))
	defer cancel()
	err := srv.Shutdown(ctx)
	if err != nil {
		logger.Error(errors.Wrapf(err, "server has shut down with an error").Error())
	} else {
		logger.Info("server has shut down succesful")
	}
}
