package router

import (
	"net/http"
	"power4/controller"
)

// New crée et retourne un nouvel objet ServeMux configuré avec les routes de l'application
func New() *http.ServeMux {
	mux := http.NewServeMux()

	// Routes de ton app
	mux.HandleFunc("/", controller.Home)
	mux.HandleFunc("/about", controller.About)
	mux.HandleFunc("/contact", controller.Contact)
	mux.HandleFunc("/jeu", controller.Jeu)

	// Ajout des fichiers statiques
	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	return mux
}
