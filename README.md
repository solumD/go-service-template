# go-service-template
Шаблон сервиса на Go, реализующий принципы Чистой Архитектуры. В основе лежит разделение на слои: model, usecase, repository и transport. Это мое личное видение, проект не претендует на абсолютную истину.

### Запуск
Для запуска необходим docker compose. В терминале выполнить команду:
```
make build-and-run
```
Будет поднят контейнер с БД Postrgres, а также автоматически накатятся миграции. После этого запустится сервер.

### Структура проекта
```
├── cmd
│   └── app
├── config
├── internal
│   ├── app
│   ├── model
│   ├── transport
│   │   ├── handler
│   │   │   ├── dto
│   └── usecase
│   ├── repository
│   │   └── postgres
├── migrations
├── pkg
│   ├── http_server
│   ├── logger
│   └── postgres
```
