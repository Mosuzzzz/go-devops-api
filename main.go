package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	_ "github.com/lib/pq"
)

// HealthCheckHandler ต้องชื่อตรงกับที่เรียกในไฟล์ Test
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "up"})
}

func main() {
	connStr := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/health", HealthCheckHandler)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := db.Ping()
		if err != nil {
			fmt.Fprintf(w, "Database Connection Error: %v", err)
			return
		}
		fmt.Fprintln(w, "🚀 Go API is running and PostgreSQL is connected!")
	})


	fmt.Println("Server starting at :8080")
	log.Fatal(http.ListenAndServe(":8080",nil))
}