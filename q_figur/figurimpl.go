package q_figur
import gfx "quoridor/gfx"
import "fmt"

type data struct {
	r, g, b uint8
	hr,hg,hb uint8
	x, y, radius uint16
	spieler int
	highlight bool
}

func New() *data {
	var f *data 
	f = new (data)
	return f
}

func (f *data) SetzeKoordinaten (x, y uint16) {
	f.x = x
	f.y = y
}

func (f *data) GebeKoordinaten () (x,y uint16) {
	return f.x, f.y
}

func (f *data) SetzeFarbe (r, g, b uint8) {
	f.r = r
	f.g = g
	f.b = b
}

func (f *data) SetzeHighlightFarbe (r, g, b uint8) {
	f.hr = r
	f.hg = g
	f.hb = b
}

func (f *data) SetzeHighlight (h bool) {
	f.highlight = h
}

func (f *data) SetzeSpieler (spieler int) {
	f.spieler = spieler
}

func (f *data) GebeSpieler () int {
	return f.spieler
}

func (f *data) SetzeGroesse (radius uint16) {
	f.radius = radius
}

func (f *data) GebeGroesse()(uint16) {
	return f.radius
}

func (f *data) GebeFarbe () (r, g, b uint8) {
	return f.r, f.g, f.b
}

func (f *data) Draw() {
	if f.highlight {
		gfx.Stiftfarbe(f.r,f.g,f.b)
	}else{
		gfx.Stiftfarbe(f.hr,f.hg,f.hb)
	}    
	gfx.Vollkreis(f.x,f.y,f.radius) 
	gfx.Stiftfarbe(0,0,0)      
	gfx.Kreis(f.x,f.y,f.radius) 
}

func (f *data) GehoertPunktZuFigur(x,y uint16) bool {
	return (int(x)-int(f.x))*(int(x)-int(f.x)) + (int(y)-int(f.y))*(int(y)-int(f.y))<=int(f.radius)*int(f.radius)
	}
	
/*??? funktion um zu sehen, ob spieler die figur setzen möchte (angeklickt hat)
 * 
 * func (f *data) IstFigurAusgewählt () {
 * var auswahl bool
	
	if GehörtPunktZuFigur = true
	&& Mauslesen1 = (taste, 1, mausx, mausy)
	auswahl = 1
	* } else {
	* auswahl = 0
	
}
*/

func (f *data) String () string {
	var erg string
	erg = erg + fmt.Sprintln ("Koordinaten: ",f.x,f.y)
	erg = erg + fmt.Sprintln("Radius: ", f.radius)
	erg = erg + fmt.Sprintln("Farbe in RGB: ",f.r,f.g,f.b)
	return erg
}
//string funktion
