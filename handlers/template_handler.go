package handlers

import (
	"fmt"
	"net/http"
	"text/template"
)

func serveTemplate(w http.ResponseWriter, templateName string, structToPage interface{}) error {
	head := fmt.Sprintf("./public/templates/%s.html", "head")
	page := fmt.Sprintf("./public/templates/%s.html", templateName)
	footer := fmt.Sprintf("./public/templates/%s.html", "footer")
	t, err := template.ParseFiles(head, page, footer)
	if err != nil {
		return err
	}
	t.ExecuteTemplate(w, templateName, structToPage)
	return nil
}
