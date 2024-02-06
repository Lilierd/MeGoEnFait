package activity

type Activity struct {
	Name     string
	Duree    int
	Parents  []Activity
	PlusTot  int
	PlusTard int
}

func New(nom string, duree int, parents []Activity) Activity {
	act := Activity{nom, duree, parents, 0, 0}
	return act
}
