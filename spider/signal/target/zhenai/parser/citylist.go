package parser

import (
	"go-tools/spider/signal/engine"
	"regexp"
)

const regExpression  = `<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)"[^>]*>([^<]+)</a>`

func ParserCitylist(content []byte) engine.ParseResult {
	reg := regexp.MustCompile(regExpression)
	all := reg.FindAllSubmatch(content, -1)
	result := engine.ParseResult{}
	for _,m := range all {
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: engine.NilParser,
		})
	}
	return result
}
