# Сервис TASKS

Команда сборки образа с приложением:
```
docker build -t task-app:v{x} .
```

Команда для запуска контейнера PostgreSQL:
```
docker run --name test-pg -d -e POSTGRES_USER=app -e POSTGRES_PASSWORD=app -e POSTGRES_DB=app -p 5432:5432 postgres:16-alpine
```

