package engine

import (
	"godemo/crawler/fetcher"
	"log"
)

/**
爬取引擎
>调用爬取方法
*/
type SimpleEngine struct{}

func (e SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds { //将参数放进数组中
		requests = append(requests, r)
	}
	//通过url开始爬数据
	for len(requests) > 0 {
		r := requests[0] //始终爬去数组的第一个数组
		requests = requests[1:]

		parseResult, err := worker(r)
		if err != nil {
			continue
		}

		requests = append(requests, parseResult.Requests...) //parseRequest.Requests...就是parseRequest.Request[0] parseRequest.Request[1]..

		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item)
		}
	}
}
func worker(r Request) (ParseResult, error) {
	log.Printf("Fetching %s", r.Url)
	body, err := fetcher.Fetch(r.Url) //通过url返回全部数据
	if err != nil {                   //错误处理 避免一个爬去错误影响后续的请求
		log.Printf("Fetcher :error "+"fetching url %s,%v", r.Url, err)
		return ParseResult{}, err
	}
	return r.ParserFunc(body), nil //处理数据 request 放到requests中 item打印出来
}