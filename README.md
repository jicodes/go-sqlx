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
run `direnv allow` to load the environment variables

Install the sqlx package

```sh
go get github.com/jmoiron/sqlx
```

Install the pq driver for Go's database/sql package

```sh
go get github.com/lib/pq
```

Install goose for database migration

```sh
go install github.com/pressly/goose/v3/cmd/goose@latest
```

This will install the goose binary to your `$GOPATH/bin` directory.

Create the `db/migrations` folder and set environment variables in `.envrc` file

```envrc
GOOSE_DRIVER=DRIVER
GOOSE_DBSTRING=DBSTRING
GOOSE_MIGRATION_DIR=MIGRATION_DIR
```

Usage of goose

```sh
goose create create_users sql
```

Replace the content in goose up section with:
`create extension if not exists citext;`
This will create the citext(case-insensetive text) extension in the database
We can use the citext data type for email address in the database

Create the users table in the database

```create table users (
  id serial primary key,
  email citext unique not null,
  password_hash varchar(255),
  inserted_at timestamp(0) not null default (now() at time zone 'utc'),
  updated_at timestamp(0) not null default (now() at time zone 'utc')
);
```
In goose up section replace the content with:
`drop table if exists users;`

Now run the goose up command to create the users table in the database

```sh
goose up
```