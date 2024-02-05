package bastien

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"project/bastien/trajet"
	ville "project/bastien/ville"
	"strconv"
	"strings"
)

func Bastien() {

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
		tab = append(tab, trajet.New(ville.New(line[0]), ville.New(line[1]), x))
	}

	if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

	return tab
}

/**
* Exemple d'instanciation d'une ville
*/
func InstanceVille() ville.Ville {
	reims := ville.New("Reims")

	fmt.Println(reims.Nom)

	return reims
}