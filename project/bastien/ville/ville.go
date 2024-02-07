package ville

/**
* Structure d'une ville
 */

type Ville struct {
	Nom     string
	Visitee bool
}

/**
* Constructeur de la struct Ville
* @param Nom : nom de la ville
 */
func New(Nom string) Ville {
	ville := Ville{Nom, false}
	return ville
}

/**
* Méthode de classe
 */
func (v Ville) IsVisited() bool {
	return v.Visitee
}

func (v *Ville) SetVisitee(isVisited bool) {
	v.Visitee = isVisited
}
