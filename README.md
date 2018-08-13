# Google-Go

Google-Go is a [Google OAuth API](https://developers.google.com/identity/protocols/OAuth2) client written in Go. Currently it can:

- Get Token Info

## Installation

```bash
$ go get github.com/swensonhe/google-go
```

## Examples

### Get Me

```go
package main

import (
    "github.com/swensonhe/google-go"
    "fmt"
)

func main() {
    client := google.NewClient()
    user, err := client.GetTokenInfo("token")
    if err != nil {
        fmt.Println(err)
        return
    }
    
    fmt.Println(user)
}

```
