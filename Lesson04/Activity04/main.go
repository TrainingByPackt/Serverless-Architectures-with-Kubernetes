package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("MYSQL_DATABASE")
	address := os.Getenv("MYSQL_ADDRESS")

	apiKey := os.Getenv("API_KEY")

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, address, dbName))

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to database..")

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS GoldPrices (timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP, price FLOAT(7,6));")
	if err != nil {
		log.Fatal(err)
	}

	r, err := http.Get(fmt.Sprintf("http://apilayer.net/api/live?access_key=%s&currencies=XAU", apiKey))
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()

	target := Data{}
	err = json.NewDecoder(r.Body).Decode(&target)
	if err != nil {
		log.Fatal(err)
	}

	if !target.Success {
		log.Fatal("Not successful data from API layer")
	}

	log.Println("Successfully retrieved data..")

	stmt, err := db.Prepare("INSERT INTO GoldPrices(price) VALUES(?)")
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(target.Quotes.USDXAU)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Successfully inserted the price: %v", target.Quotes.USDXAU)

}

type Data struct {
	Success bool
	Quotes  struct {
		USDXAU float64
	}
}
