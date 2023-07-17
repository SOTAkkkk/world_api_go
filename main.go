package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type City struct {
	id          int
	name        string
	countryCode string
	district    string
	population  int
}

var city City

func main() {
	connectToDB()
}

func connectToDB() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}
	HOST := os.Getenv("HOST")
	DBTable := os.Getenv("DB_TABLE")
	DBUser := os.Getenv("DB_USER")
	DBPass := os.Getenv("DB_PASS")
	DBPort := os.Getenv("DB_PORT")

	// データベースのハンドルを取得する
	// db, err := sql.Open("mysql", "root:rootpass@tcp(localhost:13306)/world")
	db, err := sql.Open("mysql", DBUser+":"+DBPass+"@tcp("+HOST+":"+DBPort+")/"+DBTable+"")
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

	//  | ID | Name | CountryCode | District | Population
	// SQLの実行
	rows, err := db.Query("SELECT * FROM city ")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&city.id, &city.name, &city.countryCode, &city.district, &city.population)
		if err != nil {
			panic(err.Error())
		}
		log.Println(city.id, city.name, city.countryCode, city.district, city.population)
	}

	row := db.QueryRow("SELECT * FROM city WHERE id = ?", 1)
	err = row.Scan(&city.id, &city.name, &city.countryCode, &city.district, &city.population)

	if err != nil {
		panic(err.Error())
	}
	log.Println(city.id, city.name, city.countryCode, city.district, city.population)
}
