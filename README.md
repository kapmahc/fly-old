# FLY

A complete open source e-commerce solution by Go.

## Installing

```bash
bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)
gvm install go1.8rc3 -B
gvm use go1.8rc3 --default

go get -u github.com/kapmahc/fly
cd $GOPATH/src/github.com/kapmahc/fly
make
ls -l dist
```

## Editors

### Atom plugins

- git-plus
- go-plus
- atom-beautify
- atom-react
- autosave(remember to enable it)
- [Chrome extension](https://chrome.google.com/webstore/detail/react-developer-tools/fmkadmapgofadopljbjfkapdkoienihi)

## Notes

### RabbitMQ

- The web UI is located at: <http://server-name:15672/>, (user "guest" is created with password "guest")

  ```bash
  rabbitmq-plugins enable rabbitmq_management
  ```

## Documents

- [Build Web Application with Golang](https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/preface.html)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- [JSON Style Guide](https://google.github.io/styleguide/jsoncstyleguide.xml)
- [gorm](http://jinzhu.me/gorm/)
- [martini](https://github.com/go-martini/martini)
- [cli](https://github.com/urfave/cli)
- [viper](https://github.com/spf13/viper)
- [machinery](https://github.com/RichardKnop/machinery)
- [validator](https://github.com/go-playground/validator)
- [RabbitMQ](https://www.rabbitmq.com/getstarted.html)
