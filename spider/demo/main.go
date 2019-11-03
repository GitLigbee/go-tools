package demo

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main(){
	resp,err := http.Get("http://www.zhenai.com/zhenghun")
	if  err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("http status:", resp.StatusCode)
		return
	}
	// 解码
	encoder := determineEncoding(resp.Body)
	newReader := transform.NewReader(resp.Body, encoder.NewDecoder())

	content,err := ioutil.ReadAll(newReader)
	//content,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	//fmt.Printf("%s\n", content)
	findTarget(content)
}

// 获取编码
func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}

// 寻找目标
func findTarget(content []byte) {
	reg := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)"[^>]*>([^<]+)</a>`)
	all := reg.FindAllSubmatch(content, -1)
	for _,m := range all {
		fmt.Printf("url: %s, city: %s\n", m[1], m[2])
	}
}
