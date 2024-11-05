package main

import (
	"simulador/src/scenes"
	"image/color"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2"
)

func main() {
	myApp := app.New()
	stage := myApp.NewWindow("Simulador de Carro")
	stage.CenterOnScreen()
	stage.Resize(fyne.NewSize(800, 500))
	stage.SetFixedSize(true)

	
	scene := scenes.NewScene(
		stage,
		color.RGBA{R: 150, G: 150, B: 150, A: 200},
		nil, 
	)

	scene.Inicio()

	stage.ShowAndRun()
}
