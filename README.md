## development
this repo includes a..
- Go API and sqlite3 database in [/api](api)
- Svelte + TypeScript website in [/web](web)

the database uses [migrate cli](https://github.com/golang-migrate/migrate) to manage migrations, and [sqlc](https://github.com/sqlc-dev/sqlc) to generate code from SQL

### setup
to start the api..
> [!IMPORTANT]
> make sure the migrate cli is installed, and you have an up-to-date database
> ```sh
> go install -tags 'sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate
> migrate -source file://db/migrations -database sqlite3://db/jump.db up
> ```

```sh
cd api
go run .
```

to start the website..

```sh
cd web
npm i
npm run dev
```

### todo
- document sqlc generation
- document migrations
