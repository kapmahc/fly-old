# FLY

A complete open source e-commerce solution for the Go language.

## Some packages

```bash
sudo apt-get install libmagickwand-dev
```

## Install nodejs

```bash
curl -o- https://raw.githubusercontent.com/creationix/nvm/v0.33.1/install.sh | zsh
nvm install 6
```

## Install go

```bash
zsh < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)
gvm install go1.8 -B
gvm use go1.8 --default
```

## Usage

```bash
go get -u github.com/kapmahc/fly
cd $GOPATH/src/github.com/kapmahc/fly
# download packages
make init
# build
make
ls dist
```

## Devleopment

```bash
cd $GOPATH/src/github.com/kapmahc/fly
# backend
./run.sh
# frontend
cd front && npm run dev
```

## Create database

```bash
psql -U postgres
CREATE DATABASE db-name WITH ENCODING = 'UTF8';
CREATE USER user-name WITH PASSWORD 'change-me';
GRANT ALL PRIVILEGES ON DATABASE db-name TO user-name;
```

## Issues

- "RPC failed; HTTP 301 curl 22 The requested URL returned error: 301"

  ```bash
  git config --global http.https://gopkg.in.followRedirects true
  ```

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

- Generate sitemap.xml.gz everyday

  ```bash
  @daily cd /var/www/www.change-me.com && ./fly seo
  ```

  ## Documents

  - [vue](https://vuejs.org/v2/guide/)
  - [vue-router](http://router.vuejs.org/)
  - [vue-i18n](https://kazupon.github.io/vue-i18n/en/started.html)
  - [vuex](https://vuex.vuejs.org/)
  - [bootstrap](http://v4-alpha.getbootstrap.com/getting-started/introduction/)
  - [gorm](http://jinzhu.me/gorm/)
  - [cli](https://github.com/urfave/cli)
  - [govendor](https://github.com/kardianos/govendor)
  - [gin](https://github.com/gin-gonic/gin)
  - [viper](https://github.com/spf13/viper)
  - [validator](https://godoc.org/gopkg.in/go-playground/validator.v9)
  - [rabbitmq](https://www.rabbitmq.com/documentation.html)
