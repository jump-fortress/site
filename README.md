## development
this repo includes a..
- Go API and sqlite3 database in [/api](api)
- Svelte + TypeScript website in [/web](web)

the database uses [migrate](https://github.com/golang-migrate/migrate) to manage migrations, and [sqlc](https://github.com/sqlc-dev/sqlc) to generate code from SQL

### setup
to start the api..
```sh
cd api
migrate -source file://db/migrations -database sqlite3://db/jump.db up
go run .
```

to start the website..
```sh
cd web
npm i
npm run dev
```
