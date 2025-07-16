package controller

import (
	"encoding/json"
	"go-package-gorm/config"
	"go-package-gorm/model"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var users []model.User

	err := config.DB.Find(&users).Error

	if err != nil {
		w.WriteHeader(500)
		res, _ := json.Marshal(map[string]string{"error": "Failed to retrieve users"})
		w.Write(res)
		return
	}
	
	res, _ := json.Marshal(users)
	// w.WriteHeader(200)
	w.Write(res)
}