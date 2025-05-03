# Этап 1: Сборка проекта
FROM node:20-alpine AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем package.json и package-lock.json
COPY package.json package-lock.json ./

# Устанавливаем зависимости
RUN npm install

# Копируем все исходники проекта
COPY . .

# Выполняем сборку TypeScript + Vite
RUN npm run build

# Этап 2: Production окружение с Nginx
FROM nginx:stable-alpine

# Копируем готовый билд в папку статических файлов Nginx
COPY --from=builder /app/dist /usr/share/nginx/html

# 👇 добавляем конфиг Nginx
COPY ./nginx/default.conf /etc/nginx/conf.d/default.conf

# Открываем 80 порт для доступа извне
EXPOSE 80

# Стартуем Nginx
CMD ["nginx", "-g", "daemon off;"]
