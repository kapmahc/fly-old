#!/bin/sh
while true; do
  fuser -k 8080/tcp
  cd demo && go run main.go s &
  inotifywait -e modify -e move -e create -e delete -e attrib -r .
done
