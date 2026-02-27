package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHealthCheckHandler(t *testing.T) {
	// สร้าง Request จำลอง
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	// สร้าง ResponseRecorder (ตัวบันทึกผลลัพธ์จำลอง)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HealthCheckHandler) // สมมติว่าฟังก์ชันชื่อนี้

	// รัน Handler
	handler.ServeHTTP(rr, req)

	// ตรวจสอบ Status Code (ต้องเป็น 200 OK)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// ตรวจสอบ Body (สมมติว่าต้องตอบว่า {"status": "up"})
	got := strings.TrimSpace(rr.Body.String())
	expected := `{"status":"up"}`
	if got != expected {
		t.Errorf("handler returned unexpected body: got %q want %q",
			got, expected) // ใช้ %q เพื่อให้เห็นตัวอักษรพิเศษถ้ามี
	}

}
