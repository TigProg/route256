# Система бронирования места в автобусе

Пример брони:  
`1: ABCD45 / 2022-01-01 / 7`  
внутренний идентификатор: `1` - инкремент  
маршрут: `ABCD45` - 4-10 символов  
дата: `2022-01-01` - строковое представление в формате `YYYY-MM-DD`  
номер места: `7` - от 1 до 100 включительно

### Запуск

поднятие окружения: `docker-compose up -d`  
запуск всех миграций БД: `./migrate.sh`  
сборка: `make build`  
запуск: `make run`

По-умолчанию запускаются telegram бот, gRPC сервер, HTTP-сервер 
в качестве интерфесов пользователя.
Также запускается бэкенд gRPC сервер для взаимодействия с БД, Kafka кластер и сервер метрик.

Запись данных происходит асинхронно по Kafka (3 ноды + 1 нода Zookeeper), чтение синхронное.

Возможности взаимодействия:  
1. [telegram_bot](https://t.me/tigprog_bot) -
необходимо задать переменную окружения TELEGRAM_API_KEY с токеном telegram бота  
2. gRPC сервер на `localhost:8081` (присутствует пример клиента `client/client.go`)
3. HTTP-сервер на `localhost:8080`

### Отладка

После запуска доступны:

1. [pprof](http://localhost:9876/debug/pprof/)
2. [метрики](http://localhost:9876/debug/vars), в том числе кастомные:
```
custom_consumer::income  // всего чтений Kafka-consumer
custom_consumer::success  //  успешных операций
custom_consumer::fail  // неуспешных операций
```

Метрики хранятся in-memory, после перезапуска обновляются.

### Тестирование

#### unit

1. запуск тестов: `make test`
2. coverage: `make cover`

#### integration

1. подготовка окружения:
```
docker-compose up -d
./migrate.sh
make run
```
2. запуск тестов:
```
make integration
```

### Функциональность

#### Telegram бот

1. при отправке сообщения или пустой команды (`/`) работает как эхобот
2. при отправке некорректной команды сработает предложение о вводе `/help`
3. `/help` - список доступных команд
4. `/list <offset> <limit>` - вывести все забронированые места
5. `/get <id>` - вывести бронь по идентификатору
6. `/add <route> <date> <seat>` - создань новую бронь
7. `/change_seat <id> <seat>` -
   обновить номер места на значение `seat` по идентификатору
8. `/change_date_seat <id> <date> <seat>` -
   обновить дату и номер места на значения `date` и `new_value` по идентификатору
9. `/delete <id>` - удалить бронь по идентификатору

#### gRPC

Аналогично telegram боту (за исключением echo и help).  
Proto файл находится в `api/api.proto`.

#### HTTP

Swagger находится в `gen/openapiv2/api.swagger.json`.  
Генерируется на основе `.proto` файла.
