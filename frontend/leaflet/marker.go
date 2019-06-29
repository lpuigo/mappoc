package leaflet

// Marker is a leaflet Marker: https://leafletjs.com/reference-1.5.0.html#marker.
type Marker struct {
	Layer
}

// NewMarker creates a new Marker
func NewMarker(lat, long float64) *Marker {
	return &Marker{
		Layer: Layer{
			Object: L.Call("marker", NewLatLng(lat, long)),
		},
	}
}
