package game

import "fmt"

func game() {
	// défiinition grille du p4
	// pour chercher dans la grille le premier grille[] sert à définir la ligne donc grille [0] égal premiere ligne hoizontale
	// et donc le deuxième grille[][0] la première colonne sachant quand on cherche on part de tout en haut à gauce
	grille := [][]int{
		{3, 3, 3, 3, 3, 3, 3, 3, 3},
		{3, 0, 0, 0, 0, 0, 0, 0, 3},
		{3, 0, 0, 0, 0, 0, 0, 0, 3},
		{3, 0, 0, 0, 0, 0, 0, 0, 3},
		{3, 0, 0, 0, 0, 0, 0, 0, 3},
		{3, 0, 0, 0, 0, 0, 0, 0, 3},
		{3, 0, 0, 0, 0, 0, 0, 0, 3},
		{3, 3, 3, 3, 3, 3, 3, 3, 3},
	}

	//Yo les potes j'ai fait ça pour ajouter un jeton dans la premiere ligne de la grille mais je pense ça peut s'opti avec une boucle
	choix_joueur := 0
	fmt.Println("Choisis l'emplacement de ton jeton sur la première ligne 1 = tout en bas à gauche, 7 = tout en bas à droite")
	fmt.Scan(&choix_joueur)
	switch choix_joueur {
	case 1:
		grille[6][1] = (grille[6][1]) + 1
	case 2:
		grille[6][1] = (grille[6][1]) + 1
	case 3:
		grille[6][2] = (grille[6][1]) + 1
	case 4:
		grille[6][3] = (grille[6][1]) + 1
	case 5:
		grille[6][4] = (grille[6][1]) + 1
	case 6:
		grille[6][5] = (grille[6][1]) + 1
	case 7:
		grille[6][6] = (grille[6][1]) + 1
	}
	fmt.Println(grille)
	//rajouter uine fonction pour savoir si la colonne est pleine
	//et une fonction pour checker la win

}
