# Comp tech market

0. Настроит git, установить postgres и postman
1. Прописать api ручки в одном файле main.go максимально просто, без похода в базу, данные хранить в глобальной переменной
    1. `POST` `/item` req body{"name", "price", ...}, resp 200, id; 400, err; 500, err;
    2. `GET` `/item/:id`, resp 200, body{"name", "price", ...}; 400, err; 500, err;
    3. `PUT` `/item/:id` req body{"name", "price", ...}, resp 200, id; 400, err; 500, err;
    4. `DELETE` `/item/:id`, resp 200, id; 400, err; 500, err;

Используй fiber для принятия запросов

2. слайс vs массив, что происходит при append, структура слайса
3. map, бакеты, миграции, коллизии