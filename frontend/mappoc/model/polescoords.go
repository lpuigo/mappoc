package model

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/lpuig/ewin/doe/website/frontend/tools"
	"github.com/lpuig/ewin/doe/website/frontend/tools/elements"
)

type BePole struct {
	Ref   string
	City  string
	Lat   float64
	Long  float64
	State string
}

const (
	PoleStateNotSubmitted string = "00 Not Submitted"
	PoleStateToDo         string = "10 To Do"
	PoleStateHoleDone     string = "20 Hole Done"
	PoleStateIncident     string = "25 Incident"
	PoleStateDone         string = "90 Done"
	PoleStateCancelled    string = "99 Cancelled"
)

func GetStatesValueLabel() []*elements.ValueLabel {
	return []*elements.ValueLabel{
		elements.NewValueLabel(PoleStateNotSubmitted, "Non soumis"),
		elements.NewValueLabel(PoleStateToDo, "A faire"),
		elements.NewValueLabel(PoleStateHoleDone, "Trou fait"),
		elements.NewValueLabel(PoleStateIncident, "Incident"),
		elements.NewValueLabel(PoleStateDone, "Fait"),
		elements.NewValueLabel(PoleStateCancelled, "Annulé"),
	}
}

type Pole struct {
	*js.Object
	Ref   string  `js:"Ref"`
	City  string  `js:"City"`
	Lat   float64 `js:"Lat"`
	Long  float64 `js:"Long"`
	State string  `js:"State"`
}

func NewPole(pole BePole) *Pole {
	np := &Pole{
		Object: tools.O(),
	}

	np.Ref = pole.Ref
	np.City = pole.City
	np.Lat = pole.Lat
	np.Long = pole.Long
	np.State = pole.State

	return np
}

func (p *Pole) SwitchState() {
	switch p.State {
	case PoleStateNotSubmitted:
		p.State = PoleStateToDo
	case PoleStateToDo:
		p.State = PoleStateHoleDone
	case PoleStateHoleDone:
		p.State = PoleStateIncident
	case PoleStateIncident:
		p.State = PoleStateDone
	case PoleStateDone:
		p.State = PoleStateCancelled
	case PoleStateCancelled:
		p.State = PoleStateNotSubmitted
	}
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

var Poles = []BePole{
	//{"MF22A01", "Maizières-lès-Vic", 48.68398083, 6.767729722,PoleStateToDo},
	//{"MF22A02", "Maizières-lès-Vic", 48.6867125, 6.770223056,PoleStateToDo},
	//{"MF22A03", "Maizières-lès-Vic", 48.69282139, 6.762173611,PoleStateToDo},
	//{"MF22A04", "Maizières-lès-Vic", 48.68570194, 6.778500278,PoleStateToDo},
	//{"MF22A05", "Maizières-lès-Vic", 48.68481806, 6.779818333,PoleStateToDo},
	//{"MF22A06", "Maizières-lès-Vic", 48.68511778, 6.7796225,PoleStateToDo},
	{"MF220M4", "Moussey", 48.67271083, 6.77881, PoleStateDone},
	//{"MF22001", "Rechicourt-le-Chateau", 48.6891575, 6.803713611,PoleStateToDo},
	{"MF22002", "Moussey", 48.69043278, 6.805981667, PoleStateDone},
	{"MF22003", "Moussey", 48.672515, 6.778281944, PoleStateToDo},
	{"MF22004", "Avricourt", 48.64818139, 6.805636389, PoleStateDone},
	{"MF22005", "Moussey", 48.67283361, 6.779289167, PoleStateDone},
	{"MF22006", "Maizière-lès-Vic", 48.70158583, 6.797752778, PoleStateDone},
	{"MF22007", "Moussey", 48.67306472, 6.780090278, PoleStateDone},
	{"MF22008", "Moussey", 48.67324722, 6.7805625, PoleStateDone},
	{"MF22009", "Avricourt", 48.65007417, 6.805681111, PoleStateDone},
	{"MF22011", "Avricourt", 48.65011278, 6.805285833, PoleStateToDo},
	{"MF22017", "Moussey", 48.67555806, 6.779148056, PoleStateToDo},
	{"MF22019", "Avricourt", 48.65126222, 6.806866667, PoleStateDone},
	{"MF22023", "Avricourt", 48.65164306, 6.805110278, PoleStateDone},
	{"MF22026", "Avricourt", 48.65183056, 6.803511111, PoleStateDone},
	{"MF22027", "Avricourt", 48.65202222, 6.80309, PoleStateDone},
	{"MF22028", "Moussey", 48.68827694, 6.801550556, PoleStateDone},
	{"MF22031", "Avricourt", 48.65296444, 6.800733611, PoleStateDone},
	{"MF22034", "Avricourt", 48.65093722, 6.804666944, PoleStateDone},
	{"MF22036", "Avricourt", 48.65221222, 6.805172222, PoleStateDone},
	{"MF22039", "Avricourt", 48.65268361, 6.804828611, PoleStateDone},
	{"MF22043", "Avricourt", 48.65489806, 6.804246667, PoleStateDone},
	{"MF22053", "Avricourt", 48.64927167, 6.806758611, PoleStateDone},
	{"MF22054", "Avricourt", 48.64892583, 6.806525278, PoleStateDone},
	{"MF22062", "Avricourt", 48.64828, 6.810688056, PoleStateDone},
	{"MF22064", "Avricourt", 48.6487175, 6.811341389, PoleStateDone},
	{"MF22065", "Avricourt", 48.64879778, 6.811813333, PoleStateDone},
	{"MF22070", "Avricourt", 48.64959917, 6.811748889, PoleStateDone},
	{"MF22072", "Avricourt", 48.65029611, 6.811665556, PoleStateDone},
	{"MF22077", "Avricourt", 48.65082528, 6.809250833, PoleStateDone},
	{"MF22078", "Avricourt", 48.65110833, 6.809428056, PoleStateDone},
	//{"MF22085", "Avricourt", 48.65008667, 6.808467778,PoleStateToDo}, // Not allocated
	{"MF22086", "Avricourt", 48.64994, 6.808153889, PoleStateDone},
	{"MF22090", "Avricourt", 48.65129944, 6.807456389, PoleStateDone},
	{"MF22091", "Avricourt", 48.65164417, 6.807748333, PoleStateDone},
	//{"MF22100", "Maizière-lès-Vic", 48.69961444, 6.800100556,PoleStateToDo},
	{"MF22103", "Avricourt", 48.65231528, 6.808456389, PoleStateDone},
	{"MF22104", "Moussey", 48.67394167, 6.783895278, PoleStateToDo},
	{"MF22106", "Avricourt", 48.65023917, 6.812333056, PoleStateDone},
	{"MF22120", "Moussey", 48.67238139, 6.785153333, PoleStateDone},
	{"MF22121", "Moussey", 48.67193306, 6.785250278, PoleStateDone},
	{"MF22124", "Moussey", 48.67109722, 6.785146389, PoleStateDone},
	{"MF22125", "Moussey", 48.67066889, 6.785198056, PoleStateDone},
	{"MF22129", "Moussey", 48.67294167, 6.779651389, PoleStateDone},
	{"MF22136", "Avricourt", 48.64919333, 6.816627778, PoleStateToDo},
	{"MF22140", "Avricourt", 48.65261389, 6.8225325, PoleStateToDo},
	{"MF22141", "Avricourt", 48.65228833, 6.822251389, PoleStateToDo},
	{"MF22143", "Avricourt", 48.65203167, 6.822015278, PoleStateToDo},
	{"MF22146", "Avricourt", 48.65239083, 6.820795833, PoleStateDone},
	{"MF22148", "Avricourt", 48.65144694, 6.821033889, PoleStateToDo},
	{"MF22154", "Avricourt", 48.65195167, 6.819122778, PoleStateDone},
	{"MF22156", "Avricourt", 48.65060167, 6.819590278, PoleStateDone},
	{"MF22158", "Avricourt", 48.65102222, 6.818688611, PoleStateDone},
	{"MF22162", "Avricourt", 48.65342389, 6.823642222, PoleStateDone},
	{"MF22163", "Avricourt", 48.6536425, 6.824211389, PoleStateDone},
	{"MF22164", "Avricourt", 48.65383806, 6.82472, PoleStateDone},
	{"MF22172", "Avricourt", 48.6529675, 6.822763611, PoleStateDone},
	{"MF22174", "Avricourt", 48.65325944, 6.821715278, PoleStateDone},
	{"MF22176", "Avricourt", 48.65387972, 6.820951667, PoleStateDone},
	{"MF22182", "Avricourt", 48.65311222, 6.820315556, PoleStateDone},
	{"MF22183", "Avricourt", 48.65285556, 6.820610278, PoleStateDone},
	{"MF22184", "Avricourt", 48.65248556, 6.819389167, PoleStateDone},
	{"MF22185", "Avricourt", 48.65276, 6.819709444, PoleStateDone},
	{"MF22186", "Avricourt", 48.65286889, 6.819963333, PoleStateToDo},
	{"MF22187", "Avricourt", 48.6496425, 6.806759444, PoleStateDone},
	{"MF22192", "Avricourt", 48.65418361, 6.822214444, PoleStateDone},
	//{"MF22205", "Rechicourt-le-Chateau", 48.67231417, 6.840896944,PoleStateToDo},
	//{"MF22209", "Rechicourt-le-Chateau", 48.67387333, 6.8416275,PoleStateToDo},
	//{"MF22211", "Rechicourt-le-Chateau", 48.67465972, 6.842134722,PoleStateToDo},
	//{"MF22213", "Rechicourt-le-Chateau", 48.67546472, 6.842726667,PoleStateToDo},
	{"MF22219", "Moussey", 48.68520389, 6.795860833, PoleStateToDo},
	{"MF22222", "Moussey", 48.68427167, 6.795415278, PoleStateDone},
	{"MF22237", "Moussey", 48.68329583, 6.797425556, PoleStateToDo},
	{"MF22238", "Moussey", 48.68351944, 6.797263889, PoleStateToDo},
	{"MF22243", "Moussey", 48.68381833, 6.794796111, PoleStateDone},
	{"MF22245", "Moussey", 48.68313556, 6.795824167, PoleStateDone},
	{"MF22246", "Moussey", 48.68282056, 6.796250556, PoleStateDone},
	{"MF22256", "Moussey", 48.68173278, 6.792783333, PoleStateDone},
	//{"MF22258", "Rechicourt-le-Chateau", 48.6658025, 6.840346111,PoleStateToDo},
	{"MF22260", "Moussey", 48.6807475, 6.791523056, PoleStateDone},
	//{"MF22264", "Rechicourt-le-Chateau", 48.6807325, 6.818073333,PoleStateToDo},
	{"MF22268", "Moussey", 48.68072611, 6.789259722, PoleStateToDo},
	{"MF22272", "Moussey", 48.67860667, 6.789581389, PoleStateToDo},
	//{"MF22277", "Rechicourt-le-Chateau", 48.66489083, 6.843915,PoleStateToDo},
	//{"MF22296", "Rechicourt-le-Chateau", 48.66392694, 6.842715833,PoleStateToDo},
	//{"MF22299", "Rechicourt-le-Chateau", 48.66342833, 6.841445833,PoleStateToDo},
	//{"MF22302", "Rechicourt-le-Chateau", 48.66304389, 6.841816111,PoleStateToDo},
	//{"MF22303", "Rechicourt-le-Chateau", 48.66299833, 6.842009167,PoleStateToDo},
	//{"MF22304", "Rechicourt-le-Chateau", 48.69258194, 6.847271667,PoleStateToDo},
	//{"MF22305", "Rechicourt-le-Chateau", 48.66335889, 6.843009722,PoleStateToDo},
	//{"MF22306", "Rechicourt-le-Chateau", 48.66354556, 6.843489444,PoleStateToDo},
	//{"MF22322", "Rechicourt-le-Chateau", 48.66855556, 6.8401975,PoleStateToDo},
	//{"MF22323", "Rechicourt-le-Chateau", 48.66892306, 6.840602778,PoleStateToDo},
	//{"MF22324", "Rechicourt-le-Chateau", 48.66932472, 6.840821389,PoleStateToDo},
	//{"MF22341", "Rechicourt-le-Chateau", 48.66474306, 6.839362222,PoleStateToDo},
	//{"MF22346", "Rechicourt-le-Chateau", 48.66380167, 6.839710833,PoleStateToDo},
	//{"MF22354", "Rechicourt-le-Chateau", 48.66723361, 6.832873333,PoleStateToDo},
	//{"MF22513", "Hertzing", 48.68857833, 6.953845,PoleStateToDo},
	//{"MF22521", "Hertzing", 48.68738167, 6.950633056,PoleStateToDo},
	//{"MF22522", "Hertzing", 48.68699444, 6.950835833,PoleStateToDo},
	//{"MF22524", "Hertzing", 48.68725056, 6.949720278,PoleStateToDo},
	//{"MF22535", "Gondrexange", 48.687005, 6.927591944,PoleStateToDo},
	//{"MF22536", "Gondrexange", 48.68666417, 6.927485833,PoleStateToDo},
	//{"MF22538", "Gondrexange", 48.6863075, 6.927270278,PoleStateToDo},
	//{"MF22541", "Gondrexange", 48.68560194, 6.927248611,PoleStateToDo},
	//{"MF22542", "Gondrexange", 48.68531028, 6.927771667,PoleStateToDo},
	//{"MF22543", "Gondrexange", 48.68506361, 6.9282175,PoleStateToDo},
	//{"MF22544", "Gondrexange", 48.68537389, 6.928441389,PoleStateToDo},
	//{"MF22546", "Gondrexange", 48.684665, 6.928902778,PoleStateToDo},
	//{"MF22548", "Gondrexange", 48.68513944, 6.929541389,PoleStateToDo},
	//{"MF22556", "Gondrexange", 48.68628194, 6.928623889,PoleStateToDo},
	//{"MF22562", "Gondrexange", 48.68540139, 6.926319167,PoleStateToDo},
	//{"MF22563", "Gondrexange", 48.68506778, 6.9259075,PoleStateToDo},
	//{"MF22567", "Gondrexange", 48.68434389, 6.927116111,PoleStateToDo},
	//{"MF22568", "Gondrexange", 48.68416417, 6.927418056,PoleStateToDo},
	//{"MF22573", "Gondrexange", 48.68440333, 6.928902778,PoleStateToDo},
	//{"MF22579", "Gondrexange", 48.68280611, 6.930495556,PoleStateToDo},
	//{"MF22581", "Gondrexange", 48.68263333, 6.931374167,PoleStateToDo},
	//{"MF22584", "Gondrexange", 48.68332444, 6.932414444,PoleStateToDo},
	//{"MF22593", "Gondrexange", 48.68350444, 6.927856111,PoleStateToDo},
	//{"MF22594", "Gondrexange", 48.68322778, 6.927503333,PoleStateToDo},
	//{"MF22596", "Gondrexange", 48.68279333, 6.926906111,PoleStateToDo},
	//{"MF22597", "Gondrexange", 48.68250167, 6.926476111,PoleStateToDo},
	//{"MF22613", "Gondrexange", 48.68443222, 6.925142222,PoleStateToDo},
	//{"MF22620", "Gondrexange", 48.68446944, 6.921548333,PoleStateToDo},
	//{"MF22625", "Gondrexange", 48.68345611, 6.919533333,PoleStateToDo},
	//{"MF22641", "Gondrexange", 48.68795028, 6.931528333,PoleStateToDo},
	{"MF22650", "Saint-Georges", 48.65859194, 6.930193056, PoleStateNotSubmitted},
	{"MF22651", "Saint-Georges", 48.65868333, 6.9303275, PoleStateNotSubmitted},
	{"MF22652", "Saint-Georges", 48.65994361, 6.934416389, PoleStateNotSubmitted},
	{"MF22661", "Saint-Georges", 48.6591275, 6.931368889, PoleStateNotSubmitted},
	{"MF22663", "Avricourt", 48.65011444, 6.812752778, PoleStateToDo},
	{"MF22666", "Saint-Georges", 48.65823833, 6.929964167, PoleStateNotSubmitted},
	{"MF22667", "Avricourt", 48.64926972, 6.817091111, PoleStateDone},
	{"MF22671", "Avricourt", 48.64911028, 6.804318889, PoleStateDone},
	{"MF22672", "Avricourt", 48.64882722, 6.804091111, PoleStateDone},
	{"MF22674", "Saint-Georges", 48.65759944, 6.929021111, PoleStateNotSubmitted},
	{"MF22677", "Saint-Georges", 48.65705556, 6.927816111, PoleStateNotSubmitted},
	{"MF22686", "Saint-Georges", 48.65872333, 6.928489444, PoleStateNotSubmitted},
	{"MF22693", "Saint-Georges", 48.65764167, 6.926398333, PoleStateNotSubmitted},
	{"MF22694", "Saint-Georges", 48.65752306, 6.927009444, PoleStateNotSubmitted},
	{"MF22696", "Saint-Georges", 48.657965, 6.926471667, PoleStateNotSubmitted},
	{"MF22699", "Saint-Georges", 48.65889694, 6.9254575, PoleStateNotSubmitted},
	{"MF22701", "Saint-Georges", 48.65896944, 6.924851111, PoleStateNotSubmitted},
	{"MF22714", "Richeval", 48.63694722, 6.910884722, PoleStateNotSubmitted},
	{"MF22725", "Richeval", 48.63604889, 6.911171111, PoleStateNotSubmitted},
	{"MF22739", "Richeval", 48.63620528, 6.910274444, PoleStateNotSubmitted},
	{"MF22757", "Ibigny", 48.64696389, 6.891068056, PoleStateDone},
	{"MF22773", "Ibigny", 48.64359639, 6.899089722, PoleStateToDo},
	{"MF22774", "Ibigny", 48.64335639, 6.899443333, PoleStateDone},
	{"MF22777", "Ibigny", 48.64257444, 6.898117222, PoleStateDone},
	{"MF22778", "Ibigny", 48.64218083, 6.898291667, PoleStateDone},
	{"MF22781", "Ibigny", 48.64385889, 6.900287778, PoleStateDone},
	{"MF22783", "Ibigny", 48.64348861, 6.900448056, PoleStateDone},
	//{"MF22784", "Saint-Georges", 48.6567925, 6.923207778,PoleStateToDo},
	{"MF22785", "Ibigny", 48.64311778, 6.901533056, PoleStateDone},
	//{"MF22786", "Saint-Georges", 48.656955, 6.923806389,PoleStateToDo},
	//{"MF22787", "Saint-Georges", 48.6571475, 6.924421389,PoleStateToDo},
	{"MF22798", "Foulcrey", 48.6346, 6.858000556, PoleStateDone},
	{"MF22799", "Foulcrey", 48.63497639, 6.858034444, PoleStateDone},
	{"MF22801", "Saint-Georges", 48.655285, 6.928313611, PoleStateNotSubmitted},
	{"MF22804", "Foulcrey", 48.63585389, 6.857374444, PoleStateToDo},
	{"MF22805", "Saint-Georges", 48.65287778, 6.929853056, PoleStateNotSubmitted},
	{"MF22806", "Saint-Georges", 48.65256194, 6.930313611, PoleStateNotSubmitted},
	{"MF22807", "Foulcrey", 48.63648222, 6.857576111, PoleStateDone},
	{"MF22809", "Foulcrey", 48.63703389, 6.857435556, PoleStateDone},
	{"MF22810", "Foulcrey", 48.63730139, 6.857881944, PoleStateDone},
	{"MF22811", "Foulcrey", 48.63760806, 6.858388611, PoleStateDone},
	{"MF22815", "Saint-Georges", 48.6501275, 6.933261944, PoleStateNotSubmitted},
	{"MF22817", "Foulcrey", 48.63655111, 6.858186111, PoleStateDone},
	{"MF22819", "Saint-Georges", 48.64801694, 6.934151111, PoleStateNotSubmitted},
	{"MF22820", "Foulcrey", 48.63675861, 6.856575556, PoleStateDone},
	{"MF22821", "Foulcrey", 48.63680278, 6.855143333, PoleStateDone},
	{"MF22822", "Foulcrey", 48.6376447, 6.8556019, PoleStateDone},
	{"MF22823", "Foulcrey", 48.63698667, 6.855366111, PoleStateDone},
	{"MF22824", "Saint-Georges", 48.64653111, 6.935343056, PoleStateNotSubmitted},
	{"MF22829", "Foulcrey", 48.63733611, 6.855866667, PoleStateToDo},
	{"MF22830", "Foulcrey", 48.63762806, 6.856010278, PoleStateDone},
	{"MF22833", "Foulcrey", 48.63788083, 6.858008889, PoleStateToDo},
	{"MF22834", "Foulcrey", 48.63811194, 6.857729722, PoleStateDone},
	{"MF22835", "Foulcrey", 48.63834278, 6.857174444, PoleStateDone},
	{"MF22836", "Foulcrey", 48.63847917, 6.856597778, PoleStateDone},
	{"MF22837", "Rechicourt-le-Chateau", 48.69030694, 6.806426944, PoleStateDone}, // submitted ?
	{"MF22839", "Foulcrey", 48.63761444, 6.857043611, PoleStateDone},
	{"MF22840", "Foulcrey", 48.63835083, 6.858033611, PoleStateDone},
	{"MF22841", "Rechicourt-le-Chateau", 48.68892861, 6.803174722, PoleStateDone},
	{"MF22843", "Foulcrey", 48.63872944, 6.858158333, PoleStateDone},
	{"MF22848", "Foulcrey", 48.64008167, 6.857115556, PoleStateToDo},
	{"MF22851", "Foulcrey", 48.64088278, 6.855696944, PoleStateToDo},
	{"MF22869", "Rechicourt-le-Chateau", 48.69097972, 6.80856, PoleStateDone}, // submitted ?
	//{"MF22901", "Richeval", 48.62507917, 6.898072222,PoleStateToDo},
	//{"MF22902", "Richeval", 48.62526611, 6.898253611,PoleStateToDo},
	//{"MF22903", "Richeval", 48.62556056, 6.898591944,PoleStateToDo},
	//{"MF22904", "Richeval", 48.62586167, 6.898933611,PoleStateToDo},
	//{"MF22905", "Richeval", 48.62617, 6.899272222,PoleStateToDo},
	//{"MF22906", "Richeval", 48.6264725, 6.899606944,PoleStateToDo},
	//{"MF22907", "Richeval", 48.62677278, 6.899941667,PoleStateToDo},
	//{"MF22908", "Richeval", 48.62707083, 6.900265833,PoleStateToDo},
	//{"MF22909", "Richeval", 48.6273675, 6.900603611,PoleStateToDo},
	//{"MF22910", "Richeval", 48.62767167, 6.9009575,PoleStateToDo},
	//{"MF22911", "Richeval", 48.62796528, 6.901314722,PoleStateToDo},
	//{"MF22912", "Richeval", 48.62822806, 6.901636944,PoleStateToDo},
	//{"MF22913", "Richeval", 48.62845722, 6.90194,PoleStateToDo},
	//{"MF22914", "Richeval", 48.62865583, 6.902205833,PoleStateToDo},
	//{"MF22915", "Richeval", 48.62887444, 6.902479722,PoleStateToDo},
	//{"MF22916", "Richeval", 48.62914, 6.902807778,PoleStateToDo},
	//{"MF22917", "Richeval", 48.62943389, 6.903196111,PoleStateToDo},
	//{"MF22918", "Richeval", 48.62972389, 6.903574167,PoleStateToDo},
	//{"MF22919", "Richeval", 48.62998194, 6.9039325,PoleStateToDo},
	//{"MF22920", "Richeval", 48.63025806, 6.904292222,PoleStateToDo},
	//{"MF22921", "Richeval", 48.63053528, 6.904652778,PoleStateToDo},
	//{"MF22922", "Richeval", 48.63079667, 6.904993611,PoleStateToDo},
	//{"MF22923", "Richeval", 48.63106028, 6.905323611,PoleStateToDo},
	//{"MF22924", "Richeval", 48.63133, 6.905671111,PoleStateToDo},
	//{"MF22925", "Richeval", 48.63160917, 6.906048056,PoleStateToDo},
	//{"MF22926", "Richeval", 48.63189639, 6.906417222,PoleStateToDo},
	//{"MF22927", "Richeval", 48.63218361, 6.906779444,PoleStateToDo},
	//{"MF22928", "Richeval", 48.63247111, 6.907147778,PoleStateToDo},
	//{"MF22929", "Richeval", 48.63277361, 6.907510556,PoleStateToDo},
	//{"MF22930", "Richeval", 48.63305694, 6.907878611,PoleStateToDo},
	//{"MF22931", "Richeval", 48.63334444, 6.908256111,PoleStateToDo},
	{"MF22932", "Richeval", 48.63734, 6.911486111, PoleStateNotSubmitted},
	//{"MF22933", "Richeval", 48.63756861, 6.911790833,PoleStateToDo},
	//{"MF22934", "Richeval", 48.63781361, 6.912103611,PoleStateToDo},
	//{"MF22935", "Richeval", 48.63809306, 6.912451389,PoleStateToDo},
	//{"MF22936", "Richeval", 48.63838139, 6.912816111,PoleStateToDo},
	//{"MF22937", "Richeval", 48.63867028, 6.913183056,PoleStateToDo},
	//{"MF22938", "Richeval", 48.63895417, 6.913553333,PoleStateToDo},
	//{"MF22939", "Richeval", 48.63924306, 6.913918611,PoleStateToDo},
	//{"MF22940", "Richeval", 48.63953167, 6.914283611,PoleStateToDo},
	//{"MF22941", "Richeval", 48.63981944, 6.914650278,PoleStateToDo},
	//{"MF22942", "Richeval", 48.640105, 6.915011667,PoleStateToDo},
	//{"MF22943", "Richeval", 48.6403975, 6.91539,PoleStateToDo},
	//{"MF22944", "Richeval", 48.6406975, 6.91576,PoleStateToDo},
	//{"MF22945", "Richeval", 48.64099778, 6.916133056,PoleStateToDo},
	//{"MF22946", "Richeval", 48.64131556, 6.916463611,PoleStateToDo},
	//{"MF22947", "Richeval", 48.64164306, 6.916759444,PoleStateToDo},
	//{"MF22948", "Richeval", 48.64196472, 6.917045278,PoleStateToDo},
	//{"MF22949", "Richeval", 48.6422725, 6.917327778,PoleStateToDo},
	//{"MF22950", "Richeval", 48.64257444, 6.917602778,PoleStateToDo},
	//{"MF22951", "Richeval", 48.64289028, 6.917888611,PoleStateToDo},
	//{"MF22952", "Richeval", 48.6431975, 6.918157778,PoleStateToDo},
	//{"MF22953", "Richeval", 48.6435025, 6.918427778,PoleStateToDo},
	//{"MF22954", "Richeval", 48.64378778, 6.918676944,PoleStateToDo},
	//{"MF22955", "Richeval", 48.64411139, 6.918923611,PoleStateToDo},
	//{"MF22956", "Richeval", 48.64443222, 6.919154167,PoleStateToDo},
	{"MF22957", "Richeval", 48.63936028, 6.910701111, PoleStateNotSubmitted},
	{"MF22958", "Richeval", 48.63973056, 6.910341389, PoleStateNotSubmitted},
	{"MF22959", "Richeval", 48.64005111, 6.909957222, PoleStateNotSubmitted},
	{"MF22960", "Richeval", 48.64028972, 6.909481667, PoleStateNotSubmitted},
	{"MF22961", "Richeval", 48.64042611, 6.908951944, PoleStateNotSubmitted},
	{"MF22962", "Richeval", 48.64065528, 6.908526944, PoleStateNotSubmitted},
	{"MF22963", "Richeval", 48.64093639, 6.908177778, PoleStateNotSubmitted},
	{"MF22964", "Richeval", 48.64118083, 6.907893056, PoleStateNotSubmitted},
	{"MF22965", "Richeval", 48.64107472, 6.907406944, PoleStateNotSubmitted},
	{"MF22966", "Ibigny", 48.64116639, 6.907057222, PoleStateNotSubmitted},
	{"MF22967", "Ibigny", 48.64101278, 6.90657, PoleStateNotSubmitted},
	{"MF22968", "Ibigny", 48.64085528, 6.906035833, PoleStateNotSubmitted},
	{"MF22969", "Ibigny", 48.64072833, 6.905538056, PoleStateNotSubmitted},
	{"MF22970", "Ibigny", 48.64081361, 6.905018333, PoleStateDone},
	{"MF22971", "Ibigny", 48.64092444, 6.9044975, PoleStateDone},
	{"MF22972", "Ibigny", 48.64109333, 6.904035833, PoleStateDone},
	{"MF22973", "Ibigny", 48.64141306, 6.903634722, PoleStateDone},
	{"MF22974", "Ibigny", 48.64173778, 6.903209722, PoleStateDone},
	{"MF22975", "Ibigny", 48.64210667, 6.902787222, PoleStateDone},
	{"MF22990", "Moussey", 48.68475333, 6.803333611, PoleStateDone},
	{"MF22991", "Foulcrey", 48.63615, 6.857541389, PoleStateToDo},
	{"MF22994", "Moussey", 48.68865667, 6.802375278, PoleStateDone},
	//{"MF22996", "Saint-Georges", 48.65733139, 6.925023889,PoleStateToDo},
	{"MF22997", "Foulcrey", 48.63688917, 6.855005556, PoleStateDone},
	//{"MF22998", "Rechicourt-le-Chateau", 48.68269694, 6.853323889,PoleStateToDo},
	{"MF22999", "Avricourt", 48.64946861, 6.814881389, PoleStateDone},
}
