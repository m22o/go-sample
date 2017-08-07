package main

import (
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/http/httputil"
)

func main() {
	// クッキーを保存するインスタンス作成
	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}
	// クッキーが保存可能なhttp.Clientインスタンスを作成
	client := http.Client{
		Jar: jar,
	}
	// 最初のアクセスでクッキーを受信して、2回目以降はクッキーを送信している
	for i := 0; i < 2; i++ {
		// cookiejar のライブラリの clientのGetを使って、アクセスする
		resp, err := client.Get("http://localhost:18888/cookie")
		if err != nil {
			panic(err)
		}
		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			panic(err)
		}
		log.Println(string(dump))
	}
}
