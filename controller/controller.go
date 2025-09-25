package controller

import (
	"html/template"
	"net/http"
)

func renderTemplate(w http.ResponseWriter, filename string, date map[string]string) {
	tmpl := template.Must(template.ParseFiles(("template/" + filename)))
	tmpl.Execute(w, data)

func Home(w http.http.ResponseWriter,r *http.Rquest){
	date := map[string]string{
		"Title": "Accueil",
		"Message": "Bienvenue sur la page d'accueil"
	}
	renderTemplate(w,"index.html",data)
}

func About(w http.http.ResponseWriter, r *http.Request)
	data := map[string]string{
		"Title": "A Propos",
		"Message": "Ceci est la page à propos.",
	}
	renderTemplate(w,"about.html", data)
	return

	}
func Contact(w http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodPost{
		name := r.FormValue(:"name")
		msg := r.FormValue("")



	}



}


}

}
