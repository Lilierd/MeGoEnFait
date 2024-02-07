package bastien

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/Lilierd/MeGoEnFait/project/bastien/trajet"
	"github.com/Lilierd/MeGoEnFait/project/bastien/ville"
)

func Dijsktra(nomVille1 string, nomVille2 string) {
	//Tous les trajets !
	trajets := GetTrajets()
	//Toutes les villes !
	villes := GetVilles(trajets)
	//On enregistre les distances entre les villes
	distances := make(map[string]int)
	//On enregistre les villes visitées
	visited := make(map[string]bool)

	// Initialiser les distances pour toutes les villes à l'infini, sauf pour la ville de départ à 0
	for _, v := range villes {
		distances[v.Nom] = 1000000000
	}
	distances[nomVille1] = 0

	// Parcourir les villes pour trouver le plus court chemin
	for len(visited) < len(villes) {
		minVille := ""
		minDistance := 1000000000

		// Trouver la ville non visitée la plus proche
		for _, v := range villes {
			if !visited[v.Nom] && distances[v.Nom] < minDistance {
				minVille = v.Nom
				minDistance = distances[v.Nom]
			}
		}

		// Si aucune ville non visitée n'a été trouvée, cela signifie que toutes les villes ont été visitées
		if minVille == "" {
			break
		}

		// Marquer la ville comme visitée
		visited[minVille] = true

		// Mettre à jour les distances pour les villes voisines non visitées
		for _, traj := range GetTrajetsFromVille(trajets, ville.New(minVille)) {
			autreVille := ""
			if traj.GetVille1().Nom == minVille {
				autreVille = traj.GetVille2().Nom
			} else {
				autreVille = traj.GetVille1().Nom
			}

			// Calculer la nouvelle distance
			nouvelleDistance := distances[minVille] + traj.GetDistance()

			// Mettre à jour la distance si elle est plus petite
			if nouvelleDistance < distances[autreVille] {
				distances[autreVille] = nouvelleDistance
			}
		}
	}
	for visited := range visited {
		fmt.Println(visited)
	}
	// Afficher la distance entre les deux villes si la ville d'arrivée a été visitée
	if visited[nomVille2] {
		fmt.Println("La distance entre", nomVille1, "et", nomVille2, "est de", distances[nomVille2])
	} else {
		fmt.Println("Aucun chemin trouvé entre", nomVille1, "et", nomVille2)
	}
}

func GetNbVilles(tab []trajet.Trajet) int {
	i := 0
	already := GetVilles(tab)

	for _, v := range already {
		v.IsVisited()
		i++
	}

	return i
}

func GetTrajets() []trajet.Trajet {
	f, err := os.Open("carte.txt")

	var tab []trajet.Trajet

	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		x, errs := strconv.Atoi(line[2])
		if errs != nil {
			// ... handle error
			panic(err)
		}
		v1 := ville.New(line[0])
		v2 := ville.New(line[1])
		tab = append(tab, trajet.New(v1, v2, x))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return tab
}

func GetVilles(tab []trajet.Trajet) []ville.Ville {
	var already []ville.Ville
	var alreadybis []ville.Ville

	for _, traj := range tab {
		vi := traj.GetVille1()
		vii := traj.GetVille2()

		alreadybis = append(alreadybis, vi)
		alreadybis = append(alreadybis, vii)
	}

	for _, v := range alreadybis {
		src := true
		for _, vi := range already {
			if v.Nom == vi.Nom {
				src = false
				break
			}
		}
		if src {
			already = append(already, v)
		}
	}

	return already
}

func GetTrajetsFromVille(tab []trajet.Trajet, ville ville.Ville) []trajet.Trajet {
	var already []trajet.Trajet

	for _, traj := range tab {
		if traj.GetVille1().Nom == ville.Nom || traj.GetVille2().Nom == ville.Nom {
			already = append(already, traj)
		}
	}

	return already
}
