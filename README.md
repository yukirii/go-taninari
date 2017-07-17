# go-taninari

[たになり語録](https://taninari.amebaownd.com/) _人生楽しんでますか？_


## Installation

```bash
$ go get github.com/shiftky/go-taninari/...
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

    "github.com/shiftky/go-taninari"
)

func Show(goroku *taninari.Goroku) {
    fmt.Println(goroku.PublishedAt)
    fmt.Println(goroku.Text)
    if goroku.ImageURL != "" {
        fmt.Println(goroku.ImageURL)
    }
    fmt.Println(goroku.PublishedURL)
}

func main() {
    // 語録からランダムに 1 件取得
    goroku, err := taninari.GetGoroku()
    if err != nil {
        log.Fatal(err)
    }
    Show(goroku)

    // 語録の全取得
    gorokus, err := taninari.GetAllGorokus()
    if err != nil {
        log.Fatal(err)
    }
    for _, goroku := range gorokus {
        Show(goroku)
        fmt.Print("\n")
    }
}
```


## License

[MIT](https://github.com/shiftky/go-taninari/blob/master/LICENSE)
