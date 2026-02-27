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
	// ดึงค่าจาก environment variable ที่เราตั้งไว้ใน docker-compose
	dbURL := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}

	// ทดสอบ Ping ฐานข้อมูล
	http.HandleFunc("/db-check", func(w http.ResponseWriter, r *http.Request) {
		err := db.Ping()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Database connection failed: %v", err)
			return
		}
		fmt.Fprintf(w, "Successfully connected to Database!")
	})

	log.Println("Server starting at :8080")
	http.ListenAndServe(":8080", nil)
}