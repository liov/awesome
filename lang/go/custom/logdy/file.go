package main

import (
	"bufio"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	logdyhttp "github.com/logdyhq/logdy-core/http"
	"github.com/logdyhq/logdy-core/logdy"
)

// 没啥用，没用任何搜索排序逻辑，纯把日志放前端，搜索都是前端的搜索
type Logger struct {
	logdy logdy.Logdy
}

func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.DefaultServeMux.ServeHTTP(w, r)
	paths := r.URL.Query()["path"]
	if len(paths) == 0 {
		return
	}
	path := paths[0]
	log.Println(path)
	if strings.HasSuffix(strings.Trim(path, "\""), "log") {
		go func() {
			file, _ := os.Open(path)
			defer file.Close()
			scanner := bufio.NewScanner(file)
			bufferSize := 10 * 1024 * 1024
			scanner.Buffer(make([]byte, 0, bufferSize), bufferSize)
			var n int64
			for scanner.Scan() {
				line := scanner.Text()
				err := l.logdy.LogString(line)
				if err != nil {
					log.Println(err)
				}
				n++
				if n == logdyhttp.BULK_WINDOW_MS {
					time.Sleep(time.Millisecond * 100)
					n = 0
				}
			}
		}()
	}

}

func main() {
	logger := logdy.InitializeLogdy(logdy.Config{
		HttpPathPrefix: "/logdy-ui",
	}, http.DefaultServeMux)

	addr := ":8082"
	log.Printf("server is listening at %s", addr)
	log.Fatal(http.ListenAndServe(addr, &Logger{logdy: logger}))
}
