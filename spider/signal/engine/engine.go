package engine

import (
	"go-tools/spider/signal/fetcher"
	"log"
)

func Run(seeds ...Request)  {
	var requests []Request
	for _,r :=range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("Fetcher err: url[%s]:%v", r.Url, err)
			continue
		}
		parseResult := r.ParserFunc(body)
		//requests = append(requests, parseResult.Requests...)
		for _,item := range parseResult.Items {
			log.Printf("Get item:%s", item)
		}

		for _,item := range parseResult.Requests {
			log.Printf("Get url:%s", item.Url)
		}
	}
}
