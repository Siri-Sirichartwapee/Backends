package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// โครงสร้างรับข้อมูล
type Request struct {
	Name string `json:"name"`
}

// โครงสร้างตอบกลับ
type Response struct {
	Message string `json:"message"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// ต้องเป็น POST เท่านั้น
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// อ่าน JSON จาก body
	var req Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.Name == "" {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// สร้าง response
	res := Response{
		Message: "Hello " + req.Name,
	}

	// ตั้ง header เป็น JSON
	w.Header().Set("Content-Type", "application/json")

	// ส่ง JSON กลับ
	json.NewEncoder(w).Encode(res)
}

func main() {
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}