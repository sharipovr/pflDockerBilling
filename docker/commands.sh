# Запуск вручную БД postgres:
docker compose -f docker-compose.db.yml up -d

# заход в контейнер с БД дял проверки внутренностей
docker exec -it billing_postgres  psql -U postgres

# Запуск app:
docker compose -f docker-compose.app.yml up -d --build