# Instructor LED
## Import library

Setelah melakukan clone. Langkah selanjutnya adalah melakukan import semua dependensi yang digunakan seperti berikut:

```shell
go get -u ./...
```

## Config
Silahkan buat file baru dengan nama .env yang di duplikasi dari .env.example kemudian isi pada bagian berikut:
```shell
API_PORT=
DB_HOST=
DB_PORT=
DB_NAME=
USER=
PASSWORD=
DB_DRIVER=

LOG_FILE=logger.log
CSV_FILE=report.csv
TOKEN_LIFE_TIME=1 # jam
TOKEN_ISSUE_NAME=admin1@gmail.com # issuer bisa menggunakan username dari user
TOKEN_KEY=secret # key untuk token
```

Terdapat file ddl.sql dml.sql, silahkan buat database baru di postgre lalu import file tersebut

## Run
Jika sudah, jalankan program dengan perintah berikut :
```shell
go run .
```
