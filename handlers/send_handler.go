package handlers

// import (
// 	"database/sql"
// 	"log"
// 	"net/http"

// 	"github.com/viniciusfeitosa/MyFirstAppWithGo/senders"

// 	"github.com/viniciusfeitosa/MyFirstAppWithGo/models"
// )

// // Send is the handler struct to send email
// type Send struct {
// 	db *sql.DB
// }

// // NewSender starts a new instance of send
// func NewSender(db *sql.DB) *Send {
// 	return &Send{db: db}
// }

// func (s *Send) ServeHTTP(w http.ResponseWriter, req *http.Request) {
// 	userModel := models.NewUser()
// 	users, err := userModel.SelectAll(s.db)
// 	if err != nil {
// 		log.Println(err.Error())
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	var emailList []string
// 	for _, u := range users {
// 		emailList = append(emailList, u.Email)
// 	}

// 	if err := senders.SendEmail(emailList); err != nil {
// 		log.Println(err.Error())
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json; charset=utf-8")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte("{\"message\": \"Success to send email\"}"))

// }
