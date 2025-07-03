# seg_service

## Описание

**seg_service** — сервис для управления сегментами пользователей.

- Создание, удаление, переименование сегментов
- Добавление/удаление пользователей в сегменты
- Случайное распределение сегмента на процент пользователей
- Получение сегментов пользователя через API

Архитектура: Clean Architecture (чистая архитектура).

---

## Структура проекта

```
├── .env                # Переменные окружения для docker-compose и приложения
├── docker-compose.yml  # Запуск всей инфраструктуры (Postgres, Redis, сервис)
├── README.md           # Описание проекта
└── seg_service/
    ├── Dockerfile      # Dockerfile для сборки сервиса
    ├── go.mod, go.sum  # Go-модули и зависимости
    ├── main.go         # Точка входа
    ├── config/         # Конфиг и пример .env
    ├── internal/
    │   ├── domain/     # Бизнес-модели и интерфейсы
    │   ├── repository/ # Реализация репозиториев (Postgres, Redis)
    │   ├── service/    # Бизнес-логика
    │   └── handler/    # HTTP-обработчики
    └── migrations/     # Миграции для Postgres
```

---

## Быстрый старт (Docker Compose)

1. **Создай .env в корне** (или скопируй из `.env-example`):

```
POSTGRES_DB=seg_service
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DSN=postgres://postgres:postgres@postgres:5432/seg_service?sslmode=disable
REDIS_ADDR=redis:6379
REDIS_PASS=
HTTP_PORT=:8080
```

2. **Запусти сервис и инфраструктуру:**

```sh
docker-compose up --build
```

- Сервис будет доступен на http://localhost:8080
- Postgres — на порту 5433 (логин/пароль из .env)
- Redis — на порту 6379

3. **Миграции применяются автоматически контейнером migrate**

---

## Примеры API-запросов

### Linux/Mac/Docker (curl)

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

### Windows PowerShell (Invoke-WebRequest)

> В PowerShell для POST-запросов с JSON используйте Invoke-WebRequest!

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

- Все переменные окружения должны быть заданы (лучше всего через .env в корне).
- Перед первым запуском обязательно примените миграции (docker-compose делает это автоматически).
- Redis кеш реализован как заглушка — можно доработать при необходимости.
- Логика и структура легко расширяются под production.

---

## Контакты

Если возникнут вопросы по архитектуре, запуску или расширению — пиши https://t.me/nekell3r! 
