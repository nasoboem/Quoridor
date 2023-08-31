package q_waende

//import "gfx"

type Wand interface {
	
	SetzeKoordinaten(x,y uint16)
	
	GebeKoordinaten () (x,y uint16)
	
	Setzebesitz(sp int)

	Gebebesitz() (sp int)
	
	SetzeGroesse(br,h uint16) 
	 
	GebeGroesse()(br,h uint16)
	 
	SetzeFarbe(r,g,b uint8) 
	 
	GebeFarbe()(r,g,b uint8) 
	
	SwitchOrientation ()
	
	IstSenkrecht () bool
	
	Plaziert (p bool)
	
	IstPlaziert () bool
	
	PunktgehörtzurWand(mx,my uint16) bool
	 
	Draw () 
	
}


func Kollision (w1, w2 Wand) bool {
	var x1,x2,y1,y2,b1,b2,h1,h2 uint16
	x1,y1 = w1.GebeKoordinaten()
	x2,y2 = w2.GebeKoordinaten()
	b1,h1 = w1.GebeGroesse()
	b2,h2 = w2.GebeGroesse()
	return x1+b1>x2&&x2+b2>x1&&y1+h1>y2&&y2+h2>y1
}
