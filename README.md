# Tilda Golang client

[![License](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/dimuska139/tilda-go/blob/master/LICENSE)

This is unofficial Tilda SDK for GO applications. This library contains methods for interacting with the [Tilda API](https://help.tilda.cc/api).

## Installation

```shell
go get github.com/dimuska139/tilda-go
```

## Usage

```go
package main

import (
	"context"
	"fmt"
	"net/http"

	tilda "github.com/dimuska139/tilda-go"
)

func main() {
	client := tilda.NewClient(&tilda.Config{
		PublicKey: "your_public_key",
		SecretKey: "your_secret_key",
	}, tilda.WithCustomHttpClient(http.DefaultClient))

	projects, err := client.GetProjectsList(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Println(projects)
}
```

The tests should be considered a part of the documentation. Also you can read [official docs](https://help.tilda.cc/api).

## License

Tilda GO is released under the
[MIT License](http://www.opensource.org/licenses/MIT).