# Marshaler

Adds JSON (for now) Marshal and Unmarshal to different standard GO types

## Installation

```
$> go get github.com/xescugc/marshaler
```

## Usage

It supports 2 types:

* [`url.URL`](https://golang.org/pkg/net/url/#URL): With `marshaler.URL`
* [`currency.Unit`](https://pkg.go.dev/golang.org/x/text/currency#Unit): With `marshaler.CurrencyUnit`

## Example

```go
import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/xescugc/marshaler"
)

type User struct {
  Name  string        `json:"name"`
  URL   marshaler.URL `json:"url"`
}

func main() {
  u, _ := url.Parse("http://example.com")
  usr := User{
    Name: "Pepito",
    URL:  marshaler.URL{
      URL: u,
    },
  }

  b, _ := json.Marshal(usr)

  fmt.Println(string(b))
  // { "name": "Pepito", "url": "http://example.com" }
}
```
