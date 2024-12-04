package main

import (
	"time"

	"github.com/cantylv/online-music-lib/config"
	"go.uber.org/zap"
)

func main() {
	logger := zap.Must(zap.NewProduction())
	config.Read(logger)
	time.Sleep(5 * time.Second)
}
