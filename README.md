## Requirement

- [Go](https://golang.org/doc/install)
- [MySQL](https://www.mysql.com/downloads/)

## Requirement

Create `.env` file with detail below

```
DB_USERNAME=(your database username)
DB_PASSWORD=(your database password)
DB_DATABASE=(your database name)
```

## How to Run

1. Make sure `Go` and `MySQL` had been installed and running
2. Type `go mod download` to install packages that needed
3. Type `go build` to compile executable file
4. Run `./foreign-currency-go` (in MacoS/Linux) or `foreign-currency-go.exe` (in Windows)
