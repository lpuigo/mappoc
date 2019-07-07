package polemap

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/huckridgesw/hvue"
	"github.com/lpuig/ewin/mappoc/frontend/leaflet"
	"github.com/lpuig/ewin/mappoc/frontend/mappoc/comp/leafletmap"
	"github.com/lpuig/ewin/mappoc/frontend/mappoc/model"
)

const template string = `
<div id="LeafLetMap" style="height: 100%"></div>
`

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Comp Registration

func RegisterComponent() hvue.ComponentOption {
	return hvue.Component("pole-map", componentOptions()...)
}

func componentOptions() []hvue.ComponentOption {
	return []hvue.ComponentOption{
		hvue.Template(template),
		hvue.Props("poles"),
		hvue.MethodsOf(&PoleMap{}),
		//hvue.Computed("progressPct", func(vm *hvue.VM) interface{} {
		//	wspb := &WorksiteProgressBarModel{Object: vm.Object}
		//	return wspb.ProgressPct()
		//}),
		hvue.Mounted(func(vm *hvue.VM) {
			pm := PoleMapFromJS(vm.Object)
			pm.Init()
			pm.AddPoles(pm.Poles, "Poteaux")
			print("polemap mounted", pm.Poles)
		}),
		hvue.BeforeUpdate(func(vm *hvue.VM) {
			pm := PoleMapFromJS(vm.Object)
			print("polemap beforeUpdate", pm.Poles)
			pm.AddPoles(pm.Poles, "attrib Poteaux")
		}),
		hvue.Updated(func(vm *hvue.VM) {
			pm := PoleMapFromJS(vm.Object)
			print("polemap Updated", pm.Poles)
		}),
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Comp Model

type PoleMap struct {
	leafletmap.LeafletMap

	Poles []*model.Pole `js:"poles"`

	VM *hvue.VM `js:"VM"`
}

func PoleMapFromJS(obj *js.Object) *PoleMap {
	return &PoleMap{LeafletMap: leafletmap.LeafletMap{Object: obj}}
}

func newPoleMap(vm *hvue.VM) *PoleMap {
	pm := &PoleMap{LeafletMap: *leafletmap.NewLeafletMap(vm)}
	pm.Poles = nil
	return pm
}

func (pm *PoleMap) Init() {
	pm.LeafletMap.Init()
	pm.LeafletMap.SetView(leaflet.NewLatLng(48, 6), 5)
}

func (pm *PoleMap) AddPoles(poles []*model.Pole, name string) {
	polesLayer := []*leaflet.Layer{}

	for _, pole := range poles {
		dio := leaflet.DefaultDivIconOptions()
		ico := leaflet.NewDivIcon(dio)
		mOption := leaflet.DefaultMarkerOption()
		mOption.Icon = &ico.Icon
		mOption.Opacity = 0.5
		mOption.Title = pole.Ref

		//poleMarker := leaflet.NewMarker(pole.Lat, pole.Long, mOption)
		poleMarker := NewPoleMarker(pole.Lat, pole.Long, mOption, pole)
		poleMarker.BindPopup(pole.Ref)
		poleMarker.UpdateFromState()
		poleMarker.On("click", func(o *js.Object) {
			print("event :", o)
			//mpm.ConfirmLeave = true // todo dependecy injection for dirty control
			//pole := &model.Pole{Object: o.Get("sourceTarget").Get("Pole")}
			//pole.SwitchState()
			//pole.PoleMarker.UpdateFromState()
			//pole.PoleMarker.Refresh()
		})
		polesLayer = append(polesLayer, &poleMarker.Layer)
	}

	polesGroup := leaflet.NewLayerGroup(polesLayer)
	polesGroup.AddTo(pm.LeafletMap.Map)
	pm.LeafletMap.ControlLayers.AddOverlay(&polesGroup.Layer, name)

	clat, clong, minlat, minlong, maxlat, maxlong := model.GetCenterAndBounds(poles)
	pm.LeafletMap.Map.SetView(leaflet.NewLatLng(clat, clong), 12)
	pm.LeafletMap.Map.FitBounds(leaflet.NewLatLng(minlat, minlong), leaflet.NewLatLng(maxlat, maxlong))

}
