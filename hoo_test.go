package hoo

import (
	"fmt"
	"net/http"
	"testing"
)

func TestRouterEngine_Run(t *testing.T) {
	h := New()

	h.Get("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	})

	h.Run(":8080")
}
