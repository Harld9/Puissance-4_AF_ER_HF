package controller

import (
	"encoding/json"
	"html/template"
	"net/http"
	"os"
	"power4/game"
)

// renderTemplate est une fonction utilitaire pour afficher un template HTML avec des données dynamiques
func renderTemplate(w http.ResponseWriter, filename string, data map[string]string) {
	tmpl := template.Must(template.ParseFiles("template/" + filename)) // Charge le fichier template depuis le dossier "template"
	tmpl.Execute(w, data)                                              // Exécute le template et écrit le résultat dans la réponse HTTP
}

type PageData struct {
	Title         string
	Message       string
	Tableau       [8][9]int
	Player1       string
	Player2       string
	NbTour        int
	EnCours       bool
	JoueurCourant string
	GameEnd       bool
	Victoires     []game.JoueurVictoire
}

func Home(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title:   "Accueil",
		Message: "Bienvenue Au Jeu Puissance 4 🎉",
	}
	tmpl := template.Must(template.ParseFiles("template/index.html"))
	tmpl.Execute(w, data)
}

func Leaderboard(w http.ResponseWriter, r *http.Request) {
	var joueurs []game.JoueurVictoire

	// on lit le fichier json des stats
	data, err := os.ReadFile(game.Path)
	if err != nil {
		http.Error(w, "Impossible de lire les statistiques", http.StatusInternalServerError)
		return
	}
	// Si fichier on des datas on met les données du json dans la structure joueurs
	if len(data) > 0 {
		if err := json.Unmarshal(data, &joueurs); err != nil {
			http.Error(w, "Impossible de parser les statistiques", http.StatusInternalServerError) //Chat gpt m'a donné ça car on la fonction ne peut pas retourner erreur car pas dans son type de retour
			return
		}
	}

	// On trie les joueurs par nombre de victoires décroissant
	for i := 0; i < len(joueurs); i++ {
		for j := i + 1; j < len(joueurs); j++ {
			if joueurs[j].Victoire > joueurs[i].Victoire {
				joueurs[i], joueurs[j] = joueurs[j], joueurs[i]
			}
		}
	}

	dataPage := PageData{
		Title:     "Leaderboard",
		Message:   "Classement des joueurs 🏆",
		Victoires: joueurs,
	}

	tmpl := template.Must(template.ParseFiles("template/leaderboard.html"))
	tmpl.Execute(w, dataPage)
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

var G *game.GameData = game.InitGame()

func Jeu(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		if r.FormValue("reset") == "1" {
			G = game.InitGame()
			G.Debut = false
			http.Redirect(w, r, "/jeu", http.StatusSeeOther) // Redirection après POST, un return
			return
		}
		player1 := r.FormValue("player1")
		player2 := r.FormValue("player2")
		if player1 != "" && player2 != "" {
			G.J1 = player1
			G.J2 = player2
			G.Debut = true
		} else {
			game.Tour_joueur(G, r)
			if G.Winnner != "" {
				data := PageData{
					Title:   "Fin de partie",
					Message: G.Winnner,
					Tableau: G.Tableau,
					EnCours: G.Debut,
					GameEnd: G.GameEnd,
				}
				tmpl := template.Must(template.ParseFiles("template/jeu.html"))
				tmpl.Execute(w, data)
			}
		}
		http.Redirect(w, r, "/jeu", http.StatusSeeOther) // Redirection après POST, un return

		return
	}

	var title, message string
	if !G.Debut {
		title = "Bienvenue sur le Puissance 4"
		message = "Entrez les noms des joueurs pour commencer la partie"
	} else if G.Winnner != "" {
		title = "Fin de partie"
		message = G.Winnner
	} else {
		title = "Partie en cours !"
	}

	data := PageData{
		Title:         title,
		Message:       message,
		Tableau:       G.Tableau,
		Player1:       G.J1,
		Player2:       G.J2,
		NbTour:        G.NbTour,
		EnCours:       G.Debut,
		JoueurCourant: game.Nomdesjoueurs(G),
		GameEnd:       G.GameEnd,
	}
	tmpl := template.Must(template.ParseFiles("template/jeu.html"))
	tmpl.Execute(w, data)
}
