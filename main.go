package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type RequestBody struct {
	Message string `json:"message"`
}

type ResponseBody struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/process", processJSON)
	fmt.Println("Server is listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}

func processJSON(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var reqBody RequestBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	if reqBody.Message != "" {
		fmt.Println("Received message:", reqBody.Message)
		res := ResponseBody{
			Status:  "success",
			Message: "Данные успешно приняты",
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	} else {
		http.Error(w, "Некорректное JSON-сообщение", http.StatusBadRequest)
	}
}

