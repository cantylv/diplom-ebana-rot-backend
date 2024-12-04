package config

import (
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// Read получает переменные конфигурации из среды выполнения. Если переменная не установлена, то вернется zero value.
func Read(logger *zap.Logger) {
	viper.AutomaticEnv()
	// получим ключи из .env файла
	mapEnv, err := godotenv.Read("build/.env")
	if err != nil {
		logger.Fatal(errors.Wrapf(err, ".env file not found").Error())
	}

	for key := range mapEnv {
		viper.BindEnv(key)
	}

	// запишем текущую конфигурацию запуска в декларативные файлы (создаются при каждом запуске)
	go func() {
		type config struct {
			filepath string
			filetype string
		}
		confs := []config{
			{filepath: "config/temp/config.yaml", filetype: "yaml"},
			{filepath: "config/temp/.env", filetype: "env"},
		}
		success := true
		// записи в файлы конфигурации должны выполняться последовательно, так как целевые файлы конфигурации различаются,
		// а viper - это один объект
		for _, conf := range confs {
			if err := saveCurrentConfig(conf.filepath, conf.filetype); err != nil {
				logger.Warn(err.Error())
				success = false
			}
		}
		if success {
			logger.Info("current configuration file was successful created")
		}
	}()
	logger.Info("configuration reading completed")
}

func saveCurrentConfig(filepath, filetype string) error {
	viper.SetConfigType(filetype)
	if err := viper.WriteConfigAs(filepath); err != nil {
		return errors.Wrapf(err, "configuration file %s was not found", filepath)
	}
	return nil
}
