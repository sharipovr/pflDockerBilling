# Запуск вручную БД postgres:
docker compose -f docker-compose.db.yml up -d

# заход в контейнер с БД дял проверки внутренностей
docker exec -it billing_postgres  psql -U postgres

# Запуск app:
docker compose -f docker-compose.app.yml up -d --build

# Запуск локального контейнера** (для разработки, фронтенд будет обращаться к `localhost:8080`):
docker compose -f docker-compose.app.yml --env-file .env.docker.local up -d --build

# Запуск контейнера на сервере Linode (фронтенд будет обращаться к удалённому IP):
docker compose -f docker-compose.app.yml --env-file .env.docker.prod up -d --build