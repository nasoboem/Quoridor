package q_figur


type Figur interface{
	
	SetzeKoordinaten (x,y uint16)
	
	GebeKoordinaten () (x, y uint16)
	
	GebeFarbe () (r,g,b uint8)
	
	SetzeFarbe (r,g,b uint8)
	 
	GebeGroesse () (radius uint16)
	
	SetzeGroesse (radius uint16)
	
	GehoertPunktZuFigur(x,y uint16) bool
	
	SetzeSpieler (spieler int)
	
	GebeSpieler () int
	
	SetzeHighlight (h bool)
	
	SetzeHighlightFarbe (r, g, b uint8)
	
	Draw()
	
}


