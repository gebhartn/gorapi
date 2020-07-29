# Gorapi
A small RESTful API written to absorb some Golang

### Getting started

Example `.env` file:

```
APP_DB_USERNAME=postgres
APP_DB_PASSWORD=password
APP_DB_NAME=postgres
APP_LISTEN=":8080"
```

Write the schema to your database:

```
psql -d postgres -U postgres -h localhost -f util/schema.sql
```

To run all tests:

```
make test
```

To build the app:

```
make build
```

To run the app:

```
make run
```
