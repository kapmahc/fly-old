#!/bin/sh
while true; do
  fuser -k 8080/tcp
  cd build && make -j && ./fly &
  inotifywait -e modify -e move -e create -e delete -e attrib -r `pwd`
done
