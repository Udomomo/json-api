package main

import (
	"log"
	"net/http"
	"time"
)

//Logger http.Handlerをラップし、Handlerが実行されるごとに毎回ログが吐かれるようにする
func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}
