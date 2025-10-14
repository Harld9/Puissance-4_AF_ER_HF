package game

import (
	"log"
	"net/http"
	"strconv"
)

type Grille struct {
	Tableau [8][9]int
}

func InitGame() *Grille {

	return &Grille{

		// définition grille du p4
		// pour chercher dans la grille le premier grille[] sert à définir la ligne donc grille [0] égal premiere ligne hoizontale
		// et donc le deuxième grille[][0] la première colonne sachant quand on cherche on part de tout en haut à gauce
		//  - Emrick:
		// En gros grille[i][j] = un tableau de tableaux; i = ligne 1 et j = colonne 1
		Tableau: [8][9]int{
			{3, 3, 3, 3, 3, 3, 3, 3, 3}, // mur du haut
			{3, 0, 0, 0, 0, 0, 0, 0, 3},
			{3, 0, 0, 0, 0, 0, 0, 0, 3},
			{3, 0, 0, 0, 0, 0, 0, 0, 3},
			{3, 0, 0, 0, 0, 0, 0, 0, 3},
			{3, 0, 0, 0, 0, 0, 0, 0, 3},
			{3, 0, 0, 0, 0, 0, 0, 0, 3},
			{3, 3, 3, 3, 3, 3, 3, 3, 3}, // mur du bas
		},
	}
}

func Tour_joueur(grille *Grille, r *http.Request) {
	colStr := r.FormValue("colonne")
	col, err := strconv.Atoi(colStr)
	if err != nil {
		log.Println("Erreur conversion colonne :", err)
	}

	for ligne := 6; ligne >= 1; ligne-- {
		if grille.Tableau[ligne][col] == 0 {
			grille.Tableau[ligne][col]++
			/* Ajouter dans structure position dernier pion joué:
			Position[0] == ligne
			Position[1] == col */
			break
		}
	}
}

/*func IsWinPlayer1() {
	nb_colonnes := 7
	nb_lignes := 6

	// Horizontal
	compteurHorizontal := 0
	for i := -3; i <= 3; i++ {
		if x+i >= 0 && x+i < nb_colonnes { // reste dans la grille
			if plateau[y][x+i] == 1 { // pion du joueur
				compteurHorizontal++
			} else {
				compteurHorizontal = 0
			}
			if compteurHorizontal >= 4 {
				fmt.Print("Victoire horizontal")
			}
		}
	}

	// Vertical
	compteurVertical := 0
	for i := -3; i <= 3; i++ {
		if y+i >= 0 && y+i < nb_lignes {
			if plateau[y+i][x] == 1 {
				compteurVertical++
			} else {
				compteurVertical = 0
			}
			if compteurVertical >= 4 {
				fmt.Print("Victoire vertical")
			}
		}
	}

	// Diagonale /
	compteurDiag1 := 0
	for i := -3; i <= 3; i++ {
		if x+i >= 0 && x+i < nb_colonnes && y-i >= 0 && y-i < nb_lignes {
			if plateau[y-i][x+i] == 1 {
				compteurDiag1++
			} else {
				compteurDiag1 = 0
			}
			if compteurDiag1 >= 4 {
				fmt.Print("Victoire diagonale en haut/bas droit")
			}
		}
	}

	// Diagonale \
	compteurDiag2 := 0
	for i := -3; i <= 3; i++ {
		if x+i >= 0 && x+i < nb_colonnes && y+i >= 0 && y+i < nb_lignes {
			if plateau[y+i][x+i] == 1 {
				compteurDiag2++
			} else {
				compteurDiag2 = 0
			}
			if compteurDiag2 >= 4 {
				fmt.Print("Victoire diagonale en haut/bas gauche")
			}
		}
	}
} */
