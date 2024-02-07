package bastien

import (
	"fmt"
	"strconv"

	ville "github.com/Lilierd/MeGoEnFait/project/bastien/ville"
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
