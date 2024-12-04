package main

import (
	"github.com/cantylv/online-music-lib/config"
	"github.com/cantylv/online-music-lib/internal/app"
	"go.uber.org/zap"
)

// main является точкой входа в программу
func main() {
	logger := zap.Must(zap.NewProduction())
	config.Read(logger)
	app.Run(logger)
}
