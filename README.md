#简单Web框架hoo

##Quickly strart

```go
package main

import(
	"fmt"
	"net/http"
	
	"github.com/astrosta/hoo"
)

func main(){
    h := hoo.New()
    
    h.Get("/", func(w http.ResponseWriter, req *http.Request) {
        fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
    })
    
    h.Run(":8080")
}
```