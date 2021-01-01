# Database

## Set up

The project needs to be able to access a PostgreSQL database named `no-manga`

```(commandLine)
(assumes you have PostgreSQL, psql CLI tool installed)
$ dropdb no-manga
$ createdb no-manga
$ psql -d no-manga
$ \i path-to-database/setup.sql
$ \q
(to quit psql)
```

There is a table diagram in `tableDiagram.png`, which is also shown at <https://dbdiagram.io/d/5fe996aa9a6c525a03bc7393> (although the remote diagram needs to be kept in sync manually)

## Design

To be written and git fixed-up into this file
