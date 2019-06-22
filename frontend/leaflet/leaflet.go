// Package leaflet provides a (currently minimal) wrapper around leaflet.js
// for use with gopherjs. The bindings are currently for leaflet version 1.0.2.
package leaflet

import "github.com/gopherjs/gopherjs/js"

// L is the primary leaflet javascript object.
var L = js.Global.Get("L")

// Map is a leaflet map object: http://leafletjs.com/reference-1.0.2.html#map
type Map struct {
	*js.Object
}

// NewMap creates a new map in the div specified by divID.
func NewMap(divID string, options *MapOptions) *Map {
	return &Map{
		Object: L.Call("map", divID),
	}
}

// DefaultMapOptions returns the default Map options.
func DefaultMapOptions() *MapOptions {
	return &MapOptions{
		Object: js.Global.Get("Object").New(),
	}
}

// MapOptions specify the options for a map:
// http://leafletjs.com/reference-1.0.2.html#map.
// They need to be initialized with DefaultMapOptions.
type MapOptions struct {
	Object       *js.Object
	PreferCanvas bool `js:"preferCanvas"`
}

// SetView sets the center and zoom level of the map.
func (m *Map) SetView(center *LatLng, zoom int) {
	m.Object.Call("setView", center, zoom)
}

// CreatePane creates a new Pane with the given name:
// http://leafletjs.com/reference-1.0.2.html#map-createpane
func (m *Map) CreatePane(name string) *Pane {
	return &Pane{Object: m.Object.Call("createPane", name)}
}

// Pane is a leaflet pane.
type Pane struct {
	*js.Object
}

// SetZIndex sets the Z index of the pane.
func (p *Pane) SetZIndex(index int) {
	p.Object.Get("style").Set("zIndex", index)
}

// LatLng specifies a point in latitude and longitude
type LatLng struct {
	*js.Object
}

// NewLatLng returns a new LatLng object.
func NewLatLng(lat, lng float64) *LatLng {
	return &LatLng{
		Object: L.Call("latLng", lat, lng),
	}
}

// TileLayer is a leaflet TileLayer object: http://leafletjs.com/reference-1.0.2.html#tilelayer
type TileLayer struct {
	Layer
}

// NewTileLayer creates a new TileLayer with the specified URL template and
// options.
func NewTileLayer(urlTemplate string, options *TileLayerOptions) *TileLayer {
	return &TileLayer{
		Layer: Layer{
			Object: L.Call("tileLayer", urlTemplate, options),
		},
	}
}

// TileLayerOptions specifies the options for the TileLayer: http://leafletjs.com/reference-1.0.2.html#tilelayer-option
// They need to be initialized with DefaultTileLayerOptions.
type TileLayerOptions struct {
	*js.Object
	MinZoom       int      `js:"minZoom"`
	MaxZoom       int      `js:"maxZoom"`
	MinNativeZoom int      `js:"minNativeZoom"`
	MaxNativeZoom int      `js:"maxNativeZoom"`
	Subdomains    []string `js:"subdomains"`
	ErrorTileURL  string   `js:"errorTileUrl"`
	ZoomOffset    int      `js:"zoomOffset"`
	TMS           bool     `js:"tms"`
	ZoomReverse   bool     `js:"zoomReverse"`
	DetectRetina  bool     `js:"detectRetina"`
	CrossOrigin   bool     `js:"crossOrigin"`
	Id            string   `js:"id"`
	AccesToken    string   `js:"accessToken"`

	Pane        string `js:"pane"`
	Attribution string `js:"attribution"`
}

// DefaultTileLayerOptions returns the default TileLayer options.
func DefaultTileLayerOptions() *TileLayerOptions {
	return &TileLayerOptions{
		Object: js.Global.Get("Object").New(),
	}
}

// Layer is a leaflet layer object: http://leafletjs.com/reference-1.0.2.html#layer.
type Layer struct {
	*js.Object
}

// AddTo add the receiver to the specified Map.
func (l *Layer) AddTo(m *Map) {
	l.Object.Call("addTo", m)
}

// Path is a leaflet path object: http://leafletjs.com/reference-1.0.2.html#path.
type Path struct {
	Layer
}

// SetStyle sets the style of the receiver:
// http://leafletjs.com/reference-1.0.2.html#path-setstyle.
func (p *Path) SetStyle(style *PathOptions) {
	p.Object.Call("setStyle", style)
}

// PathOptions specify the options for a path:
// http://leafletjs.com/reference-1.0.2.html#path-option.
// They need to be initialized with DefaultPathOptions.
type PathOptions struct {
	Object      *js.Object
	Stroke      bool    `js:"stroke"`
	Color       string  `js:"color"`
	Weight      int     `js:"weight"`
	Opacity     float64 `js:"opacity"`
	LineCap     string  `js:"lineCap"`
	LineJoin    string  `js:"lineJoin"`
	DashArray   string  `js:"dashArray"`
	DashOffset  string  `js:"dashOffset"`
	Fill        bool    `js:"fill"`
	FillColor   string  `js:"fillColor"`
	FillOpacity float64 `js:"fillOpacity"`
	FillRule    string  `js:"fillRule"`
}

// DefaultPathOptions returns the default TileLayer options.
func DefaultPathOptions() *PathOptions {
	return &PathOptions{
		Object: js.Global.Get("Object").New(),
	}
}

// Polyline is a leaflet polyline object: http://leafletjs.com/reference-1.0.2.html#polyline.
type Polyline struct {
	Path
}

// Polygon is a leaflet polygon object: http://leafletjs.com/reference-1.0.2.html#polygon.
type Polygon struct {
	Polyline
}

// NewPolygon creates a new polygon.
func NewPolygon(latlngs []*LatLng) *Polygon {
	return &Polygon{
		Polyline: Polyline{
			Path: Path{
				Layer: Layer{
					Object: L.Call("polygon", latlngs),
				},
			},
		},
	}
}

// GridLayer is a leaflet GridLayer: http://leafletjs.com/reference-1.0.2.html#gridlayer.
type GridLayer struct {
	Layer
}

// NewGridLayer creates a new GridLayer.
func NewGridLayer() *GridLayer {
	return &GridLayer{
		Layer: Layer{
			Object: L.Call("gridLayer"),
		},
	}
}
