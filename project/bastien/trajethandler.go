package bastien

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"project/bastien/trajet"
	"project/bastien/ville"
	"strconv"
	"strings"
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
			if(v.Nom == vi.Nom) {
				src = false
				break;
			}
		}
		if src {
			already = append(already, v)
		}
	}

	return already
}