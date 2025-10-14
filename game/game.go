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
