package trajet

import (
	ville "github.com/Lilierd/MeGoEnFait/project/bastien/ville"
)

type Trajet struct {
	ville1   ville.Ville
	ville2   ville.Ville
	distance int
}

/**
* Constructeur de la struct Ville
* @param ville1 : nom de la ville de départ
* @param ville2 : nom de la ville d'arrivée
* @param distance : distance entre les deux villes
 */
func New(ville1 ville.Ville, ville2 ville.Ville, distance int) Trajet {
	traj := Trajet{ville1, ville2, distance}
	return traj
}

func (t Trajet) GetVille1() ville.Ville {
	return t.ville1
}

func (t Trajet) GetVille2() ville.Ville {
	return t.ville2
}

func (t Trajet) GetDistance() int {
	return t.distance
}

func (t Trajet) SetDistance(distance int) {
	t.distance = distance
}
