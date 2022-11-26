package handlers

import (
	"github.com/CloudyKit/jet/v6"
	"log"
	"net/http"
)

var views = jet.NewSet(
	jet.NewOSFileSystemLoader("./html"),
	jet.InDevelopmentMode())

func Home(w http.ResponseWriter, r *http.Request) {
	err := RenderPage(w, "home.jet", nil)
	if err != nil {
		log.Println(err)
		return
	}
}

func RenderPage(w http.ResponseWriter, templateName string, data jet.VarMap) error {
	view, err := views.GetTemplate(templateName)
	if err != nil {
		return err
	}

	return view.Execute(w, data, nil)
}
