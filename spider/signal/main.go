package main

import (
	"go-tools/spider/signal/engine"
	"go-tools/spider/signal/target/zhenai/parser"
)

/*
单机版爬虫
 */
func main() {
	engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParserCitylist,
	})
}
