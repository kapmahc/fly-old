# fly

A complete open source e-commerce solution for the Go language and React.

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

  - [gorm](http://jinzhu.me/gorm/)
  - [cli](https://github.com/urfave/cli)
  - [govendor](https://github.com/kardianos/govendor)
  - [gin](https://github.com/gin-gonic/gin)
  - [viper](https://github.com/spf13/viper)
  - [validator](https://godoc.org/gopkg.in/go-playground/validator.v9)
  - [react](https://facebook.github.io/react/docs/installation.html#creating-a-single-page-application)
  - [react-redux](http://redux.js.org/docs/basics/UsageWithReact.html)
  - [react-bootstrap](https://react-bootstrap.github.io/components.html)
  - [Create React App](https://github.com/facebookincubator/create-react-app)
