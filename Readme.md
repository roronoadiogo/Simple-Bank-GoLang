# :books: Simple bank example using GoLang

This project it's to understand the mainly tools used for language GoLang to API services, ORM, Migration and another stuffs to fixed the knowledge adquired.

## :wrench: Setup and Tools used

- GoLang 1.20
- SQLC to generate models (in discussion to check other GO ORM's)
- MigrationDB
- GoDotEnv
- Docker
    - Postgres Alpine 12

This list does not complete yet, will increase with the time.

## :dart: Goals

- Understading the tools most commom used with GoLang

- Implementation using Docker to configuration in the enviroment

- Produce a front to consume the API (In discussion if Angular or VueJs)

- Try adopt the best pratices for GoLang Projects

- Understand the tests with GoLang

## :rocket: Using the project

- Open a terminal in the directory root this project and run:
```console
user@machine:~$ make postgres

```
- Create the database in the image of postgres
```console
user@machine:~$ make createdb

```
- Run the migrateDB to use the version of database until the moment
```console
user@machine:~$ make migratedb

```
This documentation it's not complete, I'm still working on this resource to automatization better

## :floppy_disk: Database initial version
<img src="docs/diagrams/database/database_version_1.png" alt="Initial version" width="700" height="300" />

This is not definitive yet, in discussion the of domain of bussiness.