package leaflet

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/lpuig/ewin/doe/website/frontend/tools"
)

// Marker is a leaflet Marker: https://leafletjs.com/reference-1.5.0.html#marker.
type Marker struct {
	Layer
}

// NewMarker creates a new Marker
func NewMarker(lat, long float64, option *MarkerOptions) *Marker {
	return &Marker{
		Layer: Layer{
			Object: L.Call("marker", NewLatLng(lat, long), option),
		},
	}
}

func (m *Marker) BindPopup(content string) {
	m.Call("bindPopup", content)
}

type MarkerOptions struct {
	*js.Object
	//Icon Icon `js:"icon"`
	Keyboard            bool    `js:"keyboard"`
	Title               string  `js:"title"`
	Alt                 string  `js:"alt"`
	ZIndexOffset        float64 `js:"zIndexOffset"`
	Opacity             float64 `js:"opacity"`
	RiseOnHover         bool    `js:"riseOnHover"`
	RiseOffset          bool    `js:"riseOffset"`
	Pane                string  `js:"pane"`
	BubblingMouseEvents bool    `js:"bubblingMouseEvents"`
	Draggable           bool    `js:"draggable"`
	AutoPan             bool    `js:"autoPan"`
	//AutoPanPadding Point `js:"autoPanPadding"`
	AutoPanSpeed float64 `js:"autoPanSpeed"`
}

func DefaultMarkerOption() *MarkerOptions {
	return &MarkerOptions{Object: tools.O()}
}

/*

// TODO Implement this

type Icon struct {
	*js.Object
}

type IconOptions struct {
	*js.Object
}
*/
