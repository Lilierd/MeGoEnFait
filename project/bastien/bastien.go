package bastien

import (
	"fmt"
	ville "project/bastien/ville"
)

func Bastien() {

}

/**
* Exemple d'instanciation d'une ville
*/
func InstanceVille(){
	reims := ville.New("Reims")

	fmt.Println(reims.Nom)
}