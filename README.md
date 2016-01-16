# web

对 gorilla/mux 进一步封装

### 使用

```go
package main

import (
	"github.com/num5/web"
	"net/http"
	"fmt"
)

func main() {
	var routes = web.Routes{
    	Route{
    		"H",
    		"GET",
    		"/h",
    		H,
    	},
    	Route{
    		"H",
    		"POST",
    		"/h",
    		H,
    	},
    	//...
    }

    r := web.Register(routes)
}

func H(ctx *web.Context) {
    name := ctx.GetString("name")
    if name == "" {
    	ctx.Json("miss query param: name")
    	return
    }

    c.Json(name)
}
```

