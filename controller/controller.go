package controller

import (
	"html/template"
	"net/http"
	"power4/game"
)

type PageData struct {
	Title   string
	Message string
	Tableau [8][9]int
}

func Home(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title:   "Accueil",
		Message: "Bienvenue chez la PUISSANCE 🎉",
	}
	tmpl := template.Must(template.ParseFiles("template/index.html"))
	tmpl.Execute(w, data)
}

func About(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title:   "À propos",
		Message: "Ceci est la page À propos ✨",
	}
	tmpl := template.Must(template.ParseFiles("template/about.html"))
	tmpl.Execute(w, data)
}

func Contact(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title:   "Contact",
		Message: "Envoie-nous un message 📩",
	}
	tmpl := template.Must(template.ParseFiles("template/contact.html"))
	tmpl.Execute(w, data)
}

func Jeu(w http.ResponseWriter, r *http.Request) {
	grille := game.Game()
	data := PageData{
		Title:   "Entrée des joueurs",
		Message: "Saisissez les noms des Joueurs",
		Tableau: grille.Tableau,
	}
	tmpl := template.Must(template.ParseFiles("template/jeu.html"))
	tmpl.Execute(w, data)
}
