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
	reims := ville.New("Reims", false, 20)

	fmt.Println(reims.Nom)
}