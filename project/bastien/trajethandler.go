package bastien

import (
	"fmt"
	"project/bastien/trajet"
	"project/bastien/ville"
)

func Dijsktra(nomVille1 string, nomVille2 string) {
	//depart := ville.New()
	
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

func GetVilles(tab []trajet.Trajet) []ville.Ville {
	var already []ville.Ville

	for _, traj := range tab {
		var vi ville.Ville
		src := true
		for _, v := range already {
			if traj.GetVille1().Nom == v.Nom || traj.GetVille2().Nom == v.Nom {
				vi = v
				fmt.Println(vi.Nom)
				src = false
				break;
			}		
		}
		if src == true {
			already = append(already, vi)
		}
	}

	return already
}