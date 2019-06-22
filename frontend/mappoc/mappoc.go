package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/huckridgesw/hvue"
	"github.com/lpuig/ewin/doe/website/frontend/tools"
	"github.com/lpuig/ewin/mappoc/frontend/leaflet"
)

//go:generate bash ./makejs.sh

const MapboxToken string = "pk.eyJ1IjoibGF1cmVudC1wdWlnIiwiYSI6ImNqeDgxazRqYzBmOGEzbnA3Z2lld3Rja2cifQ.Oq6cQfmK3uKYyVQffiIn_Q"

func main() {
	mpm := NewMainPageModel()

	hvue.NewVM(
		hvue.El("#mappoc_app"),
		hvue.DataS(mpm),
		hvue.MethodsOf(mpm),
		hvue.Mounted(func(vm *hvue.VM) {
			mpm := &MainPageModel{Object: vm.Object}
			mpm.InitMap()
		}),
	)

	js.Global.Set("mpm", mpm)
}

type MainPageModel struct {
	*js.Object

	VM        *hvue.VM `js:"VM"`
	Longitude float64  `js:"Longitude"`
	Latitude  float64  `js:"Latitude"`

	Map *leaflet.Map `js:"Map"`
}

func NewMainPageModel() *MainPageModel {
	mpm := &MainPageModel{Object: tools.O()}
	mpm.Longitude = 1
	mpm.Latitude = 1
	mpm.Map = nil
	return mpm
}

func (mpm *MainPageModel) InitMap() {
	mpm.Map = leaflet.NewMap("mapISS", leaflet.DefaultMapOptions())
	mpm.Map.SetView(leaflet.NewLatLng(mpm.Latitude, mpm.Longitude), 1)

	tileOption := leaflet.DefaultTileLayerOptions()

	//tileOption.Attribution = `&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors`
	//tileOption.Id = "mapbox.streets"
	//tileOption.AccesToken = MapboxToken
	//url := "https://api.tiles.mapbox.com/v4/{id}/{z}/{x}/{y}.png?access_token={accessToken}"

	tileOption.Attribution = `&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors`
	//tileOption.Id = "mapbox.streets"
	//tileOption.AccesToken = MapboxToken
	url := "https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
	tile := leaflet.NewTileLayer(url, tileOption)
	tile.AddTo(mpm.Map)

}
