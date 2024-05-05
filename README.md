# Resti Test Task


## Clone the project

```
$ git clone https://github.com/zsandibe/resti_task.git
$ cd resti_task
```

## Launch a project

```
$ make run
```

## Execute migrations

```
$ make migrate-up
$ make migrate-down
```

OR 
```
$ make docker-up
$ make docker-down
```

## API server provides the following endpoints:
* `GET /api/v1/accounts/create` - creates an account
* `GET /api/v1/accounts/` - returns list of accounts with all transactions
* `POST /api/v1/transactions/create` - creates a transaction
* `GET /api/v1/transactions/:id` - gets transactions by account id

# .env file
## Server configuration

```
SERVER_HOST=localhost
SERVER_PORT=8888
```

## Postgres configuration

```
DRIVER=
DB_USER=
DB_PASSWORD=
DB_HOST=
DB_PORT=
DB_NAME=
```

## To see swagger docs
```bash
    make run
``` 
OR 
```
    make docker-up
```

* Then you can see docs in at http://localhost:8888/swagger/index.html


По тестовому заданию:Посчитал нужным не передавать date в пост запросе создании транзакции.

