package leafletmap

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/huckridgesw/hvue"
	"github.com/lpuig/ewin/doe/website/frontend/tools"
	"github.com/lpuig/ewin/mappoc/frontend/leaflet"
)

const template string = `
<div id="LeafLetMap" style="height: 100%"></div>
`

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Comp Registration

func RegisterComponent() hvue.ComponentOption {
	return hvue.Component("leaflet-map", ComponentOptions()...)
}

func ComponentOptions() []hvue.ComponentOption {
	return []hvue.ComponentOption{
		hvue.Template(template),
		hvue.Props("poles"),
		hvue.MethodsOf(&LeafletMap{}),
		//hvue.Computed("progressPct", func(vm *hvue.VM) interface{} {
		//	wspb := &WorksiteProgressBarModel{Object: vm.Object}
		//	return wspb.ProgressPct()
		//}),
		hvue.Mounted(func(vm *hvue.VM) {
			llm := newLeafletMap(vm)
			llm.Init()
		}),
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Comp Model

type LeafletMap struct {
	*js.Object

	Map *leaflet.Map `js:"Map"`

	VM *hvue.VM `js:"VM"`
}

func newLeafletMap(vm *hvue.VM) *LeafletMap {
	llm := &LeafletMap{Object: tools.O()}
	llm.VM = vm
	return llm
}

func (llm *LeafletMap) Init() {
	mapOption := leaflet.DefaultMapOptions()

	llm.Map = leaflet.NewMap("LeafLetMap", mapOption)
	osmlayer := leaflet.OSMTileLayer()
	//satlayer := leaflet.MapBoxTileLayer("mapbox.satellite")
	satlayer := leaflet.MapBoxTileLayer("mapbox.streets-satellite")

	baseMaps := js.M{
		"Plan":      osmlayer,
		"Satellite": satlayer,
	}

	osmlayer.AddTo(llm.Map)

	markerLayer := []*leaflet.Layer{}
	markerGroup := leaflet.NewLayerGroup(markerLayer)
	markerGroup.AddTo(llm.Map)

	overlayMaps := js.M{
		"Poteaux": markerGroup,
	}

	leaflet.NewControlLayers(baseMaps, overlayMaps).AddTo(llm.Map)
}

func (llm *LeafletMap) SetView(center *leaflet.LatLng, zoom int) {
	llm.Map.SetView(center, zoom)
}

func (llm *LeafletMap) FitBounds(bound1, bound2 *leaflet.LatLng) {
	llm.Map.FitBounds(bound1, bound2)
}
