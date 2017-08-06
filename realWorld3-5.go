// curl -G -data-urlencode "query=hello world" http://localhost:18888

package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {
	values := url.Values{ // クエリーを文字列に変換
		"query": {"hello world"},
	}
	// URLの末尾にクエリを追加
	resp, _ := http.Get("http://localhost:18888" + "?" + values.Encode())
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println(string(body))

}
