package main

import (
	"context"
	"github.com/PuerkitoBio/goquery"
	"github.com/hopeio/utils/net/http/client"
	httpi "github.com/hopeio/utils/net/http/consts"
	"github.com/hopeio/utils/scheduler/crawler"
	"net/http"
	"path"
	"strconv"
)

func main() {
	engine := crawler.NewEngine(10)
	engine.ErrHandlerUtilSuccess()
	engine.Run(fetch("1"))
}

func fetch(page string) *crawler.Request {
	return &crawler.Request{
		Key: page,
		Run: func(ctx context.Context) ([]*crawler.Request, error) {
			reader, err := client.New().AddHeader(httpi.HeaderUserAgent, client.UserAgentIphone).Request(http.MethodGet,
				"https://m.yeitu.com/meinv/xinggan/20240321_33578_"+page+".html").DoStream(nil)
			if err != nil {
				return nil, err
			}
			doc, err := goquery.NewDocumentFromReader(reader)
			src, _ := doc.Find(".gallery-item img").Attr("src")
			var reqs []*crawler.Request
			reqs = append(reqs, downloadImg(src))
			if page == "1" {
				numStr := doc.Find(".imageset-sum").Text()
				numStr = numStr[2:]
				num, _ := strconv.Atoi(numStr)
				for i := 2; i <= num; i++ {
					reqs = append(reqs, fetch(strconv.Itoa(i)))
				}
			}
			return reqs, nil
		},
	}
}

func downloadImg(src string) *crawler.Request {
	return &crawler.Request{
		Key: src,
		Run: func(ctx context.Context) ([]*crawler.Request, error) {
			err := client.Download("E:/tmp/"+path.Base(src), src)
			if err != nil {
				return nil, err
			}
			return nil, nil
		},
	}
}
