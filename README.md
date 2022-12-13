# URL по правилам Yandex на Golang

[![Release](https://img.shields.io/github/release/alexnimoy/go-ya-slug.svg?style=flat-square)](https://github.com/alexnimoy/go-ya-slug/releases/latest)
[![CI](https://github.com/alexnimoy/go-ya-slug/actions/workflows/go.yml/badge.svg)](https://github.com/alexnimoy/go-ya-slug/actions/workflows/go.yml)

Правильные URL повышают шансы выйти на первые позиции в SERP. Страницы с корректным транслитерированным адресом поисковик ранжирует выше.

## Install

```
$ go get -u github.com/alexnimoy/go-ya-slug
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
