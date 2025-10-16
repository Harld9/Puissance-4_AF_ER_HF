package game

import (
	"log"
	"net/http"
	"strconv"
)

type Position struct {
	Ligne int
	Col   int
}

type GameData struct {
	J1         string
	J2         string
	Tableau    [8][9]int
	Position   [1]Position
	Debut      bool
	NbTour     int
	TourJoueur int
	Winnner    string
	GameEnd    bool
}

func InitGame() *GameData {
	return &GameData{
		J1: "",
		J2: "",
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
		Position: [1]Position{
			{Ligne: 0, Col: 0},
		},
		Debut:      false,
		NbTour:     1,
		TourJoueur: 1, //1 == J1; 2 == J2
		Winnner:    "",
		GameEnd:    false,
	}
}

func Nomdesjoueurs(g *GameData) string {
	if g.TourJoueur == 1 {
		return g.J1 + " (rouge) 🔴"
	} else {
		return g.J2 + " (jaune) 🟡"
	}
}

func Tour_joueur(g *GameData, r *http.Request) {

	if !g.Debut {
		log.Println("La partie n'est pas en cours.")
		return
	}

	colStr := r.FormValue("colonne")
	col, err := strconv.Atoi(colStr)
	if err != nil {
		log.Println("Erreur conversion colonne :", err)
		return
	}

	// Vérifie si la colonne reçue est entre 1 et 7
	if col < 1 || col > 7 {
		log.Println("Colonne hors borne :", col)
		return
	}

	for ligne := 6; ligne >= 1; ligne-- {
		if g.Tableau[ligne][col] == 0 {
			g.Tableau[ligne][col] = g.TourJoueur
			g.Position[0].Col = col
			g.Position[0].Ligne = ligne
			g.NbTour++
			break
		}
	}

	if WinCheck(g, g.TourJoueur) { // Check si le joueur a gagné
		var player string
		if g.TourJoueur == 1 {
			player = g.J1
		} else {
			player = g.J2
		}
		g.Winnner = "Victoire de " + player
		log.Printf("Le joueur %s gagne", player)
		g.GameEnd = true
		return
	}

	if g.NbTour == 43 { //Verif match nul si tour == 43
		g.Winnner = "Match nul"
		log.Println("Match nul")
		g.GameEnd = true
		return
	}

	// Alterne le joueur
	if g.TourJoueur == 1 {
		g.TourJoueur = 2
	} else {
		g.TourJoueur = 1
	}
}

func WinCheck(g *GameData, player int) bool {

	nb_colonnes_min := 1
	nb_colonnes_max := 7
	nb_lignes_min := 1
	nb_lignes_max := 6

	x := g.Position[0].Col
	y := g.Position[0].Ligne

	// Horizontal
	compteurHorizontal := 0
	for i := -3; i <= 3; i++ {
		col_act := x + i
		if col_act >= nb_colonnes_min && col_act <= nb_colonnes_max { // reste dans la grille
			if g.Tableau[y][col_act] == player { // pion du joueur
				compteurHorizontal++
			} else {
				compteurHorizontal = 0
			}
			if compteurHorizontal >= 4 {
				return true
			}
		}
	}

	// Vertical
	compteurVertical := 0
	for i := -3; i <= 3; i++ {
		ligne_act := y + i
		if ligne_act >= nb_lignes_min && ligne_act <= nb_lignes_max {
			if g.Tableau[ligne_act][x] == player {
				compteurVertical++
			} else {
				compteurVertical = 0
			}
			if compteurVertical >= 4 {
				return true
			}
		}
	}

	// Diagonale /
	compteurDiag1 := 0
	for i := -3; i <= 3; i++ {
		col_act := x + i
		ligne_act := y - i
		if col_act >= nb_colonnes_min && col_act <= nb_colonnes_max && ligne_act >= nb_lignes_min && ligne_act <= nb_lignes_max {
			if g.Tableau[ligne_act][col_act] == player {
				compteurDiag1++
			} else {
				compteurDiag1 = 0
			}
			if compteurDiag1 >= 4 {
				return true
			}
		}
	}

	// Diagonale \
	compteurDiag2 := 0
	for i := -3; i <= 3; i++ {
		col_act := x + i
		ligne_act := y + i
		if col_act >= nb_colonnes_min && col_act <= nb_colonnes_max && ligne_act >= nb_lignes_min && ligne_act <= nb_lignes_max {
			if g.Tableau[ligne_act][col_act] == player {
				compteurDiag2++
			} else {
				compteurDiag2 = 0
			}
			if compteurDiag2 >= 4 {
				return true
			}
		}
	}
	return false
}

func DrawCheck(g *GameData) bool {
	for ligne := 1; ligne <= 6; ligne++ {
		for colonne := 1; colonne <= 7; colonne++ {
			if g.Tableau[ligne][colonne] == 0 {
				return false
			}
		}
	}
	return true
}

/*func History(){






}

func Leaderboard() int{



}*/
