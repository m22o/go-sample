// cの所の #include<stdio.h>
package main

// 宣言したパッケージのみ使用できる
import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://localhost:18888")
	// 例外処理はなく関数の返り値としてエラーを返す
	//  http.Getのように関数の返り値の最後の項目としてエラーを返す。
	if err != nil {
		// panic はエラーを表示して、処理を終了
		panic(err)
	}
	// deferは処理後のコールバック
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))
}
