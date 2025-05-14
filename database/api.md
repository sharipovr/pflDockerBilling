
## 1. Создание нового пользователя (Dave):

```bash
curl -X POST http://172.234.232.238:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{
        "name": "Frank",
        "email": "frank@example.com",
        "password": "password123"
      }'
```

После успешного выполнения запроса пользователь Dave будет добавлен в базу данных.

---

## 2. Создание подписки через Stripe API:

```bash
curl -X POST http://172.234.232.238:8080/api/subscriptions/stripe \
  -H "Content-Type: application/json" \
  -d '{
        "email": "frank@example.com",
        "price_id": "price_1RKVADG6nxrf35RxcquoBSD2"
      }'
```