package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/huckridgesw/hvue"
	"github.com/lpuig/ewin/doe/website/frontend/tools"
	"github.com/lpuig/ewin/mappoc/frontend/mappoc/comp/polemap"
	"github.com/lpuig/ewin/mappoc/frontend/mappoc/model"
)

//go:generate bash ./makejs.sh

func main() {
	mpm := NewMainPageModel()

	hvue.NewVM(
		polemap.RegisterComponent(),
		hvue.El("#mappoc_app"),
		hvue.DataS(mpm),
		hvue.MethodsOf(mpm),
		hvue.Mounted(func(vm *hvue.VM) {
			mpm := &MainPageModel{Object: vm.Object}
			mpm.Poles = []*model.Pole{}
			BeforeUnloadConfirmation(mpm.PreventLeave)
			go mpm.LoadPole()
		}),
	)

	js.Global.Set("mpm", mpm)
}

type MainPageModel struct {
	*js.Object

	VM        *hvue.VM      `js:"VM"`
	Longitude float64       `js:"Longitude"`
	Latitude  float64       `js:"Latitude"`
	Poles     []*model.Pole `js:"Poles"`
	Dirty     bool          `js:"Dirty"`
}

func NewMainPageModel() *MainPageModel {
	mpm := &MainPageModel{Object: tools.O()}
	mpm.VM = nil
	mpm.Longitude = 1
	mpm.Latitude = 1
	mpm.Poles = []*model.Pole{}
	mpm.Dirty = false
	return mpm
}

func (mpm *MainPageModel) LoadPole() {
	mpm.Poles = model.GenPoles(model.Poles)
	mpm.UpdateMap()
}

func (mpm *MainPageModel) PreventLeave() bool {
	return mpm.Dirty
}

// BeforeUnloadConfirmation activate confirm leave alert if askBeforeLeave func return true
func BeforeUnloadConfirmation(askBeforeLeave func() bool) {
	js.Global.Get("window").Call(
		"addEventListener",
		"beforeunload",
		func(event *js.Object) {
			if !askBeforeLeave() {
				return
			}
			event.Call("preventDefault")
			event.Set("returnValue", "")
			//js.Global.Call("confirm", "Sur ?")

		},
		false)
}

func (mpm *MainPageModel) UpdateMap() {
	pm := polemap.PoleMapFromJS(mpm.VM.Refs("MapEwin"))
	pm.AddPoles(mpm.Poles, "Poteaux")
}

func (mpm *MainPageModel) MarkerClick(poleMarkerObj, event *js.Object) {
	pm := polemap.PoleMarkerFromJS(poleMarkerObj)
	pm.Pole.SwitchState()
	pm.UpdateFromState()
	pm.Refresh()
	//print("MarkerClick", event)
}
