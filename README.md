# go-taninari

[たになり語録](https://taninari.amebaownd.com/) _人生楽しんでますか？_


## Installation

```bash
$ go get github.com/yukirii/go-taninari/...
```


## Usage

### Command

```bash
$ taninari
$ taninari help
```

### Library

```go
package main

import (
	"fmt"
	"log"

	"github.com/yukirii/go-taninari"
)

func Show(msg *taninari.GorokuMessage) {
	fmt.Println(msg.PublishedAt)
	fmt.Println(msg.Text)
	if msg.ImageURL != "" {
		fmt.Println(msg.ImageURL)
	}
	fmt.Println(msg.PublishedURL)
}

func main() {
	// 語録からランダムに 1 件取得
	msg, err := taninari.GetRandomMessage()
	if err != nil {
		log.Fatal(err)
	}
	Show(msg)

	// 語録の全取得
	msgs, err := taninari.GetAllMessages()
	if err != nil {
		log.Fatal(err)
	}
	for _, msg := range msgs {
		Show(msg)
		fmt.Print("\n")
	}

	// 語録から検索
	msgs, err = taninari.SearchMessages("グレープフルーツ")
	if err != nil {
		log.Fatal(err)
	}
	for _, msg = range msgs {
		Show(msg)
		fmt.Print("\n")
	}
}
```


## License

[MIT](https://github.com/yukirii/go-taninari/blob/master/LICENSE)
