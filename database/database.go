package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var db *sql.DB

type City struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	CountryCode string `json:"countryCode"`
	District    string `json:"district"`
	Population  int    `json:"population"`
}

/*
データベースの初期化
*/
func InitDB() error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}

	/*
		データベースの環境設定
	*/
	HOST := os.Getenv("HOST")
	DBTable := os.Getenv("DB_TABLE")
	DBUser := os.Getenv("DB_USER")
	DBPass := os.Getenv("DB_PASS")
	DBPort := os.Getenv("DB_PORT")

	/*
		データベースのハンドルを取得する
	*/
	db, err = sql.Open("mysql", DBUser+":"+DBPass+"@tcp("+HOST+":"+DBPort+")/"+DBTable+"")
	if err != nil {
		return err
	}

	/*
		データベースへの接続確認
	*/
	err = db.Ping()
	if err != nil {
		return err
	}

	log.Println("データベース接続完了")
	return nil
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Home Page!")
}

func GetCityHandler(w http.ResponseWriter, r *http.Request) {

	/*
		SELECTクエリの実行
	*/
	rows, err := db.Query("SELECT * FROM city ")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	/*
		City に 情報をセットする
	*/
	var result []City
	for rows.Next() {
		var city City
		err := rows.Scan(&city.ID, &city.Name, &city.CountryCode, &city.District, &city.Population)
		if err != nil {
			http.Error(w, "Scan error", http.StatusInternalServerError)
			return
		}
		result = append(result, city)
	}

	// JSON形式でレスポンスを返す
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)

	defer db.Close()

}
