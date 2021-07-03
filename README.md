简单Web框架hoo
===

Quickly strart
---

```go
package main

import(
	"net/http"
	
	"github.com/astrosta/hoo"
)

func main(){
	h := hoo.New()
	h.GET("/", func(c *hoo.Context) {
		c.HTML(http.StatusOK, "<h1>Hello hoo</h1>")
	})
    
    h.Run(":8080")
}
```