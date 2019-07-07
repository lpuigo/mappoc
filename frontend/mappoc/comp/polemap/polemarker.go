package polemap

import (
	"github.com/lpuig/ewin/mappoc/frontend/leaflet"
	"github.com/lpuig/ewin/mappoc/frontend/mappoc/model"
)

type PoleMarker struct {
	leaflet.Marker
	Pole *model.Pole `js:"Pole"`
}

func NewPoleMarker(lat, long float64, option *leaflet.MarkerOptions, pole *model.Pole) *PoleMarker {
	np := &PoleMarker{Marker: *leaflet.NewMarker(lat, long, option)}
	np.Pole = pole
	return np
}

const (
	pmHtmlPin   string = `<i class="fas fa-map-pin fa-3x"></i>`
	pmHtmlPlain string = `<i class="fas fa-map-marker fa-3x"></i>`
	pmHtmlHole  string = `<i class="fas fa-map-marker-alt fa-3x"></i>`
)

func (pm *PoleMarker) UpdateFromState() {
	var html, class string

	switch pm.Pole.State {
	case model.PoleStateNotSubmitted:
		html = pmHtmlPin
		class = ""
	case model.PoleStateToDo:
		html = pmHtmlPlain
		class = "blue"
	case model.PoleStateHoleDone:
		html = pmHtmlHole
		class = "orange"
	case model.PoleStateIncident:
		html = pmHtmlHole
		class = "red"
	case model.PoleStateDone:
		html = pmHtmlPlain
		class = "green"
	case model.PoleStateCancelled:
		html = pmHtmlPlain
		class = ""
	}

	pm.UpdateDivIconClassname(class)
	pm.UpdateDivIconHtml(html)
}
