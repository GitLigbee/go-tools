package parser

import (
	"go-tools/spider/signal/fetcher"
	"testing"
)

func TestParserCitylist(t *testing.T) {
	content,err := fetcher.Fetch("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	var dataSize = 470
	result := ParserCitylist(content)
	if len(result.Requests) != dataSize {
		t.Errorf("result'size should be %d, but get %d", dataSize, len(result.Requests))
	}
}