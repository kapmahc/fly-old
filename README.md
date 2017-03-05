# fly

FLY is a complete open source ecommerce solution for C++.

## Development(for archlinux)

```bash
sudo pacman -S clang poco libiodbc psqlodbc
```

## Build

```bash
git clone https://github.com/kapmahc/fly.git
cd fly
git clone https://github.com/no1msd/mstch.git external/mstch
mkdir build && cd build
CC=/usr/bin/clang CXX=/usr/bin/clang++ cmake ..
make
```

## Documents

- [Google C++ Style Guide](https://google.github.io/styleguide/cppguide.html)
- [POCO](https://pocoproject.org/docs/)
- [mustache](http://mustache.github.io/)
