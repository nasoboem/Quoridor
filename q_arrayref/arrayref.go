package q_arrayref

import figur "quoridor/q_figur"

import waende "quoridor/q_waende"

type Arrayref interface{
	
	Draw ()
	
	LöscheNachbar(i,j,r int)
	
	Gibposition(sp int) (xi,yi uint16)
	
	Setzeposition(sp int, x,y uint16)
	
	SetzeSpielstein (f figur.Figur,x,y uint16)
	
	HighlighteAktuellenSpieler(cp int)
	
	GehörtPunktzuFigur (mx,my uint16) (uint16,uint16,bool)
	
	LöscheFeldHighlights ()
	
	HighlighteZüge (i,j uint16)
	
	IstGehighlightet (i,j uint16) bool
	
	WandPlazieren(w waende.Wand,mx,my uint16,ww [19]waende.Wand)
	
	SindNachbarn(ai,aj,bi,bj uint16) bool
	
	SetzeZielZonen ()
	
	Gewonnen() (int,bool)
	
	GehörtPunktzuFeld (mx,my uint16) (uint16,uint16,bool) //i,j, treffer
	
	PrintHighlight() string
	
}

