package main

import (
	"fmt"
	"net/http"
)

func main() {
	r := router.New()

	fmt.Println("Serveur démarré sur http:/localhost:8080")
	http.ListenAndServe(":8080", r)

}
