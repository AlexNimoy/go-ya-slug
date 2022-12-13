# URL по правилам Yandex на Golang

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
