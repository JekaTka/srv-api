# REST API Service

## How to run on local env?
1. Run PG: `docker run -d -p 5432:5432 -e POSTGRES_PASSWORD=postgres postgres`
2. Run migrations: `migrate -path migrations -database "postgres://postgres:postgres@localhost/postgres?sslmode=disable" up`
