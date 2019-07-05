package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/huckridgesw/hvue"
	"github.com/lpuig/ewin/doe/website/frontend/tools"
	"github.com/lpuig/ewin/mappoc/frontend/leaflet"
)

//go:generate bash ./makejs.sh

func main() {
	mpm := NewMainPageModel()

	hvue.NewVM(
		hvue.El("#mappoc_app"),
		hvue.DataS(mpm),
		hvue.MethodsOf(mpm),
		hvue.Mounted(func(vm *hvue.VM) {
			mpm := &MainPageModel{Object: vm.Object}
			mpm.InitMap()
			//js.Global.Call("confirm", "bonjour bonjour")
			js.Global.Get("window").Call("addEventListener", "beforeunload", mpm.Leave, false)
		}),
	)

	js.Global.Set("mpm", mpm)
}

type MainPageModel struct {
	*js.Object

	VM        *hvue.VM `js:"VM"`
	Longitude float64  `js:"Longitude"`
	Latitude  float64  `js:"Latitude"`

	Poles        []*Pole `js:"Poles"`
	ConfirmLeave bool    `js:"ConfirmLeave"`

	Map *leaflet.Map `js:"Map"`
}

func NewMainPageModel() *MainPageModel {
	mpm := &MainPageModel{Object: tools.O()}
	mpm.Longitude = 1
	mpm.Latitude = 1
	mpm.Poles = GenPoles(poles)
	mpm.Map = nil
	mpm.ConfirmLeave = false
	return mpm
}

func (mpm *MainPageModel) Leave(event *js.Object) {
	if !mpm.ConfirmLeave {
		return
	}
	event.Call("preventDefault")
	event.Set("returnValue", "")
	//js.Global.Call("confirm", "Sur ?")
}

func (mpm *MainPageModel) InitMap() {
	mapOption := leaflet.DefaultMapOptions()

	mpm.Map = leaflet.NewMap("mapEWIN", mapOption)
	osmlayer := leaflet.OSMTileLayer()
	//satlayer := leaflet.MapBoxTileLayer("mapbox.satellite")
	satlayer := leaflet.MapBoxTileLayer("mapbox.streets-satellite")

	baseMaps := js.M{
		"Plan":      osmlayer,
		"Satellite": satlayer,
	}

	osmlayer.AddTo(mpm.Map)

	polesLayer := []*leaflet.Layer{}

	for _, pole := range mpm.Poles {
		dio := leaflet.DefaultDivIconOptions()
		ico := leaflet.NewDivIcon(dio)
		mOption := leaflet.DefaultMarkerOption()
		mOption.Icon = &ico.Icon
		mOption.Opacity = 0.5
		mOption.Title = pole.Ref

		//marker := leaflet.NewMarker(pole.Lat, pole.Long, mOption)
		marker := NewPoleMarker(pole.Lat, pole.Long, mOption, pole)
		pole.PoleMarker = marker
		marker.BindPopup(pole.Ref)
		marker.UpdateFromState()
		marker.On("click", func(o *js.Object) {
			//print("event :", o)
			mpm.ConfirmLeave = true
			pole := &Pole{Object: o.Get("sourceTarget").Get("Pole")}
			pole.SwitchState()
			pole.PoleMarker.UpdateFromState()
			pole.PoleMarker.Refresh()
		})
		polesLayer = append(polesLayer, &marker.Layer)
	}

	polesGroup := leaflet.NewLayerGroup(polesLayer)
	polesGroup.AddTo(mpm.Map)

	overlayMaps := js.M{
		"Poteaux": polesGroup,
	}

	leaflet.NewControlLayers(baseMaps, overlayMaps).AddTo(mpm.Map)

	clat, clong, minlat, minlong, maxlat, maxlong := GetCenterAndBounds(mpm.Poles)
	mpm.Latitude, mpm.Longitude = clat, clong
	mpm.Map.SetView(leaflet.NewLatLng(mpm.Latitude, mpm.Longitude), 3)
	//mpm.Map.SetZoom(12)
	mpm.Map.FitBounds(leaflet.NewLatLng(minlat, minlong), leaflet.NewLatLng(maxlat, maxlong))

}
