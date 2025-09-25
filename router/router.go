package router

import (
	"net/http"
	"power4/controller"
)

// New crée et retourne un nouvel objet ServeMux configuré avec les routes de l'app
func New() *http.ServeMux { // 2 usages

	mux := http.NewServeMux() // Création d'un nouveau ServeMux, qui est un routeur

	// On associe les chemins URL à des fonctions spécifiques du controller
	mux.HandleFunc("/", controller.Home)           // "/" correspond à la page d'accueil
	mux.HandleFunc("/about", controller.About)     // "/about" correspond à la page "À propos"
	mux.HandleFunc("/contact", controller.Contact) // "/contact" correspond à la page contact

	return mux // On retourne le routeur configuré
}
