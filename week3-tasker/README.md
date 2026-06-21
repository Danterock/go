# Task Manager

Консольный менеджер задач на Go.

## Возможности

* Просмотр всех задач
* Добавление новой задачи
* Просмотр активных задач
* Поиск задачи по названию
* Отметка задачи как выполненной
* Удаление задачи по ID
* Просмотр статистики
* Сохранение задач в JSON

## Структура проекта

```text
cmd/tasker/main.go
task/model.go
task/print.go
task/service.go
task/validate.go
task/task_test.go
storage/jsonTasks.go
storage/json_tets.go
stats/stats.go
data/tasks.json
data/app.log
data/report.txt
```

## Запуск

```bash
go run ./cmd/tasker
```

## Запуск тестов

```bash
go test ./...
```

## Формат хранения данных

Файл:

```text
data/tasks.json
```

Пример:

```json
[
    {
        "id": 1,
        "title": "Learn Go",
        "priority": "high",
        "done": false
    }
]
```

## Меню

```text
1. Show all tasks
2. Add task
3. Show active tasks
4. Mark task as done
5. Delete task by ID
6. Search task by title
7. Show statistics
8. Save tasks
0. Exit
```

## Тесты

Реализованы тесты для:

* загрузки отсутствующего файла;
* загрузки пустого файла;
* обработки повреждённого JSON;
* сохранения и загрузки данных;
* проверки валидации названия;
* поиска задачи по ID.

## Используемые технологии

* Go
* JSON
* os
* encoding/json
* testing
