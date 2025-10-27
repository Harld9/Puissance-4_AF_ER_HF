package game

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type JoueurVictoire struct {
	Nom      string `json:"nom"`
	Victoire int    `json:"victoire"`
}

type Position struct {
	Ligne int
	Col   int
}

type GameData struct {
	J1            string
	J2            string
	Tableau       [8][9]int
	Position      [1]Position
	Debut         bool
	NbTour        int
	TourJoueur    int
	Winnner       string
	GameEnd       bool
	Encouragement string
}

const Path = "data/stats.json"

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
	Phrasealeatoire(g)
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
		g.Winnner = "Victoire de " + player + " en " + strconv.Itoa(g.NbTour-1) + " tours !"
		log.Printf("Le joueur %s gagne", player)
		g.GameEnd = true
		if err := WinLeaderboard(player); err != nil {
			log.Println("Erreur leaderboard:", err)
		}
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

	var piongagnant1 [2]int
	var piongagnant2 [2]int
	var piongagnant3 [2]int
	var piongagnant4 [2]int

	// Horizontal
	compteurHorizontal := 0
	for i := -3; i <= 3; i++ {
		col_act := x + i
		if col_act >= nb_colonnes_min && col_act <= nb_colonnes_max { // reste dans la grille
			if g.Tableau[y][col_act] == player { // pion du joueur
				switch compteurHorizontal {
				case 0:
					piongagnant1[0] = col_act
					piongagnant1[1] = y
				case 1:
					piongagnant2[0] = col_act
					piongagnant2[1] = y
				case 2:
					piongagnant3[0] = col_act
					piongagnant3[1] = y
				case 3:
					piongagnant4[0] = col_act
					piongagnant4[1] = y
				}
				compteurHorizontal++
			} else {
				compteurHorizontal = 0
			}
			if compteurHorizontal >= 4 {
				// Marquer les pions gagnants
				g.Tableau[piongagnant1[1]][piongagnant1[0]] = player + 2
				g.Tableau[piongagnant2[1]][piongagnant2[0]] = player + 2
				g.Tableau[piongagnant3[1]][piongagnant3[0]] = player + 2
				g.Tableau[piongagnant4[1]][piongagnant4[0]] = player + 2
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
				switch compteurVertical {
				case 0:
					piongagnant1[0] = x
					piongagnant1[1] = ligne_act
				case 1:
					piongagnant2[0] = x
					piongagnant2[1] = ligne_act
				case 2:
					piongagnant3[0] = x
					piongagnant3[1] = ligne_act
				case 3:
					piongagnant4[0] = x
					piongagnant4[1] = ligne_act
				}
				compteurVertical++
			} else {
				compteurVertical = 0
			}
			if compteurVertical >= 4 {
				g.Tableau[piongagnant1[1]][piongagnant1[0]] = player + 2
				g.Tableau[piongagnant2[1]][piongagnant2[0]] = player + 2
				g.Tableau[piongagnant3[1]][piongagnant3[0]] = player + 2
				g.Tableau[piongagnant4[1]][piongagnant4[0]] = player + 2
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
				switch compteurDiag1 {
				case 0:
					piongagnant1[0] = col_act
					piongagnant1[1] = ligne_act
				case 1:
					piongagnant2[0] = col_act
					piongagnant2[1] = ligne_act
				case 2:
					piongagnant3[0] = col_act
					piongagnant3[1] = ligne_act
				case 3:
					piongagnant4[0] = col_act
					piongagnant4[1] = ligne_act
				}
				compteurDiag1++
			} else {
				compteurDiag1 = 0
			}
			if compteurDiag1 >= 4 {
				g.Tableau[piongagnant1[1]][piongagnant1[0]] = player + 2
				g.Tableau[piongagnant2[1]][piongagnant2[0]] = player + 2
				g.Tableau[piongagnant3[1]][piongagnant3[0]] = player + 2
				g.Tableau[piongagnant4[1]][piongagnant4[0]] = player + 2
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
				switch compteurDiag2 {
				case 0:
					piongagnant1[0] = col_act
					piongagnant1[1] = ligne_act
				case 1:
					piongagnant2[0] = col_act
					piongagnant2[1] = ligne_act
				case 2:
					piongagnant3[0] = col_act
					piongagnant3[1] = ligne_act
				case 3:
					piongagnant4[0] = col_act
					piongagnant4[1] = ligne_act
				}
				compteurDiag2++
			} else {
				compteurDiag2 = 0
			}
			if compteurDiag2 >= 4 {
				g.Tableau[piongagnant1[1]][piongagnant1[0]] = player + 2
				g.Tableau[piongagnant2[1]][piongagnant2[0]] = player + 2
				g.Tableau[piongagnant3[1]][piongagnant3[0]] = player + 2
				g.Tableau[piongagnant4[1]][piongagnant4[0]] = player + 2
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

func WinLeaderboard(nomGagnant string) error {
	//déclare une slice de joueur pour stocker les scores
	var joueurs []JoueurVictoire
	//chemin du fichier json

	//lecture du fichier json
	data, err := os.ReadFile(Path)
	if err != nil { //si erreur lors de la lecture
		return err //retourne erreur
	}

	//si le fichier a des data on essaye de parser en slice de joueur
	if len(data) > 0 {
		err := json.Unmarshal(data, &joueurs)
		if err != nil { //decode json
			return err //erreur si invalide
		}
	}

	//variable pour savoir si joueur existe deja
	existe := false

	//parcourt tous les joueurs pour voir si nomGagnant est deja present
	for i := 0; i < len(joueurs); i++ {
		if strings.EqualFold(joueurs[i].Nom, nomGagnant) { //comparaison insensible a la casse
			joueurs[i].Victoire++ //increment victoire
			existe = true         //joueur existe
			break                 //sort de la boucle
		}
	}

	//si joueur existe pas on l'ajoute
	if !existe {
		joueurs = append(joueurs, JoueurVictoire{Nom: nomGagnant, Victoire: 1})
	}

	//encode la slice en json avec indentation
	data, err = json.MarshalIndent(joueurs, "", "  ")
	if err != nil { //si erreur encode
		return err
	}

	//ecrit les datas dans le fichier
	if err := os.WriteFile(Path, data, 0644); err != nil { //0644 permissions fichier
		return err //retourne erreur si echec
	}

	//tout est ok
	return nil
}
func Phrasealeatoire(g *GameData) {

	phrase := []string{
		"T'as presque gagné… enfin, si on compte à l'envers.",
		"La logique est là, mais elle est encore en train de charger.",
		"Si Einstein voyait ça, il refermerait son tableau.",
		"T'es à deux coups de la gloire… ou du désastre.",
		"On sent le génie, mais il est un peu en RTT.",
		"Ta stratégie, c'est de la poésie incomprise.",
		"L'adversaire tremble… de rire ou de peur, on sait pas encore.",
		"C'est pas perdu tant que t'as pas crié “PUISSANCE 3”.",
		"T'as un sixième sens, mais il est en panne.",
		"Le cerveau chauffe, mais le résultat reste tiède.",
		"On dirait un grand maître d'échec… mais dans un autre jeu.",
		"T'as mis la puissance, il manque juste le 4.",
		"La logique est une chose, le hasard en est une autre. T'as choisi la troisième.",
		"Continue comme ça, la victoire est au coin du plateau.",
		"On sent la concentration… et un peu la panique aussi.",
		"T'es pas loin de gagner, juste à un océan de cases près.",
		"Chaque jeton posé, c'est une étincelle d'espoir… ou de désespoir.",
		"Tactique de ninja : lente, silencieuse… et parfois complètement à côté.",
		"La chance te sourit ! Mais elle louche un peu.",
		"C'est pas une erreur, c'est une feinte de génie incomprise.",
		"T'as posé un jeton comme un champion, dommage que ce soit le mauvais endroit.",
		"Tu joues avec le feu, mais sans extincteur.",
		"On sent la réflexion intense… ou le vide absolu.",
		"La patience est une vertu, mais là c'est long.",
		"Chaque coup est un mystère pour l'humanité.",
		"Tu joues à Puissance 4 comme si c'était un Sudoku.",
		"Ton cerveau fait du sport, mais sans échauffement.",
		"Le plateau n'était pas prêt pour autant de créativité.",
		"T'as confondu stratégie et freestyle.",
		"La tension est palpable, surtout chez l'adversaire.",
		"Ce n'est pas du hasard, c'est du talent mal orienté.",
		"Si tu gagnes, ce sera historique. Si tu perds, ce sera logique.",
		"La victoire te fait coucou, mais de très loin.",
		"Tu vis dangereusement, jeton après jeton.",
		"C'est pas un move, c'est une déclaration d'intention.",
		"On voit la stratégie… faut juste un microscope.",
		"Tu pourrais enseigner l'art du presque.",
		"T'es à deux doigts d'un coup de génie, mais c'est pas ceux-là.",
		"L'adversaire doute, mais toi aussi, apparemment.",
		"Ton cerveau demande une pause café.",
		"Chaque jeton posé te rapproche du mystère.",
		"Tu joues comme un philosophe, pas comme un stratège.",
		"Si le but c'était de faire joli, t'as gagné.",
		"Le plateau pleure, mais il t'encourage.",
		"T'es dans la zone… juste pas la bonne.",
		"Tu fais vibrer la logique et pleurer les mathématiques.",
		"Un coup comme ça, on en voit qu'en expérimental.",
		"Tu vises la victoire avec un GPS cassé.",
		"Le suspense est à son comble, surtout pour ton prochain move.",
		"On dirait que tu joues contre le destin.",
		"C'est du grand art, dans une autre dimension.",
		"T'es pas dans le game, t'es au-dessus du game.",
		"Le plateau se souviendra de toi… pas forcément bien.",
		"Chaque tour, c'est une aventure émotionnelle.",
		"Tu donnes tout, même les erreurs.",
		"Le génie frappe parfois, mais il t'a raté de peu.",
		"Ton adversaire transpire, et toi… tu inspires la confusion.",
		"T'as voulu faire un piège, t'as fini par te piéger.",
		"Les dieux du Puissance 4 observent, perplexes.",
		"T'as la logique d'un poète en pleine inspiration.",
		"Un jour, les historiens parleront de ce coup.",
		"Le chaos a trouvé son maître.",
		"T'as tenté un move risqué, et le risque a gagné.",
		"Tu écris la légende du “presque gagnant”.",
		"C'est pas du hasard, c'est du destin mal coordonné.",
		"Ton style est unique, personne n'aurait osé ça.",
		"L'adversaire ne comprend plus rien, et toi non plus.",
		"T'as trouvé un nouveau concept : la défaite artistique.",
		"On sent que t'y crois, c'est déjà la moitié du chemin.",
		"Le plateau te regarde avec compassion.",
		"T'as pas perdu, t'as juste offert un spectacle.",
		"Le cerveau tourne à fond, mais pas dans le bon sens.",
		"Tu poses des jetons comme Picasso peignait : sans plan précis.",
		"T'es un pionnier du move imprévisible.",
		"Le suspense est réel, même si la logique l'est moins.",
		"Chaque case vide est une promesse… souvent brisée.",
		"Tu joues avec ton cœur, pas avec ta tête.",
		"On reconnaît les grands joueurs à leurs grandes erreurs.",
		"T'as confondu Puissance 4 et roulette russe.",
		"Le jeu est serré, mais ta logique aussi.",
		"T'as la vision, mais sans lunettes.",
		"Un coup digne des plus grands… comiques.",
		"On sent que t'étais proche d'avoir raison.",
		"Tu fais douter les lois du jeu.",
		"Ton move a surpris tout le monde, même toi.",
		"T'as créé une nouvelle stratégie : l'improvisation totale.",
		"L'adversaire panique, mais pour d'autres raisons.",
		"Ton jeton vient d'atterrir dans la légende… ou la honte.",
		"Tu joues comme un artiste abstrait.",
		"T'es un danger public pour les colonnes.",
		"Chaque tour est un poème d'incertitude.",
		"T'as la patience d'un moine et la logique d'un poulpe.",
		"La victoire approche, mais elle a pris un détour.",
		"T'as misé sur le chaos, et le chaos t'a remercié.",
		"Ce coup… personne n'aurait osé.",
		"On voit la stratégie, mais faut la deviner.",
		"Tu fais rêver ton adversaire, surtout de gagner.",
		"Le plateau est témoin d'un concept révolutionnaire : la confusion tactique.",
		"T'as presque inventé un nouveau jeu.",
		"Ton jeton a fait un beau voyage, c'est déjà ça.",
		"T'as la passion, il te manque juste la précision.",
		"Chaque coup est une surprise, même pour toi.",
		"Tu fais du Puissance 4 version freestyle.",
		"Le silence du plateau en dit long.",
		"T'es pas loin de la victoire… géographiquement.",
		"Ton move restera dans les mémoires, mais pas pour les bonnes raisons.",
		"Le destin te teste, et il s'amuse bien.",
		"Tu joues comme si chaque coup était le dernier — et parfois, il l'est.",
		"T'as prouvé que la logique, c'est surfait.",
		"Ce n'est pas une défaite, c'est une expérience sociale.",
		"On ne gagne pas toujours, mais toi, tu marques les esprits.",
	}
	r := rand.Intn(100)
	phrasealeatoire := phrase[r]
	g.Encouragement = phrasealeatoire

}
