# Project Name

Test Backend API Tahap 1.

## Requirements

To run this project, you need:

- Go 1.22
- GoLand 2024.1.4
- GoFiber V2.x

## Struktur Folder

Struktur folder dari project ini adalah sebagai berikut:

```plaintext
.
├── cmd
│   └── app
│       └── main.go
├── docs
│   ├── docs.go
│   └── swagger.json
├── internal
│   ├── domain
│   │   ├── order.go
│   │   ├── order_item.go
│   │   ├── printer.go
│   │   ├── product.go
│   │   └── response.go
│   ├── middleware
│   │   └── <validation>
│   │       └── <validation>.go
│   ├── cofig
│   │   └── config.go
│   ├── order
│   │   ├── http_handler.go
│   │   ├── mysql_repository.go
│   │   └── service.go
│   ├── order_item
│   │   ├── http_handler.go
│   │   ├── mysql_repository.go
│   │   └── service.go
│   ├── product
│   │   ├── http_handler.go
│   │   ├── mysql_repository.go
│   │   └── service.go
│   ├── infrastructure
│   │   ├── gorm.go
│   │   ├── error_handler.go
│   │   ├── container.go
│   │   └── fiber.go
│   └── utilities
│       └── validation.go
├── pkg
│   └── xlogger
│       └── xlogger.go
├── go.mod
├── gen.go
├── data.sql
└── readme.md
```

## Cara Penggunaan
Project ini menggunakan Go Modules, sehingga tidak perlu melakukan go get untuk mengunduh library yang digunakan.

Untuk menjalankan project ini, silahkan jalankan perintah berikut:

```Bash
go run cmd/app/main.go
```

Untuk melakukan build project ini, silahkan jalankan perintah berikut:

```Bash
go build -o bin/<app-name> cmd/app/main.go
```

Pastikan untuk mengganti <app-name> dengan nama aplikasi yang diinginkan. Jika anda menggunakan Windows, maka perintah di atas akan menghasilkan file bin/<app-name>.exe.

## Environtment

Daftar environtment yang digunakan:

| Key              | Description                          | Example                                                                                              | Default                   |
|------------------|--------------------------------------|------------------------------------------------------------------------------------------------------|---------------------------|
| `HOST`           | Alamat untuk binding service         | `localhost`                                                                                          | `localhost`               |
| `PORT`           | Port untuk binding service           | `3000`                                                                                               | `3000`                    |
| `IS_DEVELOPMENT` | Mode development                     | `true`                                                                                               | `false`                   |
| `PROXY_HEADER`   | Header untuk mendapatkan IP asli     | `X-Real-IP` atau `X-Forwarded-For`                                                                   |                           |
| `LOG_FIELDS`     | Field yang akan ditampilkan pada log | `method,path,ip` lihat [disini](https://github.com/gofiber/contrib/blob/main/fiberzerolog/config.go) | `latency,status,method,url,error` |
| `DB_DRIVER`      | Driver database                      | `mysql`                                                                                              | `mysql` (in memory)       |
| `DB_DSN`         | Name database                        | `enterkomputer`                                                                                      |                           |
| `DB_USERNAME`    | Username database                    | `user`                                                                                               | `root`                    |
| `DB_PASSWORD`    | Password database                    | `password`                                                                                           |                           |
| `DB_PORT`        | Port database                        | `3306`                                                                                               | `3306`                    |

## Swagger Generation

Project ini menggunakan library [swaggo/swag](https://github.com/swaggo/swag) sebagai library untuk generate swagger. Untuk
mengenerate swagger, silahkan jalankan perintah berikut:

```bash
go generate ./...
```
