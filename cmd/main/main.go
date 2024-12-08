package main

import (
	"github.com/cantylv/online-music-lib/config"
	"github.com/cantylv/online-music-lib/internal/app"
	"go.uber.org/zap"
)

//	@title			Swagger API для сервиса получения текста песен
//	@version		1.0
//	@description	Сервис работает с текстами песнями, спектр действий - CRUD. Доступны пагинация и фильтрация.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	Лобанов И.И.
//	@contact.url	http://t.me/cantylv
//	@contact.email	physic2003@mail.ru

//	@host		localhost:8080
//	@BasePath	/api/v1
func main() {
	logger := zap.Must(zap.NewProduction())
	config.Read(logger)
	app.Run(logger)
}
