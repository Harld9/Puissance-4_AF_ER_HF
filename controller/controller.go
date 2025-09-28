package controller

import (
	"html/template"
	"net/http"
)

// renderTemplate est une fonction utilitaire pour afficher un template HTML avec des données dynamiques
func renderTemplate(w http.ResponseWriter, filename string, data map[string]string) {
	tmpl := template.Must(template.ParseFiles("template/" + filename)) // Charge le fichier template depuis le dossier "template"
	tmpl.Execute(w, data)                                              // Exécute le template et écrit le résultat dans la réponse HTTP
}

// Home gère la page d'accueil
func Home(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"Title":   "Accueil",                       // Titre de la page
		"Message": "Bienvenue chez la PUISSANCE 🎉", // Message affiché dans le template
	}
	renderTemplate(w, "index.html", data) // Affiche le template index.html avec les données
}

// About gère la page "À propos"
func About(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"Title":   "À propos",
		"Message": "Ceci est la page À propos ✨",
	}
	renderTemplate(w, "about.html", data) // Affiche le template about.html avec les données
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
	if r.Method == http.MethodPost { // Si le formulaire est soumis en POST
		// Récupération des données du formulaire
		joueur1 := r.FormValue("joueur1") // Récupère le champ "joueur1"
		joueur2 := r.FormValue("joueur2") // Récupère le champ "joueur2"
		data := map[string]string{
			"Title":   "Jeu en cours",
			"Message": "Joueurs : " + joueur1 + " vs " + joueur2,
		}
		renderTemplate(w, "jeu.html", data)
		return // On termine ici pour ne pas exécuter la partie GET
	}
	data := map[string]string{
		"Title":   "Entrée des joueurs",
		"Message": "Saisissez les noms des Joueurs",
	}
	renderTemplate(w, "jeu.html", data)

	/*func IsWin(){
	  	zqz
	  	compteur := 0
	  	for {
	  		//Check à droite
	  		for i := 0; i < 4; i++ {
	  			if pion at x+i y == 1 {
	  				compteur++
	  			} else {
	  			break
	  			}
	  		}
	  		//Check à gauche
	  		for i := 0; i < 4; i++ {
	  			if pion at x-i y == 1 {
	  				compteur++
	  			} else {
	  			break
	  			}
	  		}
	  		//Check en haut
	  		for i := 0; i < 4; i++ {
	      		if pion at x y+i == 1 {
	      			compteur++
	      		} else {
	     	    	break
	      		}
	  		}
	  		//Check en bas
	  		for i := 0; i < 4; i++ {
	   		   if pion at x y-i == 1 {
	    		    	compteur++
	  		    } else {
	  	        break
	  		    }
	  		}
	      	//check en diagonalesupérieur droite
	  		for i := 0; i < 4; i++ {
	  		    if pion at x+i y+i == 1 {
	  				compteur++
	  	    	} else {
	  	        break
	  	    	}
	  		}
	      	//check en diagonale supérieur gauche
	  		for i := 0; i < 4; i++ {
	  		    if pion at x-i y+i == 1 {
	  	 	       compteur++
	  		    } else {
	  	        break
	  		    }
	  		}
	  		//check en diagonale inferieure gauche
	  		for i := 0; i < 4; i++ {
	  		    if pion at x-i y-i == 1 {
	  		        compteur++
	  		    } else {
	  		    break
	  		    }
	  		}
	  		//check en diagonale inferieure droite
	  		for i := 0; i < 4; i++ {
	  		    if pion at x+i y-i == 1 {
	  		        compteur++
	  		    } else {
	  		    break
	  		    }
	  		}
	  		if compteur = 4 {
	  			return true
	  		}
	  	}
	  }*/
}
