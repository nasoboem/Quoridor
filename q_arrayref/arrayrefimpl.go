package q_arrayref

import (
  gfx "quoridor/gfx"
  "fmt"
  figur "quoridor/q_figur"
  waende "quoridor/q_waende"
)


type arrayref struct {
	array [9][9]*felder
}

type felder struct {
	x,y,größe uint16
	nachbar [4]*felder //0 = oben,1 = rechts, 2, = unten , 3 = links
	spielstein figur.Figur
	highlight bool
	zielvon int
}


func New(x,y,größe uint16) *arrayref {
	var a *arrayref
	a =new(arrayref)

	for i:=0;i<len(a.array);i++{
		for j:=0;j<len(a.array[i]);j++{
				var f *felder
				f = new(felder)
				f.zielvon = -1
				f.größe = größe
				f.setzeKoordinaten(x,y,größe,i,j)
				a.array[i][j] = f
		}
	}
	for i:=0;i<len(a.array);i++{
		for j:=0;j<len(a.array[i]);j++{
			var o,r,u,l *felder
			if j-1>=0 {
				o = a.array[i][j-1]
			}
			if j+1<len(a.array[i]) {
				u = a.array[i][j+1]
			}
			if i-1>=0 {
				l = a.array[i-1][j]
			}
			if i+1<len(a.array) {
				r = a.array[i+1][j]
			}
			a.array[i][j].setzeNachbarn(o,r,u,l)
		}
	}
	return a
}

func (f *felder) setzeNachbarn(o,r,u,l *felder) {
	f.nachbar[0] = o
	f.nachbar[1] = r
	f.nachbar[2] = u
	f.nachbar[3] = l
}

func (a *arrayref) PrintHighlight() string {
	var erg string
	for i:=0;i<len(a.array);i++{
		for j:=0;j<len(a.array[i]);j++{
			if j==len(a.array[i])-1 {
				erg = erg + fmt.Sprintln(a.array[i][j].highlight)
			 }else{
				 erg = erg + fmt.Sprint(a.array[i][j].highlight)
			 }
		 }
	 }
	 return erg
 }


func (a *arrayref) Gewonnen() (int,bool) {
	var erg bool
	var sieger int
outer:
	for i:=0;i<len(a.array);i++{
		for j:=0;j<len(a.array[i]);j++{
			if a.array[i][j].spielstein!=nil{
				if a.array[i][j].spielstein.GebeSpieler()==a.array[i][j].zielvon{
					erg = true
					sieger = a.array[i][j].spielstein.GebeSpieler()
					break outer
				}
			}
		}
	}
	return sieger,erg
}

func (a *arrayref) cralToEnd () bool {
	var spnr []int
	for i:=0;i<len(a.array);i++{
		for j:=0;j<len(a.array[i]);j++{
			if a.array[i][j].spielstein!=nil{
				spnr = append(spnr,a.array[i][j].spielstein.GebeSpieler (),i,j)
			}
		}
	}
	var enderg []bool
	for i:=0;i<len(spnr);i = i+3{
		var erg bool
		erg = a.array[spnr[i+1]][spnr[i+2]].craler(spnr[i])
		enderg = append(enderg,erg)
		a.LöscheFeldHighlights()
	}
	var erg bool
	erg = true
	for i:=0;i<len(enderg);i++{
		 if !enderg[i] {
			 return false
		 }
	}
	return erg 
}

func (f *felder) craler (sp int) bool {
	var erg bool
	f.highlight = true
	if f.zielvon==sp {
		erg = true
		return erg
	} else {
		for i:=0;i<len(f.nachbar);i++{
			if f.nachbar[i]!=nil{
				if !f.nachbar[i].highlight {
					erg = erg || f.nachbar[i].craler(sp)
				}
			}
		}
	}
	return erg
}
			
				

func (a *arrayref) SetzeZielZonen () {
	for i:=0;i<len(a.array);i++{
		for j:=0;j<len(a.array[i]);j++{
			if a.array[i][j].spielstein!=nil{
				for k:=0;k<4;k++{
					if a.array[i][j].nachbar[k]==nil {
						for l:=0;l<len(a.array);l++{
							for m:=0;m<len(a.array[l]);m++{
								if a.array[l][m].nachbar[(k+2)%4]==nil {
									a.array[l][m].zielvon = a.array[i][j].spielstein.GebeSpieler()
								}
							}
						}
					}
				}
			}
		}
	}
}

func (f *felder) löschenachbar(i int) *felder{
	var oldf *felder
	if f.nachbar[i]!=nil {
		oldf = f.nachbar[i]
	}
	f.nachbar[i] = nil
	return oldf 
}

func (f *felder) setzeEinzelNachbar (nummer int, of *felder) {
	f.nachbar[nummer]=of
}


func (f *felder) hatnachbar(i int) bool {
	return f.nachbar[i] != nil
}


func (a *arrayref) WandPlazieren(w waende.Wand,mx,my uint16,ww [19]waende.Wand) {
	var ol,or,ul,ur *felder
	for i:=0;i<len(a.array);i++{
		for j:=0;j<len(a.array[i]);j++{
			var x,y,größe uint16
			größe = a.array[i][j].gebeGröße()
			x,y = a.array[i][j].gebeKoordinaten()
			switch {
				case mx>=x+größe&&mx<=x+größe+größe/3&&my>=y+größe&&my<=y+größe+größe/3:
					ol = a.array[i][j]
				case mx>=x-größe/3&&mx<=x&&my>=y+größe&&my<=y+größe+größe/3:
					or = a.array[i][j]
				case mx>=x+größe&&mx<=x+größe+größe/3&&my>=y-größe/3&&my<=y:
					ul = a.array[i][j]
				case mx>=x-größe/3&&mx<=x&&my>=y-größe/3&&my<=y:
					ur = a.array[i][j]
			}
		}
	}
	if ol!=nil&&or!=nil&&ul!=nil&&ur!=nil {
		if w.IstSenkrecht() {
					größe:=ol.gebeGröße()
					x,y:=ol.gebeKoordinaten()
					w.SetzeKoordinaten(x+größe,y)
			}else{
					größe:=ol.gebeGröße()
					x,y:=ol.gebeKoordinaten()
					w.SetzeKoordinaten(x,y+größe)
			}
		var kollision bool 
		for i:=0;i<len(ww);i++{
			if waende.Kollision(w,ww[i]) {
				kollision = true
				break
			}
		}
		if !kollision {
			var ool,oor,oul,our *felder
			var plazierbar bool
			if w.IstSenkrecht() {
					ool = ol.löschenachbar(1)
					oor = or.löschenachbar(3)
					oul = ul.löschenachbar(1)
					our = ur.löschenachbar(3)
			}else{
					ool = ol.löschenachbar(2)
					oor = or.löschenachbar(2)
					oul = ul.löschenachbar(0)
					our = ur.löschenachbar(0)
			}
			plazierbar = a.cralToEnd()
			if plazierbar {
				w.Plaziert(true)
			} else {
				if w.IstSenkrecht() {
					ol.setzeEinzelNachbar(1,ool)
					or.setzeEinzelNachbar(3,oor)
					ul.setzeEinzelNachbar(1,oul)
					ur.setzeEinzelNachbar(3,our)
				}else{
					ol.setzeEinzelNachbar(2,ool)
 					or.setzeEinzelNachbar(2,oor)
					ul.setzeEinzelNachbar(0,oul)
					ur.setzeEinzelNachbar(0,our)
			}
		}
		}
	}
}

func (a *arrayref) Gibposition(sp int) (x,y uint16) {
outer:
	for i:=0;i<len(a.array);i++{
		for j:=0;j<len(a.array[i]);j++{
			if a.array[i][j].spielstein!=nil{
				if a.array[i][j].spielstein.GebeSpieler() == sp {
					x = uint16(i)
					y = uint16(j)
					break outer
				}
			}
		}
	}
	return
}

func (a *arrayref) IstGehighlightet (i,j uint16) bool {
	return a.array[int(i)][int(j)].highlight
}

func (a *arrayref) HighlighteZüge (i,j uint16) {
	for k:=0;k<4;k++{																						//vier Richtungen 0 = oben, 1 = rechts,2 = unten,3 = links 
		if a.array[int(i)][int(j)].nachbar[k]!=nil && a.array[int(i)][int(j)].nachbar[k].spielstein==nil{	//wenn ein nachbar Feld da ist und es frei ist
			a.array[int(i)][int(j)].nachbar[k].highlight=true								    			//setze highlight
		}else if  a.array[int(i)][int(j)].nachbar[k]!=nil && a.array[int(i)][int(j)].nachbar[k].spielstein!=nil {										//wenn das nachbar Feld nicht frei ist
			if a.array[int(i)][int(j)].nachbar[k].nachbar[k]!=nil&&a.array[int(i)][int(j)].nachbar[k].nachbar[k].spielstein==nil{ 	//wenn der nachbar vom nachbar existiert und frei ist
					a.array[int(i)][int(j)].nachbar[k].nachbar[k].highlight=true													//dann setze highlight beim Nachbarn vom Nachbarn
			}else if  a.array[int(i)][int(j)].nachbar[k].nachbar[k]==nil {															//wenn der der Nachbar vom Nachbarn nicht existiert											
					if a.array[int(i)][int(j)].nachbar[k].nachbar[(k+1)%4]!=nil  { //Teste jeweils ob die seiten existieren 
						a.array[int(i)][int(j)].nachbar[k].nachbar[(k+1)%4].highlight=true
					}
					if a.array[int(i)][int(j)].nachbar[k].nachbar[(k+3)%4]!=nil {
						a.array[int(i)][int(j)].nachbar[k].nachbar[(k+3)%4].highlight=true
					}
			}	
		}
	}
}

func (a *arrayref) LöscheFeldHighlights () {
	for i:=0;i<len(a.array);i++{
		for j:=0;j<len(a.array[i]);j++{
			for k:=0;k<4;k++{
				if a.array[i][j].nachbar[k]!=nil{
					a.array[i][j].nachbar[k].highlight=false
				}
			}
		}
	}
}

func (f *felder) istnachbar (g *felder) bool {
	var erg bool
	if g!=nil {
		erg = f.nachbar[0]==g||f.nachbar[1]==g||f.nachbar[2]==g||f.nachbar[3]==g
	}
	return erg
}

func (a *arrayref) SindNachbarn(ai,aj,bi,bj uint16) bool {
	return a.array[int(ai)][int(aj)].istnachbar(a.array[int(bi)][int(bj)])
}

func (a *arrayref) Setzeposition(sp int, x,y uint16) {
	var xf,yf uint16
	xf,yf = a.Gibposition(sp)
	a.array[int(xf)][int(yf)].spielstein,a.array[int(x)][int(y)].spielstein = a.array[int(x)][int(y)].spielstein,a.array[int(xf)][int(yf)].spielstein
	a.updateSpielstein ()
}

func (a *arrayref) SetzeSpielstein (f figur.Figur,x,y uint16) {
	var fx,fy,fgröße uint16
	fx,fy = a.array[int(x)][int(y)].gebeKoordinaten()
	fgröße = a.array[int(x)][int(y)].gebeGröße()
	f.SetzeKoordinaten(fx+fgröße/2,fy+fgröße/2)
	a.array[int(x)][int(y)].spielstein=f
}

func (a *arrayref) updateSpielstein () {
	for i:=0;i<len(a.array);i++{
			for j:=0;j<len(a.array[i]);j++{
					if a.array[i][j].spielstein!=nil{
						var fx,fy,fgröße uint16
						fx,fy = a.array[i][j].gebeKoordinaten()
						fgröße = a.array[i][j].gebeGröße()
						a.array[i][j].spielstein.SetzeKoordinaten(fx+fgröße/2,fy+fgröße/2)
					}
			}
	}
}


func (a *arrayref) LöscheNachbar(i,j,r int) {
	if a.array[i][j].hatnachbar(r){
		a.array[i][j].löschenachbar(r)
		switch r {
			case 0: //oben
			a.array[i][j-1].löschenachbar(2)
			case 1: //rechts
			a.array[i+1][j].löschenachbar(3)
			case 2: //unten
			a.array[i][j+1].löschenachbar(0)
			case 3: // links
			a.array[i-1][j].löschenachbar(1)
		}
	}
}

func (a *arrayref) HighlighteAktuellenSpieler(cp int) {
	for i:=0;i<len(a.array);i++{
		for j:=0;j<len(a.array[i]);j++{
			if a.array[i][j].spielstein!=nil {
				a.array[i][j].spielstein.SetzeHighlight(false)
			}
		}
	}
	x,y:=a.Gibposition(cp)
	a.array[int(x)][int(y)].spielstein.SetzeHighlight(true)
}


func (f *felder) draw () {
	var x,y,größe uint16
	x = f.x
	y = f.y
	größe = f.größe
	if f.highlight {
		gfx.Stiftfarbe(150,150,0)
	}else{
		gfx.Stiftfarbe(136,136,136)
	}
	gfx.Vollrechteck(x,y,größe,größe)
	for i:=0;i<len(f.nachbar);i++{
		if f.nachbar[i]!=nil{
			switch i {
				case 0: //oben
				gfx.Stiftfarbe(136,0,0)
				gfx.Vollrechteck(x,y-größe/3,größe,größe/3)
				case 1: //rechts
				gfx.Stiftfarbe(136,0,0)
				gfx.Vollrechteck(x+größe,y,größe/3,größe)
				case 2: //unten
				gfx.Stiftfarbe(136,0,0)
				gfx.Vollrechteck(x,y+größe,größe,größe/3)
				case 3: // links
				gfx.Stiftfarbe(136,0,0)
				gfx.Vollrechteck(x-größe/3,y,größe/3,größe)
			}
		}
	}
}
		
func (a *arrayref) GehörtPunktzuFigur (mx,my uint16) (uint16,uint16,bool) {
	var x,y uint16
	var erg bool
outer:
	for i:=0;i<len(a.array);i++{
		for j:=0;j<len(a.array[i]);j++{
			if  a.array[i][j].spielstein!=nil{
				if a.array[i][j].spielstein.GehoertPunktZuFigur(mx,my) {
					erg = a.array[i][j].spielstein.GehoertPunktZuFigur(mx,my)
					x = uint16(i)
					y = uint16(j)
					break outer
				}
			}
		}
	}
	return x,y,erg
}

func (a *arrayref) GehörtPunktzuFeld (mx,my uint16) (uint16,uint16,bool) {
	var x,y,fx,fy,größe uint16
	var erg bool
outer:
	for i:=0;i<len(a.array);i++{
		for j:=0;j<len(a.array[i]);j++{
			größe = a.array[i][j].gebeGröße()
			fx,fy = a.array[i][j].gebeKoordinaten()
			if mx>=fx && mx<=fx+größe && my>=fy && my<=fy+größe {
				erg = true
				x = uint16(i)
				y = uint16(j)
				break outer
			}
		}
	}
	return x,y,erg
}

func (f *felder) setzeKoordinaten (x,y,größe uint16, i,j int) {
	f.x = x+uint16(i)*(größe+größe/3)
	f.y = y+uint16(j)*(größe+größe/3)
}

func (f *felder) gebeGröße () uint16 {
	return f.größe
}

func (f *felder) gebeKoordinaten () (uint16,uint16) {
	return f.x,f.y
}

func (a *arrayref) Draw () {
	for i:=0;i<len(a.array);i++{
		for j:=0;j<len(a.array[i]);j++{
			a.array[i][j].draw()
			if a.array[i][j].spielstein!=nil {
				a.array[i][j].spielstein.Draw()
			
		}
	}
	//if a.spieler1[0]!=nil&&a.spieler1[1]!=nil{												So besprochen, jedoch int kann nicht nil sein
		//gfx.Stiftfarbe(255,0,0)																deshalb figurzeichnen über figurpaket.
		//i,j:=uint16(a.spieler1[0]),uint16(a.spieler1[1])										wenn nicht 4 spieler, werden feldangaben auf 10,10 gesetzt, so keine komplikationen beim bewegenbefehl
		//gfx.Vollkreis(x+i*(größe+größe/3)+größe/2,y+j*(größe+größe/3)+größe/2,größe/2)
	//}
	//if a.spieler2[0]!=nil&&a.spieler2[1]!=nil{
		//gfx.Stiftfarbe(0,255,0)
		//i,j:=uint16(a.spieler2[0]),uint16(a.spieler2[1])
		//gfx.Vollkreis(x+i*(größe+größe/3)+größe/2,y+j*(größe+größe/3)+größe/2,größe/2)
	//}
	//if a.spieler3[0]!=nil&&a.spieler3[1]!=nil{
		//gfx.Stiftfarbe(0,0,255)
		//i,j:=uint16(a.spieler3[0]),uint16(a.spieler3[1])
		//gfx.Vollkreis(x+i*(größe+größe/3)+größe/2,y+j*(größe+größe/3)+größe/2,größe/2)
	//}
	//if a.spieler4[0]!=nil&&a.spieler4[1]!=nil{
		//gfx.Stiftfarbe(255,255,0)
		//i,j:=uint16(a.spieler4[0]),uint16(a.spieler4[1])
		//gfx.Vollkreis(x+i*(größe+größe/3)+größe/2,y+j*(größe+größe/3)+größe/2,größe/2)
	//}
	
	
	}
}
