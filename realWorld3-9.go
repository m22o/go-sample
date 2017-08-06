package main

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func main() {
	var buffer bytes.Buffer                //マルチパート部を組み立て後のバイト列を格納するバッファを宣言
	writer := multipart.NewWriter(&buffer) //マルチパートを組み立てるライターを作る
	writer.WriteField("name", "matsuo")    // file以外のフィールドは WriteFieldを使って登録する

	// part := make(textproto.MIMEHeader)
	// part.Set("Content-Type", "image/jped")
	// part.Set("Content-Disposition", `form-data; name=thumbnail; filename="photo.jpg"`)

	// ここからしでファイルを読み込む
	fileWriter, err := writer.CreateFormFile("thumbnail", "photo.jpg")
	if err != nil {
		panic(err)
	}

	readFile, err := os.Open("photo.jpg")
	if err != nil {
		panic(err)
	}

	defer readFile.Close()

	// io.writerに書き込む
	io.Copy(fileWriter, readFile)

	// マルチパートをクローズして、バッファに書き込む
	writer.Close()

	resp, err := http.Post("http://localhost:18888", writer.FormDataContentType(), &buffer)
	if err != nil {
		panic(err)
	}
	log.Println("status:", resp.Status)
}
