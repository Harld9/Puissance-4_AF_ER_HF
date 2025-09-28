package game

import "fmt"

func game() {
	// définition grille du p4
	// pour chercher dans la grille le premier grille[] sert à définir la ligne donc grille [0] égal premiere ligne hoizontale
	// et donc le deuxième grille[][0] la première colonne sachant quand on cherche on part de tout en haut à gauce
	//  - Emrick:
	// En gros grille[i][j] = un tableau de tableaux; i = ligne 1 et j = colonne 1
	grille := [][]int{
		{3, 3, 3, 3, 3, 3, 3, 3, 3}, //Ligne 1
		{3, 0, 0, 0, 0, 0, 0, 0, 3}, //Ligne 2
		{3, 0, 0, 0, 0, 0, 0, 0, 3}, //Ligne 3
		{3, 0, 0, 0, 0, 0, 0, 0, 3}, //Ligne 4
		{3, 0, 0, 0, 0, 0, 0, 0, 3}, //Ligne 5
		{3, 0, 0, 0, 0, 0, 0, 0, 3}, //Ligne 6
		{3, 0, 0, 0, 0, 0, 0, 0, 3}, //Ligne 7
		{3, 3, 3, 3, 3, 3, 3, 3, 3}, //Ligne 8
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
