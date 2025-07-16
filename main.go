package main

import (
	"encoding/json"
	"go-package-gorm/config"
	"go-package-gorm/controller"
	"go-package-gorm/middleware"

	// "go-package-gorm/model"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

func main() {
	// Load configuration
	config.LoadConfig()

	// Connect to the database
	config.ConnectDB()

	// Run the Migrations
	// err := config.DB.AutoMigrate(&model.User{})
	// if err != nil {
	// 	log.Fatal("Gagal migrasi:", err)
	// } else {
	// 	log.Println("Migration completed successfully")
	// }

	r := mux.NewRouter()

	r.HandleFunc("/users", controller.Index).Methods("GET")
	
	r.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		//Cara 1
		// json.NewEncoder(w).Encode(map[string]bool{"ok": true})

		// Cara 2
		res, _ := json.Marshal(map[string]bool{"ok": true})
		w.Write(res)
	}).Methods("GET")

	r.Use(middleware.LoggingMiddleware)

	log.Println("Server is starting on port 8080")
	http.ListenAndServe(":8080", r)
}
