package q_waende

import . "quoridor/gfx"
import "fmt"

type data struct {
	x,y uint16
	h,br uint16
	r,g,b uint8
	sp int
	plaziert bool
}

func New() *data {
	var w *data
	w = new(data)
	return w
}

func (w *data) SetzeKoordinaten(x,y uint16) {
	w.x = x
	w.y = y
}

func (w *data) Plaziert (p bool) {
	w.plaziert = p
}

func (w *data) IstPlaziert () bool{
	return w.plaziert
}

func (w *data) Setzebesitz(sp int) {
	w.sp=sp
}

func (w *data) Gebebesitz() (sp int) {
	return w.sp
}

func (w *data) GebeKoordinaten () (x,y uint16) {
	return w.x, w.y
}

func (w *data) SetzeGroesse(br,h uint16) {
	w.br = br
	w.h = h
}

func (w *data) IstSenkrecht() bool {
	return w.h>w.br
}

func (w *data) GebeGroesse()(br,h uint16) {
	return w.br, w.h
}

func (w *data) SetzeFarbe(r,g,b uint8) {
	w.r = r
	w.g = g
	w.b = b
}

func (w *data) SwitchOrientation () {
	w.h,w.br = w.br,w.h
}

func (w *data) GebeFarbe()(r,g,b uint8) {
	return w.r, w.g, w.b
}

func (w *data) Draw () {
	Stiftfarbe (w.r,w.g,w.b)
	Vollrechteck(w.x,w.y,w.br,w.h)
}

func (w *data) PunktgehörtzurWand(mx,my uint16) bool {
	return mx>=w.x&&mx<=w.x+w.br&&my>=w.y&&my<=w.y+w.h
}


func (w *data) String () string {
	var erg string
	erg = erg + fmt.Sprintln ("Koordinaten: ",w.x,w.y)
	erg = erg + fmt.Sprintln ("Groesse: ",w.br,w.h)
	erg = erg + fmt.Sprintln ("Farbe: ",w.r,w.g,w.b)
	return erg
}




