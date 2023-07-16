package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	connectToDB()
}

func connectToDB() {
	// データベースのハンドルを取得する
	db, err := sql.Open("mysql", "root:rootpass@tcp(localhost:13306)/world")
	if err != nil {
		// ここではエラーを返さない
		log.Fatal(err)
	}
	defer db.Close()

	// 実際に接続する
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("データベース接続完了")
	}
}
