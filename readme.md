## Go Task API

Простой REST API сервис на Go (Golang) для управления задачами (Task), построенный с использованием Gin и структурного логгера slog.

## Возможности
    Добавление задачи
    Получение задачи по ID
    Получение всех задач
    Удаление задачи
    JSON-логирование через log/slog
    Встроенный хранилище на map[int]Task
    Расширяемая архитектура (логгирование, слои)


## Структура проекта

go-gin-task-api/
├── cmd/
│   └── main.go           # Точка входа
├── internal/
│   ├── controllers/      # HTTP-обработчики
│   ├── models/           # Определение сущностей
│   ├── routes/           # Роутинг
│   └── services/         # Бизнес-логика
├── pkg/
│   └── logger/           # Логирование
├── go.mod
├── go.sum
└── README.md

## Используемые технологии
    Go 1.21+
    Gin Web Framework
    log/slog (новый структурированный логгер Go)
    Встроенное хранилище на map[int]Task
    Простая и масштабируемая архитектура

## Установка и запуск

1 Клонировать репозиторий
git clone https://github.com/yourusername/go-gin-task-api.git
cd go-gin-task-api

2 Инициализировать модули
go mod tidy

3 Запустить приложение
go run cmd/main.go

После запуска сервис будет доступен по адресу:
http://localhost:8080

## Примеры API-вызовов (через curl)

Отобразить список задач
curl http://localhost:8080/tasks/

Создать задачу
curl -X POST http://localhost:8080/task/ \
  -H "Content-Type: application/json" \
  -d "{\"title\":\"Сделать задачу\",\"content\":\"Закончить упражнение по Go\"}"

Получить задачу по ID
curl http://localhost:8080/task/1

Удалить задачу
curl -X DELETE http://localhost:8080/task/1
