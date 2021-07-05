package hoo

import (
	"log"
	"net/http"
	"time"
)

func LoggerMiddleware(next MiddleWareHandleFunc) MiddleWareHandleFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("logger middleware begin")
		timeStart := time.Now()
		next(w, r)
		log.Printf("elapse time: %v", time.Since(timeStart))
	}
}

func TimeoutMiddleware(next MiddleWareHandleFunc) MiddleWareHandleFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("timeoutMiddleware begin")
		//待加入
		next(w, r)
		log.Println("timeoutMiddleware end")
	}
}

func RateLimitMiddleware(next MiddleWareHandleFunc) MiddleWareHandleFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("rateLimitMiddleware begin")
		//待加入
		next(w, r)
		log.Println("rateLimitMiddleware end")
	}
}
