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
