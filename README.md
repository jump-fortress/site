# site

this repo includes a..

- Go API and sqlite3 database in [/api](api)
- Svelte + TypeScript frontend in [/web](web)

development tracking currently takes place within a private Linear workspace

## starting the api

the database uses [migrate cli](https://github.com/golang-migrate/migrate) to manage migrations, and [sqlc](https://github.com/sqlc-dev/sqlc) to generate code from SQL

api documentation is provided at `/docs` (see [.env.local.example](api/env/.env.local.example?plain=1#L17))

> [!IMPORTANT]
> make sure the migrate cli is installed, and you have an up-to-date database
>
> ```sh
> go install -tags 'sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate
> migrate -source file://db/migrations -database sqlite3://db/jump.db up
> ```

```sh
cd api
go run .
```

> [!TIP]
> to generate go code from migrations and sql..
>
> ```sh
> cd api
> sqlc generate
> ```

## starting the website

the frontend uses [openapi-typescript](https://github.com/openapi-ts/openapi-typescript) to generate types from the API's schema.

```sh
cd web
npm i
npm run dev
```

> [!TIP]
> to generate typescript types from the api.. (with the api up!!)
>
> ```sh
> cd web
> npm run openapi-typescript
> ```

### todo

- document migrations
- document file tree
