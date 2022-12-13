# URL по правилам Yandex на Golang

[![Release](https://img.shields.io/github/release/deweppro/go-badges.svg?style=flat-square)](https://github.com/deweppro/go-badges/releases/latest)
[![CI](https://github.com/deweppro/go-badges/actions/workflows/ci.yml/badge.svg)](https://github.com/deweppro/go-badges/actions/workflows/ci.yml)

Правильные URL повышают шансы выйти на первые позиции в SERP. Страницы с корректным транслитерированным адресом поисковик ранжирует выше.

## Install

```
$ go get github.com/alexnimoy/go-ya-slug
```

## Usage

```Go
package main

import (
	"fmt"

	yaslug "github.com/alexnimoy/go-ya-slug"
)

func main() {
	url := yaslug.Url("Урл url")
	fmt.Println(url) //=> url-url

}
```

## License

MIT

## Author

Alexander Polyakov
