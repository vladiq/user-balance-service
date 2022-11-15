# Микросервис по работе с балансом пользователей

## Setup
1. Запустить контейнер с БД и сервисом.
```bash
docker-compose up
```
2. Файлы с запросами на создание и удаление таблиц лежат в /migrations.


## Запросы
Примеры запросов лежат в [request-examples.http](request-examples.http).

Также есть swagger: http://0.0.0.0:7000/balance-service/swagger/index.html

## Возникшие вопросы
1. Суммы каких порядков будет обрабатывать сервис? Я выбрал DECIMAL(15, 2).
2. Про метод признания выручки написано: "Принимает id пользователя, ИД услуги, ИД заказа, сумму". Я сделал признание выручки и
отмену резервирования по ID резервирования, т.к. несколько записей могут иметь одинаковые параметры из требования выше.
3. Как лучше поступать в случае окончания резервирования (отмена, признание выручки)? Можно либо удалять запись из таблицы, либо хранить ключ is_active
и менять его значение. Я реализовал удаление.
