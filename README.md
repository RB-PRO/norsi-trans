#### Тестовое задание №2

### Задание:

*Небольшой микросервис с http API. Три метода на добавление, получение удаление данных. Данные хранить в любой БД (н-р postgres)*

##### Запуск:

```bash
make
```

##### Запустить тесты:

```bash
go test -v ./...
```

### Структура проекта:

[cmd/main](cmd/main/) - Запуск программы

[internal ](internal/)- Логика запуска и настройки API методов проекта

[internal/api](internal/api/) - Модуль API с использованием [Gin ](https://github.com/gin-gonic/gin)с методами [пингования](internal/api/PING.go), [добавления](internal/api/POST.go), [получения](internal/api/GET.go), [удаления](internal/api/DELETE.go) данных.

[internal/bd](internal/bd/) - Модуль базы данных с методами [добавления](internal/bd/insert.go), [получения](internal/bd/select.go), [удаления ](internal/bd/delete.go)данных.
