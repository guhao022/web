# web

对 gorilla/mux 进一步封装,功能简单,只提供了 Group Get Post Delete Put 几个方法

### 使用

```go
package main

import (
	"github.com/num5/web"
	"net/http"
	"fmt"
)

func main() {
	r := web.New()

    // 使用group
	g := r.Group("/g")      // localhost:9900/g/a
	g.Get("/a", H)
	g.Post("/p", H)         // localhost:9900/g/p
	// ...

	// 不使用
	//r.Get("/a", H)    // localhost:9900/a
	// ...

	http.Handle("/", r)

	fmt.Println("监听端口 :9900...")
	err := http.ListenAndServe(":9900", nil)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
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

每个方法返回的都是 *mux.Route, 所以可以继续使用 gorilla/mux 提供的方法
``` go
r.Get("/a", H).Schemes("https")
```

