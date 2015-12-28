# mux

对 gorilla/mux 进一步封装,功能简单,只提供了 Group Get Post Delete Put 几个方法

### 使用

```go
package main

import (
    ctx "github.com/num5/context"
	"github.com/num5/mux"
	"net/http"
	"fmt"
)

func main() {
	r := mux.New()

    // 使用group
	g := r.Group("/g")
	g.Get("/a", H)
	g.Post("/p", H)
	// ...

	// 不使用
	//r.Get("/a", H)
	// ...

	http.Handle("/", r)


	fmt.Println("监听端口 :9900...")
	err := http.ListenAndServe(":9900", nil)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
}

func H(ctx *ctx.Context) {
    name := ctx.GetString("name")
    if name == "" {
    	ctx.Json("miss query param: name")
    	return
    }

    c.Json(name)
}
```

