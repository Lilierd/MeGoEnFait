package ville

/**
* Structure d'une ville
 */

type Ville struct {
	Nom      string
	Visitee  bool
	Distance int
}

/**
* Constructeur de la struct Ville
* @param Nom : nom de la ville
*/
func New(Nom string) Ville {
	ville := Ville{Nom, false, 999999}
	return ville
}
