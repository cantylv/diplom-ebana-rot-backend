package postgresql

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// Init инициализирует клиента PostgreSQL, пингует сервер баз данных, в случае успеха возвращает канал общения с сервером.
func Init(logger *zap.Logger) *pgx.Conn {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		viper.GetString("POSTGRES_USER"),
		viper.GetString("POSTGRES_PASSWORD"),
		viper.GetString("POSTGRES_CONNECTION_HOST"),
		viper.GetInt("POSTGRES_PORT"),
		viper.GetString("POSTGRES_DATABASE"),
		viper.GetString("POSTGRES_SSLMODE_ENABLED"),
	)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	conn, err := pgx.Connect(ctx, connString)
	if err != nil {
		logger.Fatal(errors.Wrapf(err, "error while connecting to postgresql").Error())
	}

	// проверим, что БД доступна
	maxPingAttempts := 3
	successConn := false
	for i := 0; i < maxPingAttempts; i++ {
		ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
		err := conn.Ping(ctx)
		cancel()
		if err == nil {
			successConn = true
			break
		}
		logger.Warn(errors.Wrapf(err, "error while ping#%d to postgresql", i+1).Error())
	}
	if !successConn {
		logger.Fatal("can't establish connection to postgresql")
	}

	logger.Info("postgresql connected successfully")
	return conn
}
