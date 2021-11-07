# Tests

## Backend

The project's backend tests are all integration tests. They spin up a real server, respond to real httptest requests, and access a real instance of a PostgreSQL database, without mocking anything. This means your machine needs a PostgreSQL database named `no-manga-test` for them to access.

```(commandLine)
(assumes you have PostgreSQL, psql CLI tool installed)
$ dropdb no-manga-test
$ createdb no-manga-test
$ psql -d no-manga-test
$ \i path-to-database/setup.sql
$ \q
(to quit psql)
```

To run the tests, run

```(commandLine)
cd server
go test ./...
```

## Frontend

No frontend written
