# Микросервис по работе с балансом пользователей


## Setup
1. Запуск.
```bash
docker-compose up
```

2. SQL-файлы лежат в ./migrations, DSN: postgres://user:password@localhost:5432/postgres

## Используемые библиотеки

- [zerolog](https://github.com/rs/zerolog) - Логирование. Также реализовал middleware, логирующий входящие запросы.
- [sqlx](https://github.com/jmoiron/sqlx) - Взаимодействие с БД.
- [google/uuid](https://github.com/google/uuid) - Все ID в БД хранятся в UUID.
- [ozzo-validation](https://github.com/go-ozzo/ozzo-validation) - Валидация реквестов (корректность UUID, неотрицательная сумма денег, значение месяца между 1 и 12, ...).
- [go-chi](https://github.com/go-chi/chi/v5) - Routing и middleware.
- [go-chi/render](https://github.com/go-chi/render) - Рендеринг респонсов.
- [http-swagger](https://github.com/swaggo/http-swagger) - Документация API.

## Запросы

Примеры запросов для goland'а лежат в [request-examples.http](request-examples.http).

Также есть swagger: http://0.0.0.0:7000/balance-service/swagger/index.html

![swagger](./assets/swagger.png)

### Создание аккаунта пользователя

```text
### Создание аккаунта с заданным балансом

POST http://localhost:7000/balance-service/accounts
Content-Type: application/json

{
  "amount": 234.23
}

--------
HTTP/1.1 201 Created

<Response body is empty>
```

```text
### Запрос на создание аккаунта с отрицательным балансом

POST http://localhost:7000/balance-service/accounts
Content-Type: application/json

{
  "amount": -1
}

--------
HTTP/1.1 400 Bad Request

validating amount: must be no less than 0
```

### Получение баланса пользователя

```text
### Получение информации об аккаунте по ID
GET http://localhost:7000/balance-service/accounts/e104b74f-b804-43ad-84df-ad8ba376dcdf

--------
HTTP/1.1 200 OK

{
  "id": "e028eda5-ae49-4e8c-a52f-fa5363acf071",
  "balance": 1000
}
```

```text
### Получение информации по невалидному UUID
GET http://localhost:7000/balance-service/accounts/12345

--------
HTTP/1.1 400 Bad Request

binding body: invalid UUID length: 5
```

### Перевести деньги на баланс пользователя

```text
### Перевод денег на баланс пользователя

PUT http://localhost:7000/balance-service/accounts/deposit/303371c3-3073-488c-971f-f57a6f5aeb6f
Content-Type: application/json

{
  "amount": 453.50
}

--------
HTTP/1.1 204 No Content

<Response body is empty>
```

### Списать деньги с баланса пользователя

```text
PUT http://localhost:7000/balance-service/accounts/withdraw/696a8d48-5a40-4ec2-803d-1f9bd6fb2949
Content-type: application/json

{
"amount": 1.10
}

--------
HTTP/1.1 204 No Content

<Response body is empty>
```

### Перевод денег между пользователями

```text
### Перевод между пользователями
POST http://localhost:7000/balance-service/transfers
Content-Type: application/json

{
  "amount": 100,
  "from_id": "10473c75-7cb7-405d-bc23-886fd8564cb6",
  "to_id": "01a223c8-462d-4466-94ab-4f98a7ed4c5f"
}

--------

HTTP/1.1 201 Created

<Response body is empty>
```

```text
### Попытка списать слишком много денег

PUT http://localhost:7000/balance-service/accounts/withdraw/b567e9d6-025c-4d8f-8615-88f391c9e9be
Content-type: application/json

{
  "amount": 1000000000000
}

--------
HTTP/1.1 400 Bad Request

not enough funds to withdraw: bad request
```

### Резервирование денег

```text
### Зарезервировать деньги за заказ

POST http://localhost:7000/balance-service/reservations
Content-Type: application/json

{
  "user_id": "6c8e7526-95c6-4822-97f8-168c4d2d2fb9",
  "service_id": "e028eda5-ae49-4e8c-a52f-fa5363acf071",
  "order_id": "521ab429-832d-4bc7-b870-26f491ad0b27",
  "amount": 100
}

--------
HTTP/1.1 204 No Content

<Response body is empty>
```

### Отмена резервирования

```text
### Отмена резервирования, возвращение денег пользователю

DELETE http://localhost:7000/balance-service/reservations/cancel-reservation/a51a02f4-a827-4c11-bbbe-43650296b588

--------
HTTP/1.1 204 No Content

<Response body is empty>
```

### Подтверждение резервирования

```text
### Подтверждение резервации, внесение записи в отчёт для бухгалтерии

DELETE http://localhost:7000/balance-service/reservations/confirm-reservation/17da187d-2ea9-4376-94c4-4af09011eaea

--------

HTTP/1.1 204 No Content

<Response body is empty>
```

### Список транзакций пользователя с пагинацией и сортировкой по сумме и дате

ID следующей страницы возвращается в JSON'е, ключ next_page_id.

```text
### Список транзакций пользователя, страница 1. Сортировка по сумме.

GET http://localhost:7000/balance-service/transfers/reports/4509b5d2-bca1-43be-a5e6-44fc9293de2c?order-by=amount
Accept: application/json

--------
HTTP/1.1 200 OK

{
  "items": [
    {
      "ID": "059651fc-1410-43e1-8471-1a3f245cb482",
      "AccountID": "4509b5d2-bca1-43be-a5e6-44fc9293de2c",
      "IsAccrual": false,
      "Amount": 100,
      "Info": "merchant withdrawal",
      "CreatedAt": "2022-11-15T14:09:14.599789Z"
    },
    {
      "ID": "1a3d1d56-1490-473f-9e50-80b33b08b089",
      "AccountID": "4509b5d2-bca1-43be-a5e6-44fc9293de2c",
      "IsAccrual": false,
      "Amount": 120,
      "Info": "merchant withdrawal",
      "CreatedAt": "2022-11-15T14:09:36.344897Z"
    },
    {
      "ID": "28b60740-f0f3-40f2-9147-71bc08ba7472",
      "AccountID": "4509b5d2-bca1-43be-a5e6-44fc9293de2c",
      "IsAccrual": false,
      "Amount": 256.30,
      "Info": "merchant withdrawal",
      "CreatedAt": "2022-11-15T14:09:32.11071Z"
    },
    {
      "ID": "3669fe18-1755-429f-b597-ebf9e9007915",
      "AccountID": "4509b5d2-bca1-43be-a5e6-44fc9293de2c",
      "IsAccrual": false,
      "Amount": 100000,
      "Info": "merchant withdrawal",
      "CreatedAt": "2022-11-15T14:09:27.399031Z"
    }
  ],
  "next_page_id": "37c07fc9-7747-487c-aac6-f44cd9536b9c"
}
```


```text
### Страница 2, сортировка по времени

GET http://localhost:7000/balance-service/transfers/reports/4509b5d2-bca1-43be-a5e6-44fc9293de2c?order-by=date&page-id=37c07fc9-7747-487c-aac6-f44cd9536b9c
Accept: application/json

--------
HTTP/1.1 200 OK

{
  "items": [
    {
      "ID": "37c07fc9-7747-487c-aac6-f44cd9536b9c",
      "AccountID": "4509b5d2-bca1-43be-a5e6-44fc9293de2c",
      "IsAccrual": true,
      "Amount": 100,
      "Info": "merchant deposit",
      "CreatedAt": "2022-11-15T14:08:54.407696Z"
    },
    {
      "ID": "4b0bd916-bc00-45e2-9eef-c9f2e669ddac",
      "AccountID": "4509b5d2-bca1-43be-a5e6-44fc9293de2c",
      "IsAccrual": false,
      "Amount": 100,
      "Info": "merchant withdrawal",
      "CreatedAt": "2022-11-15T14:09:35.750302Z"
    },
    {
      "ID": "50f8fb9c-af76-49f8-ad19-3de7f9d5f5a3",
      "AccountID": "4509b5d2-bca1-43be-a5e6-44fc9293de2c",
      "IsAccrual": false,
      "Amount": 100,
      "Info": "merchant withdrawal",
      "CreatedAt": "2022-11-15T14:09:38.767387Z"
    },
    {
      "ID": "69a9ff7f-0abe-4ea3-bbc7-3536c04f73cd",
      "AccountID": "4509b5d2-bca1-43be-a5e6-44fc9293de2c",
      "IsAccrual": false,
      "Amount": 100,
      "Info": "merchant withdrawal",
      "CreatedAt": "2022-11-15T14:09:39.313725Z"
    }
  ],
  "next_page_id": "7a364945-6dc7-423f-939d-56ae230e8272"
}
```

### Месячный отчёт для бухгалтерии

```text
### Отчёт для бухгалтерии за определённый месяц

GET http://localhost:7000/balance-service/reports?month=11
Accept: text/csv

--------
HTTP/1.1 200 OK
Content-Disposition: attachment; filename="user-report.csv"
Content-Type: text/csv

ServiceID                           ,Amount
c7fe41c5-200a-4d69-89fc-599c9f5e7cdb,861
e028eda5-ae49-4e8c-a52f-fa5363acf071,200
```

```text
### Отчёт для бухгалтерии за невалидный месяц

GET http://localhost:7000/balance-service/reports?month=-2
Accept: text/csv

--------
HTTP/1.1 400 Bad Request

validating month: must be no less than 1
```
