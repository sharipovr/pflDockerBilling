#!/bin/bash
# See /Users/rustemsharipov/workspace/github.com/sharipovr/portfolios/pflDockerBilling/docs/folders.md

echo "Создаю структуру подпапок для pflDockerBilling..."

folders=(
  "backend"
  "frontend"
  "database"
  "infrastructure"
  "docker"
  "docs"
  "scripts"
  "tests"
  ".github"
)

for folder in "${folders[@]}"; do
  mkdir -p "$folder"
  echo "Создана папка: $folder"
done

echo "Все папки успешно созданы."
