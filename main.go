package main

import (
	"fmt"
	"log"
	"net/http"

	"world_api_go/database" // 追加: database パッケージをインポート

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// データベースの初期化
	err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	cityHandler()
}

func cityHandler() {
	// マルチプレクサの初期化
	mux := http.NewServeMux()

	// ハンドラ関数とURLパスの登録
	mux.HandleFunc("/", database.HomeHandler)
	mux.HandleFunc("/city", database.GetCityHandler)

	// マルチプレクサを使用してHTTPサーバーを立ち上げ
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("サーバーを起動できませんでした。エラー:", err)
	}
}
