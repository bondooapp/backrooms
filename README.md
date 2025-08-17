## README

### How to use?

```shell
go get github.com/bondooapp/backrooms@latest
```

### How to test locally?

#### Start test

* Replace the backrooms path with your local path

```shell
go mod edit -replace=github.com/bondooapp/backrooms=../backrooms

go mod tidy
```

* Test your code
* Restore reference

```shell
go mod edit -dropreplace=github.com/bondooapp/backrooms

go mod tidy
```

### Postgres Module

#### Params

```text
POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_USER=root
POSTGRES_PASSWORD=bondoo
POSTGRES_DB_NAME=bd-project-service
POSTGRES_SSL_MODE=disable
POSTGRES_MAX_OPEN_CONNS=20
POSTGRES_MAX_IDLE_CONNS=5
POSTGRES_CONN_MAX_LIFETIME=60
```

### Redis Module

#### Params

```text
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=
REDIS_DB=0
REDIS_POOL_SIZE=100
```
