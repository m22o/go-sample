package main

import (
  "io/ioutil"
  "log"
  "net/http"
)

func main(){
  resp, _ :=http.Get("http://localhost:18888")
  defer resp.Body.Close()
  body,_:= ioutil.ReadAll(resp.Body)
  log.Println(string(body))
  log.Print("status:",resp.Status) //ステータスコード
  log.Print("Headers:",resp.Header)
  log.Println("Conten-Length:",resp.Header.Get("Content-Length"))
  
}
