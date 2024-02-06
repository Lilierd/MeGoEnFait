package bastien

import (
	"fmt"
	ville "project/bastien/ville"
	"strconv"
)

func Bastien() {

}

/**
* Exemple d'instanciation d'une ville
*/
func InstanceVille() ville.Ville {
	reims := ville.New("Reims")

	fmt.Println(reims.Nom)

	return reims
}

func parseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Error parsing int:", err)
		return 0
	}
	return i
}
