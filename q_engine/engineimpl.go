package q_engine

import (
		"fmt"
		gfx "quoridor/gfx"
		arrayref "quoridor/q_arrayref"
		spieler "quoridor/q_spieler"
		figur "quoridor/q_figur"
		waende "quoridor/q_waende"
		)

type engine struct {
	a arrayref.Arrayref
	w [20]waende.Wand
	h,b,x,y,größe uint16
	sp [4]spieler.Spieler
	spielerz int
}

func New() *engine {
	var e *engine
	e=new(engine)
	return e
}

func (e *engine) Spiel () {
	h,b:=e.start()
	gfx.Fenster(b,h)
	
	
	var cp,gewinner int //current player
	gewinner = -1
	var beendet bool
	for !beendet{
		e.zug(cp)
		gewinner,beendet = e.a.Gewonnen()
		gfx.UpdateAus()
		e.draw(false)
		gfx.UpdateAn()
		//if e.figurzug(cp) {					//noch nicht fertig,nur für figur, diese 2 zeilen sind demo/test
			//fmt.Println("figur")
		//}
		//if e.Figurzug {
			//e.Figurbewegen(cp)
			//e.Hatfigurgewonnen(cp)
		//} else {
			//e.Wandsetzen(cp)
		//}
		//e.draw()
		//gfx.TastaturLesen1()		//da bisher nur anzeige von Startfeld
		cp=(cp+1)%e.spielerz
	}
	var siegername string
	fmt.Println("Herzlichen Glückwunsch!!")
	for i:=0;i<len(e.sp);i++{
		if e.sp[i]!=nil{
			if e.sp[i].GebeNummer()==gewinner {
				siegername = e.sp[i].GebeName()
			}
		}
	}
	fmt.Println(siegername,"hat gewonnen!!!")
}

func (e *engine) zug (cp int) {
	e.a.HighlighteAktuellenSpieler(cp)
outer:
	for {
	start:
	gfx.UpdateAus()
	e.draw(false)
	gfx.UpdateAn()
		t,s,mx,my:=gfx.MausLesen1()
		for l:=0;l<len(e.w);l++{
			if t==1&&s==1&&e.w[l].PunktgehörtzurWand(mx,my) && e.w[l].Gebebesitz()==cp {
				var ux,uy uint16
				ux,uy=e.w[l].GebeKoordinaten ()
				gfx.UpdateAus()
				e.draw(true)
				gfx.UpdateAn()
				for !e.w[l].IstPlaziert(){
					t,s,mx,my:=gfx.MausLesen1()
					if t==1&&s==1{
						var wliste [19]waende.Wand
						for i:=0;i<len(e.w);i++{
							if i !=l && l>i {
								wliste[i] = e.w[i]
							} else if i !=l && i>l {
								wliste[i-1] = e.w[i]
							}
						}
						e.a.WandPlazieren(e.w[l],mx,my,wliste)
						if e.w[l].IstPlaziert(){
							break outer
						}
					}else if t==3&&s==1{
						e.w[l].SwitchOrientation()
					}
					br,h:=e.w[l].GebeGroesse()
					e.w[l].SetzeKoordinaten(mx-br/2,my-h/2)
					gfx.UpdateAus()
					e.draw(true)
					gfx.UpdateAn()
					i,j,erg:=e.a.GehörtPunktzuFigur(mx,my)
					ci,cj:=e.a.Gibposition(cp)	
					if erg&&i==ci&&j==cj&&t==1&&s==1{
						if e.w[l].IstSenkrecht(){
							e.w[l].SwitchOrientation()
						}
						e.w[l].SetzeKoordinaten(ux,uy)
						goto start
					}
				}
			}
		}
					
					
		i,j,erg:=e.a.GehörtPunktzuFigur (mx,my)		
		if erg {
			ci,cj:=e.a.Gibposition(cp)
			if i==ci&&j==cj&&t==1&&s==1{
				e.a.HighlighteZüge(i,j)
				gfx.UpdateAus()
				e.draw(false)
				gfx.UpdateAn()
				for {
					t,s,mx,my:=gfx.MausLesen1()
					k,l,erg:=e.a.GehörtPunktzuFeld (mx,my)
					if erg&&e.a.IstGehighlightet(k,l)&&t==1&&s==1 {
						e.a.Setzeposition(cp,k,l)
						e.a.LöscheFeldHighlights ()
						gfx.UpdateAus()
						e.draw(false)
						gfx.UpdateAn()
						break outer
					}
					for l:=0;l<len(e.w);l++{
						if t==1&&s==1&&e.w[l].PunktgehörtzurWand(mx,my) && e.w[l].Gebebesitz()==cp {
							e.a.LöscheFeldHighlights ()
							goto start
						}
					}
					
				}
			}
		}
	}
}
				

//func (e *engine) figurzug (cp int) bool{
	//var figur bool
	//for {
		//t,s,mx,my:=gfx.MausLesen1()
		//if e.sf[cp].GehoertPunktZuFigur(mx,my) {
			//switch cp {
				//case 0:
					//e.sf[cp].SetzeFarbe(0,0,0)		//hier geeignete farben suchen
				//case 1:
					//e.sf[cp].SetzeFarbe(0,0,0)
				//case 2:
					//e.sf[cp].SetzeFarbe(0,0,0)
				//case 3:
					//e.sf[cp].SetzeFarbe(0,0,0)
			//}
			//if s!=0&&t!=50 {							//t musste benutzt werden, ich konnte t declared but not used nicht umgehen
				//figur=true
				//e.figurfarbwechsel=false
				//break
			//}
		//}
		////for i:=0;i<len(e.w);i++ {
			////if e.w[i].GehoertPunktZuFigur(mx,my) {
				////e.w[i].SetzeFarbe(50,50,50)
			////}
		////}
		//e.draw()
	//}
	//return figur
//}

func (e *engine) start () (uint16,uint16) {		//spielerz= anzhal spieler
	var h,b,x,y,größe uint16
	var z int
	for {
		fmt.Println("Wieviele Spieler seid ihr?")
		fmt.Println("(2-4)")
		fmt.Scanln(&z)
		if 1<z&&z<5 {
			break
		}
	}
	e.spielerz=z
	for i:=0;i<e.spielerz;i++ {
		e.sp[i]=spieler.New()
		var name string
		fmt.Println("Name Spieler",i+1,":")
		fmt.Scan(&name)
		e.sp[i].SetzeName(name)
		e.sp[i].SetzeNummer(i)
	}
	fmt.Println("höhe in pixeln:")
	fmt.Scanln(&h)
	b=13*h/10			//b=1,3h bei angewendeten verhältnissen
	größe= (3*h)/41		// für volle befüllung des Bildschirms
	x=(b-(13*(größe+größe/3)+größe/3))/2+3*größe	//mittige Ausrichtung
	y=(h-(9*(größe+größe/3)+größe/3))/2+größe/3
	e.a= arrayref.New(x,y,größe)
	
	for i:=0;i<len(e.w);i++ {
		e.w[i]=waende.New()
		e.w[i].SetzeGroesse(2*größe+größe/3,größe/3)
	}
	for i:=0;i<len(e.w)/2;i++ {
		e.w[i].SetzeKoordinaten(x-3*größe+größe/3,y-größe/3+uint16(i)*(größe+größe/3))
	}
	for i:=len(e.w)/2;i<len(e.w);i++ {
		e.w[i].SetzeKoordinaten(x+9*(größe+größe/3),y-größe/3+uint16(i-len(e.w)/2)*(größe+größe/3))
		e.w[i].Setzebesitz(1)
	}
	
	for i:=0;i<e.spielerz;i++ {
		var f figur.Figur
		f=figur.New()
		f.SetzeGroesse(größe/2)
		switch i {
			case 0:
				f.SetzeSpieler(i)
				f.SetzeFarbe(255,255,255)
				f.SetzeHighlightFarbe(150,150,150)
				e.a.SetzeSpielstein(f,4,0)
			case 1:
				f.SetzeSpieler(i)
				f.SetzeFarbe(255,255,0)
				f.SetzeHighlightFarbe(150,150,0)
				e.a.SetzeSpielstein(f,4,8)
			case 2:
				f.SetzeSpieler(i)
				f.SetzeFarbe(0,0,255)
				f.SetzeHighlightFarbe(0,0,150)
				e.a.SetzeSpielstein(f,0,4)
			case 3:
				f.SetzeSpieler(i)
				f.SetzeFarbe(0,255,0)
				f.SetzeHighlightFarbe(0,150,0)
				e.a.SetzeSpielstein(f,8,4)
		}
	}
	e.a.SetzeZielZonen()
	e.h,e.b,e.x,e.y,e.größe=h,b,x,y,größe
	return h,b
}

func (e *engine) draw (ww bool) {
	x,y,größe:=e.x,e.y,e.größe
	gfx.Stiftfarbe(136,0,0)	
	gfx.Vollrechteck(x-3*größe,y-größe,13*(größe+größe/3)+größe/3,9*(größe+größe/3)+5*größe/3)		//Grundrechteck das über felder hinausgeht, ist feldummantelung
	if ww {
		gfx.Stiftfarbe(0,255,0)
		gfx.Vollrechteck(x,y,(größe*9+8*größe/3)-2,(größe*9+8*größe/3)-2)	//"Wände", wird durch a.Draw übermalt wo keine wand ist
	}
	e.a.Draw()
	for i:=0;i<len(e.w);i++ {
		e.w[i].Draw()
	}
}
