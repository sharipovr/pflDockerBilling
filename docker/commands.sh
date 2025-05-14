# Запуск вручную БД postgres:
docker compose -f docker-compose.db.yml up -d

# заход в контейнер с БД для проверки внутренностей
docker exec -it billing_postgres  psql -U postgres

# заход в контейнер с БД для проверки внутренностей (обновленная версия)
docker exec -it billing_postgres psql -U mickey pfl_docker_billing

# Запуск app:
docker compose -f docker-compose.app.yml up -d --build

# Запуск локального контейнера** (для разработки, фронтенд будет обращаться к `localhost:8080`):
docker compose -f docker-compose.app.yml --env-file .env.docker.local up -d --build

# Запуск контейнера на сервере Linode (фронтенд будет обращаться к удалённому IP):
docker compose -f docker-compose.app.yml --env-file .env.docker.prod up -d --build

# Полный перезапуск контейнера с БД на сервере
docker stop billing_postgres
docker rm billing_postgres
cd app/
docker compose -f docker-compose.db.yml up -d

# Решение проблем. Например контейнер с бэкэндом упал сразу старта. Вызовем его лог
docker logs e87337e8c7d1
# или
docker logs billing_backend


# Инспектирование специально созданной докер сети
docker network inspect billing_net