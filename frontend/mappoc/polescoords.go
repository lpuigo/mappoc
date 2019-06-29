package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/lpuig/ewin/doe/website/frontend/tools"
)

type BePole struct {
	Ref  string
	City string
	Lat  float64
	Long float64
}

type Pole struct {
	*js.Object
	Ref        string      `js:"Ref"`
	City       string      `js:"City"`
	Lat        float64     `js:"Lat"`
	Long       float64     `js:"Long"`
	Done       bool        `js:"Done"`
	PoleMarker *PoleMarker `js:"PoleMarker"`
}

func NewPole(pole BePole) *Pole {
	np := &Pole{
		Object: tools.O(),
	}

	np.Ref = pole.Ref
	np.City = pole.City
	np.Lat = pole.Lat
	np.Long = pole.Long
	np.Done = false
	np.PoleMarker = nil

	return np
}

func GetCenterAndBounds(poles []*Pole) (clat, clong, blat1, blong1, blat2, blong2 float64) {
	if len(poles) == 0 {
		return 47, 5, 46.5, 4.5, 47.5, 5.5
	}

	min := func(pole *Pole) {
		if pole.Lat < blat1 {
			blat1 = pole.Lat
		}
		if pole.Long < blong1 {
			blong1 = pole.Long
		}
	}

	max := func(pole *Pole) {
		if pole.Lat > blat2 {
			blat2 = pole.Lat
		}
		if pole.Long > blong2 {
			blong2 = pole.Long
		}
	}

	blat1, blong1 = 500, 500
	for _, pole := range poles {
		clat += pole.Lat
		clong += pole.Long
		min(pole)
		max(pole)
	}

	nb := float64(len(poles))
	clat /= nb
	clong /= nb
	return
}

func GenPoles(poles []BePole) []*Pole {
	res := make([]*Pole, len(poles))

	for i, pole := range poles {
		res[i] = NewPole(pole)
	}
	return res
}

var poles = []BePole{
	//{"MF22A01", "Maizières-lès-Vic", 48.68398083, 6.767729722},
	//{"MF22A02", "Maizières-lès-Vic", 48.6867125, 6.770223056},
	//{"MF22A03", "Maizières-lès-Vic", 48.69282139, 6.762173611},
	//{"MF22A04", "Maizières-lès-Vic", 48.68570194, 6.778500278},
	//{"MF22A05", "Maizières-lès-Vic", 48.68481806, 6.779818333},
	//{"MF22A06", "Maizières-lès-Vic", 48.68511778, 6.7796225},
	{"MF22039", "Avricourt", 48.65268361, 6.804828611},
	{"MF22009", "Avricourt", 48.65007417, 6.805681111},
	{"MF22011", "Avricourt", 48.65011278, 6.805285833},
	{"MF22006", "Maizière-lès-Vic", 48.70158583, 6.797752778},
	{"MF22103", "Avricourt", 48.65231528, 6.808456389},
	{"MF22019", "Avricourt", 48.65126222, 6.806866667},
	{"MF22023", "Avricourt", 48.65164306, 6.805110278},
	{"MF22026", "Avricourt", 48.65183056, 6.803511111},
	{"MF22027", "Avricourt", 48.65202222, 6.80309},
	{"MF22031", "Avricourt", 48.65296444, 6.800733611},
	{"MF22036", "Avricourt", 48.65221222, 6.805172222},
	{"MF22043", "Avricourt", 48.65489806, 6.804246667},
	{"MF22187", "Avricourt", 48.6496425, 6.806759444},
	{"MF22053", "Avricourt", 48.64927167, 6.806758611},
	{"MF22054", "Avricourt", 48.64892583, 6.806525278},
	{"MF22062", "Avricourt", 48.64828, 6.810688056},
	{"MF22064", "Avricourt", 48.6487175, 6.811341389},
	{"MF22065", "Avricourt", 48.64879778, 6.811813333},
	{"MF22072", "Avricourt", 48.65029611, 6.811665556},
	{"MF22077", "Avricourt", 48.65082528, 6.809250833},
	{"MF22078", "Avricourt", 48.65110833, 6.809428056},
	//{"MF22085", "Avricourt", 48.65008667, 6.808467778}, // Not allocated
	{"MF22090", "Avricourt", 48.65129944, 6.807456389},
	{"MF22091", "Avricourt", 48.65164417, 6.807748333},
	{"MF22070", "Avricourt", 48.64959917, 6.811748889},
	{"MF22086", "Avricourt", 48.64994, 6.808153889},
	{"MF22106", "Avricourt", 48.65023917, 6.812333056},
	{"MF22141", "Avricourt", 48.65228833, 6.822251389},
	{"MF22143", "Avricourt", 48.65203167, 6.822015278},
	{"MF22148", "Avricourt", 48.65144694, 6.821033889},
	{"MF22146", "Avricourt", 48.65239083, 6.820795833},
	{"MF22154", "Avricourt", 48.65195167, 6.819122778},
	{"MF22156", "Avricourt", 48.65060167, 6.819590278},
	{"MF22158", "Avricourt", 48.65102222, 6.818688611},
	{"MF22162", "Avricourt", 48.65342389, 6.823642222},
	{"MF22163", "Avricourt", 48.6536425, 6.824211389},
	{"MF22164", "Avricourt", 48.65383806, 6.82472},
	{"MF22172", "Avricourt", 48.6529675, 6.822763611},
	{"MF22140", "Avricourt", 48.65261389, 6.8225325},
	{"MF22174", "Avricourt", 48.65325944, 6.821715278},
	{"MF22176", "Avricourt", 48.65387972, 6.820951667},
	{"MF22182", "Avricourt", 48.65311222, 6.820315556},
	{"MF22183", "Avricourt", 48.65285556, 6.820610278},
	{"MF22186", "Avricourt", 48.65286889, 6.819963333},
	{"MF22185", "Avricourt", 48.65276, 6.819709444},
	{"MF22184", "Avricourt", 48.65248556, 6.819389167},
	{"MF22129", "Moussey", 48.67294167, 6.779651389},
	{"MF22005", "Moussey", 48.67283361, 6.779289167},
	{"MF220M4", "Moussey", 48.67271083, 6.77881},
	{"MF22003", "Moussey", 48.672515, 6.778281944},
	{"MF22007", "Moussey", 48.67306472, 6.780090278},
	{"MF22008", "Moussey", 48.67324722, 6.7805625},
	{"MF22017", "Moussey", 48.67555806, 6.779148056},
	{"MF22104", "Moussey", 48.67394167, 6.783895278},
	{"MF22120", "Moussey", 48.67238139, 6.785153333},
	{"MF22121", "Moussey", 48.67193306, 6.785250278},
	{"MF22124", "Moussey", 48.67109722, 6.785146389},
	{"MF22125", "Moussey", 48.67066889, 6.785198056},
	{"MF22994", "Moussey", 48.68865667, 6.802375278},
	{"MF22028", "Moussey", 48.68827694, 6.801550556},
	{"MF22219", "Moussey", 48.68520389, 6.795860833},
	{"MF22222", "Moussey", 48.68427167, 6.795415278},
	{"MF22238", "Moussey", 48.68351944, 6.797263889},
	{"MF22237", "Moussey", 48.68329583, 6.797425556},
	{"MF22245", "Moussey", 48.68313556, 6.795824167},
	{"MF22246", "Moussey", 48.68282056, 6.796250556},
	{"MF22243", "Moussey", 48.68381833, 6.794796111},
	{"MF22256", "Moussey", 48.68173278, 6.792783333},
	{"MF22260", "Moussey", 48.6807475, 6.791523056},
	{"MF22268", "Moussey", 48.68072611, 6.789259722},
	{"MF22272", "Moussey", 48.67860667, 6.789581389},
	//{"MF22277", "Rechicourt-le-Chateau", 48.66489083, 6.843915},
	//{"MF22296", "Rechicourt-le-Chateau", 48.66392694, 6.842715833},
	//{"MF22299", "Rechicourt-le-Chateau", 48.66342833, 6.841445833},
	//{"MF22302", "Rechicourt-le-Chateau", 48.66304389, 6.841816111},
	//{"MF22322", "Rechicourt-le-Chateau", 48.66855556, 6.8401975},
	//{"MF22323", "Rechicourt-le-Chateau", 48.66892306, 6.840602778},
	//{"MF22324", "Rechicourt-le-Chateau", 48.66932472, 6.840821389},
	//{"MF22341", "Rechicourt-le-Chateau", 48.66474306, 6.839362222},
	//{"MF22346", "Rechicourt-le-Chateau", 48.66380167, 6.839710833},
	//{"MF22354", "Rechicourt-le-Chateau", 48.66723361, 6.832873333},
	//{"MF22513", "Hertzing", 48.68857833, 6.953845},
	//{"MF22521", "Hertzing", 48.68738167, 6.950633056},
	//{"MF22522", "Hertzing", 48.68699444, 6.950835833},
	//{"MF22524", "Hertzing", 48.68725056, 6.949720278},
	//{"MF22535", "Gondrexange", 48.687005, 6.927591944},
	//{"MF22536", "Gondrexange", 48.68666417, 6.927485833},
	//{"MF22538", "Gondrexange", 48.6863075, 6.927270278},
	//{"MF22541", "Gondrexange", 48.68560194, 6.927248611},
	//{"MF22542", "Gondrexange", 48.68531028, 6.927771667},
	//{"MF22543", "Gondrexange", 48.68506361, 6.9282175},
	//{"MF22544", "Gondrexange", 48.68537389, 6.928441389},
	//{"MF22546", "Gondrexange", 48.684665, 6.928902778},
	//{"MF22573", "Gondrexange", 48.68440333, 6.928902778},
	//{"MF22556", "Gondrexange", 48.68628194, 6.928623889},
	//{"MF22548", "Gondrexange", 48.68513944, 6.929541389},
	//{"MF22562", "Gondrexange", 48.68540139, 6.926319167},
	//{"MF22563", "Gondrexange", 48.68506778, 6.9259075},
	//{"MF22567", "Gondrexange", 48.68434389, 6.927116111},
	//{"MF22568", "Gondrexange", 48.68416417, 6.927418056},
	//{"MF22579", "Gondrexange", 48.68280611, 6.930495556},
	//{"MF22581", "Gondrexange", 48.68263333, 6.931374167},
	//{"MF22584", "Gondrexange", 48.68332444, 6.932414444},
	//{"MF22593", "Gondrexange", 48.68350444, 6.927856111},
	//{"MF22594", "Gondrexange", 48.68322778, 6.927503333},
	//{"MF22596", "Gondrexange", 48.68279333, 6.926906111},
	//{"MF22597", "Gondrexange", 48.68250167, 6.926476111},
	//{"MF22613", "Gondrexange", 48.68443222, 6.925142222},
	//{"MF22620", "Gondrexange", 48.68446944, 6.921548333},
	//{"MF22625", "Gondrexange", 48.68345611, 6.919533333},
	//{"MF22641", "Gondrexange", 48.68795028, 6.931528333},
	{"MF22652", "Saint-Georges", 48.65994361, 6.934416389},
	{"MF22661", "Saint-Georges", 48.6591275, 6.931368889},
	{"MF22666", "Saint-Georges", 48.65823833, 6.929964167},
	{"MF22674", "Saint-Georges", 48.65759944, 6.929021111},
	{"MF22677", "Saint-Georges", 48.65705556, 6.927816111},
	{"MF22686", "Saint-Georges", 48.65872333, 6.928489444},
	{"MF22693", "Saint-Georges", 48.65764167, 6.926398333},
	{"MF22694", "Saint-Georges", 48.65752306, 6.927009444},
	{"MF22696", "Saint-Georges", 48.657965, 6.926471667},
	{"MF22699", "Saint-Georges", 48.65889694, 6.9254575},
	{"MF22701", "Saint-Georges", 48.65896944, 6.924851111},
	{"MF22932", "Richeval", 48.63734, 6.911486111},
	{"MF22714", "Richeval", 48.63694722, 6.910884722},
	{"MF22725", "Richeval", 48.63604889, 6.911171111},
	{"MF22739", "Richeval", 48.63620528, 6.910274444},
	{"MF22757", "Ibigny", 48.64696389, 6.891068056},
	{"MF22785", "Ibigny", 48.64311778, 6.901533056},
	{"MF22783", "Ibigny", 48.64348861, 6.900448056},
	{"MF22781", "Ibigny", 48.64385889, 6.900287778},
	{"MF22774", "Ibigny", 48.64335639, 6.899443333},
	{"MF22777", "Ibigny", 48.64257444, 6.898117222},
	{"MF22778", "Ibigny", 48.64218083, 6.898291667},
	{"MF22975", "Ibigny", 48.64210667, 6.902787222},
	{"MF22799", "Foulcrey", 48.63497639, 6.858034444},
	{"MF22798", "Foulcrey", 48.6346, 6.858000556},
	{"MF22991", "Foulcrey", 48.63615, 6.857541389},
	{"MF22804", "Foulcrey", 48.63585389, 6.857374444},
	{"MF22807", "Foulcrey", 48.63648222, 6.857576111},
	{"MF22820", "Foulcrey", 48.63675861, 6.856575556},
	{"MF22809", "Foulcrey", 48.63703389, 6.857435556},
	{"MF22810", "Foulcrey", 48.63730139, 6.857881944},
	{"MF22811", "Foulcrey", 48.63760806, 6.858388611},
	{"MF22821", "Foulcrey", 48.63680278, 6.855143333},
	{"MF22823", "Foulcrey", 48.63698667, 6.855366111},
	{"MF22833", "Foulcrey", 48.63788083, 6.858008889},
	{"MF22834", "Foulcrey", 48.63811194, 6.857729722},
	{"MF22839", "Foulcrey", 48.63761444, 6.857043611},
	{"MF22840", "Foulcrey", 48.63835083, 6.858033611},
	{"MF22836", "Foulcrey", 48.63847917, 6.856597778},
	{"MF22830", "Foulcrey", 48.63762806, 6.856010278},
	{"MF22829", "Foulcrey", 48.63733611, 6.855866667},
	{"MF22843", "Foulcrey", 48.63872944, 6.858158333},
	{"MF22848", "Foulcrey", 48.64008167, 6.857115556},
	{"MF22817", "Foulcrey", 48.63655111, 6.858186111},
	{"MF22835", "Foulcrey", 48.63834278, 6.857174444},
	{"MF22851", "Foulcrey", 48.64088278, 6.855696944},
	//{"MF22100", "Maizière-lès-Vic", 48.69961444, 6.800100556},
	//{"MF22209", "Rechicourt-le-Chateau", 48.67387333, 6.8416275},
	//{"MF22205", "Rechicourt-le-Chateau", 48.67231417, 6.840896944},
	//{"MF22213", "Rechicourt-le-Chateau", 48.67546472, 6.842726667},
	//{"MF22211", "Rechicourt-le-Chateau", 48.67465972, 6.842134722},
	//{"MF22998", "Rechicourt-le-Chateau", 48.68269694, 6.853323889},
	//{"MF22304", "Rechicourt-le-Chateau", 48.69258194, 6.847271667},
	//{"MF22258", "Rechicourt-le-Chateau", 48.6658025, 6.840346111},
	//{"MF22787", "Saint-Georges", 48.6571475, 6.924421389},
	//{"MF22784", "Saint-Georges", 48.6567925, 6.923207778},
	//{"MF22786", "Saint-Georges", 48.656955, 6.923806389},
	//{"MF22996", "Saint-Georges", 48.65733139, 6.925023889},
	{"MF22806", "Saint-Georges", 48.65256194, 6.930313611},
	{"MF22805", "Saint-Georges", 48.65287778, 6.929853056},
	{"MF22801", "Saint-Georges", 48.655285, 6.928313611},
	{"MF22819", "Saint-Georges", 48.64801694, 6.934151111},
	{"MF22824", "Saint-Georges", 48.64653111, 6.935343056},
	{"MF22815", "Saint-Georges", 48.6501275, 6.933261944},
	{"MF22650", "Saint-Georges", 48.65859194, 6.930193056},
	{"MF22651", "Saint-Georges", 48.65868333, 6.9303275},
	{"MF22997", "Foulcrey", 48.63688917, 6.855005556},
	{"MF22667", "Avricourt", 48.64926972, 6.817091111},
	{"MF22663", "Avricourt", 48.65011444, 6.812752778},
	{"MF22672", "Avricourt", 48.64882722, 6.804091111},
	{"MF22999", "Avricourt", 48.64946861, 6.814881389},
	{"MF22671", "Avricourt", 48.64911028, 6.804318889},
	{"MF22990", "Moussey", 48.68475333, 6.803333611},
	//{"MF22841", "Rechicourt-le-Chateau", 48.68892861, 6.803174722},
	//{"MF22001", "Rechicourt-le-Chateau", 48.6891575, 6.803713611},
	//{"MF22869", "Rechicourt-le-Chateau", 48.69097972, 6.80856},
	//{"MF22837", "Rechicourt-le-Chateau", 48.69030694, 6.806426944},
	{"MF22002", "Moussey", 48.69043278, 6.805981667},
	//{"MF22264", "Rechicourt-le-Chateau", 48.6807325, 6.818073333},
	//{"MF22306", "Rechicourt-le-Chateau", 48.66354556, 6.843489444},
	{"MF22773", "Ibigny", 48.64359639, 6.899089722},
	//{"MF22303", "Rechicourt-le-Chateau", 48.66299833, 6.842009167},
	//{"MF22305", "Rechicourt-le-Chateau", 48.66335889, 6.843009722},
	{"MF22034", "Avricourt", 48.65093722, 6.804666944},
	{"MF22004", "Avricourt", 48.64818139, 6.805636389},
	{"MF22136", "Avricourt", 48.64919333, 6.816627778},
	{"MF22192", "Avricourt", 48.65418361, 6.822214444},
	{"MF22958", "Richeval", 48.63973056, 6.910341389},
	{"MF22959", "Richeval", 48.64005111, 6.909957222},
	{"MF22960", "Richeval", 48.64028972, 6.909481667},
	{"MF22961", "Richeval", 48.64042611, 6.908951944},
	{"MF22962", "Richeval", 48.64065528, 6.908526944},
	{"MF22963", "Richeval", 48.64093639, 6.908177778},
	{"MF22964", "Richeval", 48.64118083, 6.907893056},
	{"MF22965", "Richeval", 48.64107472, 6.907406944},
	{"MF22966", "Ibigny", 48.64116639, 6.907057222},
	{"MF22967", "Ibigny", 48.64101278, 6.90657},
	{"MF22969", "Ibigny", 48.64072833, 6.905538056},
	{"MF22970", "Ibigny", 48.64081361, 6.905018333},
	{"MF22971", "Ibigny", 48.64092444, 6.9044975},
	{"MF22973", "Ibigny", 48.64141306, 6.903634722},
	{"MF22974", "Ibigny", 48.64173778, 6.903209722},
	//{"MF22956", "Richeval", 48.64443222, 6.919154167},
	//{"MF22955", "Richeval", 48.64411139, 6.918923611},
	//{"MF22954", "Richeval", 48.64378778, 6.918676944},
	//{"MF22953", "Richeval", 48.6435025, 6.918427778},
	//{"MF22951", "Richeval", 48.64289028, 6.917888611},
	//{"MF22950", "Richeval", 48.64257444, 6.917602778},
	//{"MF22949", "Richeval", 48.6422725, 6.917327778},
	//{"MF22948", "Richeval", 48.64196472, 6.917045278},
	//{"MF22947", "Richeval", 48.64164306, 6.916759444},
	//{"MF22942", "Richeval", 48.640105, 6.915011667},
	//{"MF22945", "Richeval", 48.64099778, 6.916133056},
	//{"MF22944", "Richeval", 48.6406975, 6.91576},
	//{"MF22940", "Richeval", 48.63953167, 6.914283611},
	//{"MF22938", "Richeval", 48.63895417, 6.913553333},
	//{"MF22936", "Richeval", 48.63838139, 6.912816111},
	//{"MF22935", "Richeval", 48.63809306, 6.912451389},
	//{"MF22933", "Richeval", 48.63756861, 6.911790833},
	//{"MF22934", "Richeval", 48.63781361, 6.912103611},
	//{"MF22931", "Richeval", 48.63334444, 6.908256111},
	//{"MF22929", "Richeval", 48.63277361, 6.907510556},
	//{"MF22927", "Richeval", 48.63218361, 6.906779444},
	//{"MF22926", "Richeval", 48.63189639, 6.906417222},
	//{"MF22924", "Richeval", 48.63133, 6.905671111},
	//{"MF22923", "Richeval", 48.63106028, 6.905323611},
	//{"MF22922", "Richeval", 48.63079667, 6.904993611},
	//{"MF22919", "Richeval", 48.62998194, 6.9039325},
	//{"MF22918", "Richeval", 48.62972389, 6.903574167},
	//{"MF22917", "Richeval", 48.62943389, 6.903196111},
	//{"MF22916", "Richeval", 48.62914, 6.902807778},
	//{"MF22915", "Richeval", 48.62887444, 6.902479722},
	//{"MF22913", "Richeval", 48.62845722, 6.90194},
	//{"MF22911", "Richeval", 48.62796528, 6.901314722},
	//{"MF22910", "Richeval", 48.62767167, 6.9009575},
	//{"MF22909", "Richeval", 48.6273675, 6.900603611},
	//{"MF22908", "Richeval", 48.62707083, 6.900265833},
	//{"MF22907", "Richeval", 48.62677278, 6.899941667},
	//{"MF22905", "Richeval", 48.62617, 6.899272222},
	//{"MF22904", "Richeval", 48.62586167, 6.898933611},
	//{"MF22903", "Richeval", 48.62556056, 6.898591944},
	//{"MF22902", "Richeval", 48.62526611, 6.898253611},
	//{"MF22901", "Richeval", 48.62507917, 6.898072222},
	//{"MF22930", "Richeval", 48.63305694, 6.907878611},
	//{"MF22906", "Richeval", 48.6264725, 6.899606944},
	//{"MF22912", "Richeval", 48.62822806, 6.901636944},
	//{"MF22914", "Richeval", 48.62865583, 6.902205833},
	//{"MF22928", "Richeval", 48.63247111, 6.907147778},
	//{"MF22925", "Richeval", 48.63160917, 6.906048056},
	//{"MF22920", "Richeval", 48.63025806, 6.904292222},
	//{"MF22921", "Richeval", 48.63053528, 6.904652778},
	//{"MF22952", "Richeval", 48.6431975, 6.918157778},
	//{"MF22937", "Richeval", 48.63867028, 6.913183056},
	//{"MF22939", "Richeval", 48.63924306, 6.913918611},
	//{"MF22941", "Richeval", 48.63981944, 6.914650278},
	//{"MF22946", "Richeval", 48.64131556, 6.916463611},
	//{"MF22943", "Richeval", 48.6403975, 6.91539},
	{"MF22968", "Ibigny", 48.64085528, 6.906035833},
	{"MF22972", "Ibigny", 48.64109333, 6.904035833},
	{"MF22957", "Richeval", 48.63936028, 6.910701111},
}
