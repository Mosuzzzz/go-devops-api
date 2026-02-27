package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	connStr := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := db.Ping()
		if err != nil {
			fmt.Fprintf(w, "Database Connection Error: %v", err)
			return
		}
		fmt.Fprint(w, "🚀 Go API is running and PostgreSQL is connected!")
	})


	fmt.Println("Server starting at :8080")
	log.Fatal(http.ListenAndServe(":8080",nil))
}