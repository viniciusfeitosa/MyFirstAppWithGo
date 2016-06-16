package handlers

// import (
// 	"fmt"
// 	"text/template"

// 	"net/http"
// )

// // Index is an example of struct
// type Index struct {
// 	Name  string
// 	Email string
// }

// func (i *Index) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	if err := myTemplate(w, "index", nil); err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}
// }

// func myTemplate(w http.ResponseWriter, templateName string, structToPage interface{}) error {
// 	index := fmt.Sprintf("./public/templates/%s.html", templateName)
// 	t, err := template.ParseFiles(index)
// 	if err != nil {
// 		return err
// 	}
// 	t.Execute(w, structToPage)
// 	return nil
// }
