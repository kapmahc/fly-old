# fly

A complete open source ecommerce solution for for the Go language.

## Install go

```bash
zsh < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)
gvm install go1.8rc3 -B
gvm use go1.8rc3 --default
```

## Usage

```bash
go get -u github.com/kardianos/govendor
go get -u github.com/kapmahc/fly
cd $GOPATH/src/github.com/kapmahc/fly
make
ls dist
```

## Create database

```bash
psql -U postgres
CREATE DATABASE db-name WITH ENCODING = 'UTF8';
CREATE USER user-name WITH PASSWORD 'change-me';
GRANT ALL PRIVILEGES ON DATABASE db-name TO user-name;
```

## Issues

- 'Peer authentication failed for user', open file "/etc/postgresql/9.5/main/pg_hba.conf" change line:

  ```
  local   all             all                                     peer  
  TO:
  local   all             all                                     md5
  ```

- Generate openssl certs

  ```bash
  openssl genrsa -out www.change-me.com.key 2048
  openssl req -new -x509 -key www.change-me.com.key -out www.change-me.com.crt -days 3650 # Common Name:*.change-me.com
  ```

  ## Documents

  - [gorm](http://jinzhu.me/gorm/)
  - [mux](https://github.com/gorilla/mux)
  - [negroni](https://github.com/urfave/negroni)
  - [cli](https://github.com/urfave/cli)  
  - [govendor](https://github.com/kardianos/govendor)
