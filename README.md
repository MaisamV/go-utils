# github.com/mitchallen/go-lib

## Usage

### Initialize your module

```
$ go mod init example.com/my-golib-demo
```

### Get the go-utils module

Note that you need to include the **v** in the version tag.

```
$ go get github.com/MaisamV/go-utils@v0.1.0
```

```go
package main

import (
    "fmt"

    utils "github.com/MaisamV/go-utils"
)

func main() {
	item := 2
	list := []int{20, 2, -5}
    fmt.Println(utils.Contains(item, list))

	stringItem := "green"
	stringList := []string{"blue", "green", "red"}
	fmt.Println(utils.Contains(stringItem, stringList))
}
```

## Testing

```
$ go test
```

## Tagging

```
$ git tag v0.1.0
$ git push origin --tags
```