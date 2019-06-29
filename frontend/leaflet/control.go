package leaflet

import "github.com/gopherjs/gopherjs/js"

// Control is a leaflet Control object: https://leafletjs.com/reference-1.5.0.html#control.
type Control struct {
	*js.Object
}

// AddTo add the receiver to the specified Map.
func (c *Control) AddTo(m *Map) {
	c.Object.Call("addTo", m)
}

// ControlLayers is a leaflet Control.Layers object: https://leafletjs.com/reference-1.5.0.html#control-layers.
type ControlLayers struct {
	Control
}

func NewControlLayers(baseLayer js.M, overlays js.M) *ControlLayers {
	return &ControlLayers{
		Control{L.Get("control").Call("layers", baseLayer, overlays)},
	}
}
