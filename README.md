# Payment Test

## Description
 
Solution for getting payment buttons for services:
 * Google Pay
 * Stripe
 * Apple Pay
 * Paypal
 
##  API

POST: http://localhost:3001/v1/api/payment-buttons

Request:
```
{
    "product_id": "1"
}
``` 
 
## Simple build and run
  
```cd cmd/app```
```go build```
```./app```

Default port: 8080. For custom port:

```APP_API_PORT=3001 ./app```


## Docker build and run 

```docker build -t test .```
```docker run test```

## Docker-compose run


## Замечания

1. Обычно использую gorilla/mux для решений rest api
2. Добавил модуль Storage для получения id на конкретных сервисах (если они отличаются от базового id)
3. Возможно не совсем понял задание, потому как полез в API оплат конкретных сервисов. Есть решения - репозитории с реализализацией API оплат,  но так глубоко не лез.
4. Работаю в Goland (JetBrains IDE), где для тестов использую автогенерацию уже с даленейшим редактированием. Не очень хватило времени сейчас.