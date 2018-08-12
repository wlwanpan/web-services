# Web Service

## Mount
```bash
docker-compose up -d
```

## Migrate Data
```bash
docker exec -it web-service /bin/bash
cd migration && go run migration.go
```

## API

- Simple Queries
```bash
curl localhost:8080/unique-users
```

- Filters
```bash
curl localhost:8080/unique-users?device=1,os=2
```

- Example response
```js
{
  count: 123
}
```