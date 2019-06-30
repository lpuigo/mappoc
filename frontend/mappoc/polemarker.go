package main

import "github.com/lpuig/ewin/mappoc/frontend/leaflet"

type PoleMarker struct {
	leaflet.Marker
	Pole *Pole `js:"Pole"`
}

func NewPoleMarker(lat, long float64, option *leaflet.MarkerOptions, pole *Pole) *PoleMarker {
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
	case PoleStateNotSubmitted:
		html = pmHtmlPin
		class = ""
	case PoleStateToDo:
		html = pmHtmlPlain
		class = "blue"
	case PoleStateHoleDone:
		html = pmHtmlHole
		class = "orange"
	case PoleStateIncident:
		html = pmHtmlHole
		class = "red"
	case PoleStateDone:
		html = pmHtmlPlain
		class = "green"
	case PoleStateCancelled:
		html = pmHtmlPlain
		class = ""
	}

	pm.UpdateDivIconClassname(class)
	pm.UpdateDivIconHtml(html)
}
