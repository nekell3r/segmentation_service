# seg_service

## Описание

**seg_service** — это минимальный сервис для управления сегментами пользователей.
Сервис позволяет:
- создавать, удалять, переименовывать сегменты,
- добавлять/удалять пользователей в сегменты,
- случайно распределять сегмент на процент пользователей,
- получать список сегментов пользователя через API.

Архитектура построена по принципам чистой архитектуры (Clean Architecture).

---

## Структура проекта

```
seg_service/
│
├── config/
│   ├── config.go         # Загрузка конфигурации из переменных окружения
│   └── .env.example      # Пример .env файла
│
├── internal/
│   ├── domain/           # Бизнес-модели и интерфейсы (чистая архитектура)
│   ├── repository/       # Реализация репозиториев (PostgreSQL, Redis)
│   ├── service/          # Бизнес-логика (работа с сегментами)
│   └── handler/          # HTTP-обработчики и роутинг
│
├── migrations/           # Миграции для PostgreSQL
│
├── main.go               # Точка входа, сборка приложения
├── go.mod                # Go-модуль и зависимости
└── README.md             # Описание проекта (этот файл)
```

---

## Быстрый старт

### 1. Настрой .env

Создай файл `config/.env` (или скопируй из `.env.example`) и укажи параметры подключения к PostgreSQL, Redis и порт сервера:

```
POSTGRES_DSN=postgres://user:password@localhost:5432/seg_service?sslmode=disable
REDIS_ADDR=localhost:6379
REDIS_PASS=
HTTP_PORT=:8080
```

### 2. Применить миграции

Создай базу данных и примени миграции (например, с помощью goose):

```sh
goose -dir seg_service/migrations postgres "postgres://user:password@localhost:5432/seg_service?sslmode=disable" up
```

### 3. Запусти сервис

```sh
go run main.go
```
или
```sh
go build -o seg_service && ./seg_service
```

---

## Примеры API-запросов

### Для Linux/Mac или Docker (curl)

#### Создать сегмент
```sh
curl -X POST -H "Content-Type: application/json" -d '{"name":"MAIL_GPT"}' http://localhost:8080/segment/create
```

#### Добавить пользователя в сегмент
```sh
curl -X POST -H "Content-Type: application/json" -d '{"UserID":15230,"Segment":"MAIL_GPT"}' http://localhost:8080/segment/add_user
```

#### Получить сегменты пользователя
```sh
curl "http://localhost:8080/user/segments?user_id=15230"
```

#### Удалить сегмент
```sh
curl -X POST "http://localhost:8080/segment/delete?name=MAIL_GPT"
```

#### Переименовать сегмент
```sh
curl -X POST -H "Content-Type: application/json" -d '{"OldName":"MAIL_GPT","NewName":"MAIL_GPT_NEW"}' http://localhost:8080/segment/rename
```

#### Удалить пользователя из сегмента
```sh
curl -X POST -H "Content-Type: application/json" -d '{"UserID":15230,"Segment":"MAIL_GPT"}' http://localhost:8080/segment/remove_user
```

#### Случайно распределить сегмент на процент пользователей
```sh
curl -X POST -H "Content-Type: application/json" -d '{"Segment":"MAIL_GPT","Percent":30}' http://localhost:8080/segment/distribute
```

---

## Примеры для Windows PowerShell (Invoke-WebRequest)

> **Важно:** В PowerShell не используйте curl для POST-запросов с JSON — используйте Invoke-WebRequest!

#### Создать сегмент
```powershell
Invoke-WebRequest -Uri "http://localhost:8080/segment/create" `
  -Method POST `
  -Body '{"name":"MAIL_GPT"}' `
  -ContentType "application/json"
```

#### Добавить пользователя в сегмент
```powershell
Invoke-WebRequest -Uri "http://localhost:8080/segment/add_user" `
  -Method POST `
  -Body '{"UserID":15230,"Segment":"MAIL_GPT"}' `
  -ContentType "application/json"
```

#### Получить сегменты пользователя
```powershell
Invoke-WebRequest -Uri "http://localhost:8080/user/segments?user_id=15230" -Method GET
```

#### Удалить сегмент
```powershell
Invoke-WebRequest -Uri "http://localhost:8080/segment/delete?name=MAIL_GPT" -Method POST
```

#### Переименовать сегмент
```powershell
Invoke-WebRequest -Uri "http://localhost:8080/segment/rename" `
  -Method POST `
  -Body '{"OldName":"MAIL_GPT","NewName":"MAIL_GPT_NEW"}' `
  -ContentType "application/json"
```

#### Удалить пользователя из сегмента
```powershell
Invoke-WebRequest -Uri "http://localhost:8080/segment/remove_user" `
  -Method POST `
  -Body '{"UserID":15230,"Segment":"MAIL_GPT"}' `
  -ContentType "application/json"
```

#### Случайно распределить сегмент на процент пользователей
```powershell
Invoke-WebRequest -Uri "http://localhost:8080/segment/distribute" `
  -Method POST `
  -Body '{"Segment":"MAIL_GPT","Percent":30}' `
  -ContentType "application/json"
```

---

## Важно

- Все переменные окружения должны быть заданы (лучше всего через .env и godotenv).
- Перед первым запуском обязательно примени миграции к базе данных.
- Redis кеш реализован как заглушка — можно доработать при необходимости.
- Логика и структура легко расширяются под production.
- **curl-примеры пригодятся для docker и Linux-среды.**

---

## Контакты

Если возникнут вопросы по архитектуре, запуску или расширению — пиши! 