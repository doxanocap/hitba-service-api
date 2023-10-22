# hitba-service-api

1. Clone the repository:

   ```shell
   git clone https://github.com/doxanocap/hitba-service-api.git
   ```

### Run migrations:

```shell
migrate -path migrations -database "postgresql://postgres:eldoseldos@localhost:5432/service_api?sslmode=disable" up
```

### Run services 


Postgres:
```shell
docker run --name service-api-pg -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=eldoseldos -e POSTGRES_DB=service_api -d postgres:14.5
```

Redis:
```shell
docker run --name service-api-redis -d redis
```