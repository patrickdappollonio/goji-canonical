# goji-canonical

Package `canonical` allows to set up a Goji middleware to redirect all users to a default, canonical domain. This is useful if, for example, you upload your website to Heroku, but instead of allowing your users to access both `example.herokuapp.com` and `example.com`, you can set up this middleware to enforce your users to access through `example.com` by redirecting them if they access over another domain (such as the `herokuapp.com` one).

### Usage example
```go
package main

import (
	"github.com/theosomefactory/goji-canonical"
	"github.com/zenazn/goji"
)

func main() {
	goji.Use(canonical.Canonical("example.com"))
	goji.Get("/", myHandler)
	goji.Serve()
}
```