package q_spieler


type Spieler interface {

	//Vor.: -
	//Eff.: Name wird gesetzt
	SetzeName (name string)
	String () string
	GebeNummer () int
	GebeName () string
	SetzeNummer (nr int)
}

