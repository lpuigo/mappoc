package main

import (
	"time"

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
			//mpm.Poles = model.GenPoles(model.Poles)
			print("main mounted", mpm.Poles, mpm.Object)
			BeforeUnloadConfirmation(mpm.CanLeave)
			go mpm.LoadPole()
		}),
	)

	js.Global.Set("mpm", mpm)
}

type MainPageModel struct {
	*js.Object

	VM           *hvue.VM      `js:"VM"`
	Longitude    float64       `js:"Longitude"`
	Latitude     float64       `js:"Latitude"`
	Poles        []*model.Pole `js:"Poles"`
	ConfirmLeave bool          `js:"ConfirmLeave"`
}

func NewMainPageModel() *MainPageModel {
	mpm := &MainPageModel{Object: tools.O()}
	mpm.VM = nil
	mpm.Longitude = 1
	mpm.Latitude = 1
	mpm.Poles = []*model.Pole{}
	//mpm.Poles = model.GenPoles(model.Poles)
	mpm.ConfirmLeave = false
	return mpm
}

func (mpm *MainPageModel) LoadPole() {
	print("LoadPole started", mpm.Object)
	time.Sleep(3 * time.Second)
	mpm.Poles = model.GenPoles(model.Poles)
	mpm.UpdateMap()
	print("LoadPole done")
}

func (mpm *MainPageModel) CanLeave() bool {
	return !mpm.ConfirmLeave
}

func BeforeUnloadConfirmation(canLeave func() bool) {
	js.Global.Get("window").Call(
		"addEventListener",
		"beforeunload",
		func(event *js.Object) {
			if canLeave() {
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
