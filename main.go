package main

import (
	"log"
	"net/http"

	"github.com/viniciusfeitosa/MyFirstAppWithGo/handlers"
	"github.com/viniciusfeitosa/MyFirstAppWithGo/sqlite"
)

const port = ":8080"

func main() {

	db, err := sqlite.NewConn()
	if err != nil {
		log.Fatal(err)
	}
	if err := sqlite.CreateTables(db, sqlite.QuerysToMigrate); err != nil {
		log.Fatal(err.Error())
	}

	homeHandle := new(handlers.Home)
	userHandle := handlers.NewUser(db)
	sendHandle := handlers.NewSender(db)

	// This handler serve all static files on the app
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./public/assets"))))
	http.Handle("/", homeHandle)
	http.Handle("/list", userHandle)
	http.Handle("/send_email", sendHandle)
	http.HandleFunc("/create_user", http.HandlerFunc(userHandle.Create))
	http.HandleFunc("/remove_user", http.HandlerFunc(userHandle.Remove))
	http.HandleFunc("/update_user", http.HandlerFunc(userHandle.Update))
	http.HandleFunc("/find_all_users", http.HandlerFunc(userHandle.FindAll))
	http.HandleFunc("/find_user", http.HandlerFunc(userHandle.FindOne))

	log.Println("Start app on port", port)
	http.ListenAndServe(port, nil)
}
