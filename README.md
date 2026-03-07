# go-service-template
Шаблон сервиса на Go, реализующий принципы Чистой Архитектуры. В основе лежит разделение на слои: model, usecase, repository и delivery. Это мое личное видение, проект не претендует на абсолютную истину.

### Установка
git clone github.com/solumD/go-service-template
cd go-service-template/
go mod tidy

### Запуск
Отредактировать файл .env, если нужно. Для запуска должен быть установлен docker compose. В терминале выполнить команду:
```
make build-and-run
```
Будут подняты контейнеры с БД Postrgres и приложением, а также автоматически накатятся миграции. После этого запустится сервер.

### Структура проекта
```
├── cmd
│   └── app
├── config
├── internal
│   ├── app
│   ├── model
│   ├── delivery
│   │   ├── http
│   │   │   ├── v1
│   │   │   │   └── dto
│   └── usecase
│   ├── repository
│   │   └── postgres
├── migrations
├── pkg
│   ├── http_server
│   ├── logger
│   └── postgres
├── .env
├── docker-compose.yaml
└── Makefile
```
