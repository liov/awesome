package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	logdy "github.com/logdyhq/logdy-core/logdy"
)

// 没啥用，没用任何搜索排序逻辑，纯把日志放前端，搜索都是前端的搜索
type Logger struct {
	logdy   logdy.Logdy
	handler http.Handler
}

func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	l.handler.ServeHTTP(w, r)

	// If this is a request to Logdy backend, ignore it
	if strings.HasPrefix(r.URL.Path, l.logdy.Config().HttpPathPrefix) {
		return
	}

	l.logdy.Log(logdy.Fields{
		"ua":     r.Header.Get("user-agent"),
		"method": r.Method,
		"path":   r.URL.Path,
		"query":  r.URL.RawQuery,
		"time":   time.Since(start),
	})
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/v1/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	mux.HandleFunc("/v1/time", func(w http.ResponseWriter, r *http.Request) {
		curTime := time.Now().Format(time.Kitchen)
		w.Write([]byte(fmt.Sprintf("the current time is %v", curTime)))
	})

	logger := logdy.InitializeLogdy(logdy.Config{
		HttpPathPrefix: "/logdy-ui",
	}, mux)

	addr := ":8082"
	log.Printf("server is listening at %s", addr)
	log.Fatal(http.ListenAndServe(addr, &Logger{logdy: logger, handler: mux}))
}
