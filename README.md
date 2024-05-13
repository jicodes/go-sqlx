# Connect postgresql db with sqlx package in Golang project

Run the postgres db in the Docker

```sh
docker run --name pg-container -p 5432:5432 -e POSTGRES_PASSWORD=xxx -d postgres:16-alpine
```

Create the user and db with Docker exec

```sh
docker exec -it pg-container createdb -U postgres testdb
```

Create `.envrc` file with environment variables 

Install the sqlx package
```sh
go get github.com/jmoiron/sqlx
```

Install the pq driver for Go's database/sql package
```sh
go get github.com/lib/pq
```
