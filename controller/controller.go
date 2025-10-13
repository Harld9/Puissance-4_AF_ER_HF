package controller

import (
	"html/template"
	"net/http"
	"power4/game"
)

// renderTemplate est une fonction utilitaire pour afficher un template HTML avec des données dynamiques
func renderTemplate(w http.ResponseWriter, filename string, data map[string]string) {
	tmpl := template.Must(template.ParseFiles("template/" + filename)) // Charge le fichier template depuis le dossier "template"
	tmpl.Execute(w, data)                                              // Exécute le template et écrit le résultat dans la réponse HTTP
}

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

// Contact gère la page de contact
func Contact(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost { // Si le formulaire est soumis en POST
		// Récupération des données du formulaire
		name := r.FormValue("name") // Récupère le champ "name"
		msg := r.FormValue("msg")   // Récupère le champ "msg"

		data := map[string]string{
			"Title":   "Contact",
			"Message": "Merci " + name + " pour ton message : " + msg, // Message personnalisé après soumission
		}
		renderTemplate(w, "contact.html", data)
		return // On termine ici pour ne pas exécuter la partie GET
	}

	// Si ce n'est pas un POST, on affiche simplement le formulaire
	data := map[string]string{
		"Title":   "Contact",
		"Message": "Envoie-nous un message 📩",
	}
	renderTemplate(w, "contact.html", data)
}

func Jeu(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		grille := game.Game()
		data := PageData{
			Title:   "Jeu en cours",
			Message: "C'est au tour de ",
			Tableau: grille.Tableau,
		}
		tmpl := template.Must(template.ParseFiles("template/jeu.html"))
		tmpl.Execute(w, data)
		return
	}

	data := PageData{
		Title:   "Entrée des joueurs",
		Message: "Saisissez les noms des Joueurs",
	}
	tmpl := template.Must(template.ParseFiles("template/jeu.html"))
	tmpl.Execute(w, data)
}
