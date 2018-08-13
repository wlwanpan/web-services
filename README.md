# Web Service

## Mount and start
```bash
docker-compose up -d
```

To import from data.csv upon start, set:
```bash
MIGRATE_ON_START=true
```
in the docker-compose file.

Note: the data.csv need to be located at the root of the repo.
## API

- Simple Queries
```bash
curl localhost:8080/unique-users
```

- With filters
```bash
curl localhost:8080/unique-users?device=1,os=2
```

- Example response
```js
{
  count: 123
}
```