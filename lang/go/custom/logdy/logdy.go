package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/logdyhq/logdy-core/logdy"
)

// 没啥用，没用任何搜索排序逻辑，纯把日志放前端，搜索都是前端的搜索
type Logger struct {
	logdy logdy.Logdy
}

func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	http.DefaultServeMux.ServeHTTP(w, r)
	
	l.logdy.Log(logdy.Fields{
		"ua":     r.Header.Get("user-agent"),
		"method": r.Method,
		"path":   r.URL.Path,
		"query":  r.URL.RawQuery,
		"time":   time.Since(start),
	})
}

func main() {

	http.DefaultServeMux.HandleFunc("/v1/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	http.DefaultServeMux.HandleFunc("/v1/time", func(w http.ResponseWriter, r *http.Request) {
		curTime := time.Now().Format(time.Kitchen)
		w.Write([]byte(fmt.Sprintf("the current time is %v", curTime)))
	})

	logger := logdy.InitializeLogdy(logdy.Config{
		HttpPathPrefix: "/logdy-ui",
	}, http.DefaultServeMux)

	addr := ":8082"
	log.Printf("server is listening at %s", addr)
	log.Fatal(http.ListenAndServe(addr, &Logger{logdy: logger}))
}
