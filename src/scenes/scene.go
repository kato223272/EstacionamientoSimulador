package scenes

import (
	"image/color"
	"simulador/src/views"
	"simulador/src/models/Observador"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"time"
)

type Scene struct {
	scene           fyne.Window
	container       *fyne.Container
	backgroundColor color.Color
	startButton     *widget.Button
	resetButton     *widget.Button
	carros          []*views.Carro
}

func NewScene(scene fyne.Window, backgroundColor color.Color, theme fyne.Theme) *Scene {
	return &Scene{
		scene:           scene,
		container:       nil,
		backgroundColor: backgroundColor,
		startButton:     nil,
		resetButton:     nil,
		carros:          make([]*views.Carro, 0),
	}
}

func (s *Scene) Inicio() {
	rect := canvas.NewRectangle(s.backgroundColor)
	rect.Resize(fyne.NewSize(800, 500))
	rect.Move(fyne.NewPos(0, 0))

	s.container = container.NewWithoutLayout(rect)
	s.scene.SetContent(s.container)

	s.DibujarEstacionamiento()

	s.startButton = widget.NewButton("Iniciar", func() {
		s.Iniciar()
	})
	s.startButton.Move(fyne.NewPos(50, 450))
	s.container.Add(s.startButton)

	s.resetButton = widget.NewButton("Reiniciar", func() {
		s.Reiniciar()
	})
	s.resetButton.Move(fyne.NewPos(150, 450))
	s.container.Add(s.resetButton)
}

func (s *Scene) Iniciar() {
	s.startButton.Disable()

	for i := 0; i < 5; i++ {
	
		initialPosition := Observador.Pos{X: int32(60 + int(30*i)), Y: int32(200)} 
		carro.AgregarCarro(s.container)
		carro.Update(initialPosition) 
		s.carros = append(s.carros, carro)

		go s.MoverCarro(carro)
	}
}

func (s *Scene) MoverCarro(carro *views.Carro) {
	espacioAncho := float32(30) 

	for i := 0; i < 10; i++ { 
		carroPos := Observador.Pos{X: int32(90 + int(espacioAncho)*i), Y: 400}
		carro.Update(carroPos)
		time.Sleep(500 * time.Millisecond)
	}

	for i := 0; i < 10; i++ {
		carroPos := Observador.Pos{X: int32(740 - espacioAncho), Y: int32(200 + int(espacioAncho)*i)} 
		carro.Update(carroPos)
		time.Sleep(500 * time.Millisecond) 
	}
}

func (s *Scene) dibujarLineas(posX, posY, ancho, alto float32, lineColor color.Color) {
	
	linea := canvas.NewLine(lineColor)
	linea.StrokeWidth = 2
	linea.Position1 = fyne.NewPos(posX, posY)
	linea.Position2 = fyne.NewPos(posX+ancho, posY)
	s.container.Add(linea)
}


func (s *Scene) DibujarEstacionamiento() {
	spaceColor := color.RGBA{R: 230, G: 230, B: 230, A: 255}
	lineColor := color.RGBA{R: 100, G: 100, B: 100, A: 255}

	espacioAncho := float32(60)
	espacioAlto := float32(30)
	blockWidth := 5
	blockHeight := 2

	for i := 0; i < blockWidth; i++ {
		posX := 60 + float32(i)*(espacioAncho+10)
		posY := float32(400)

		rect := canvas.NewRectangle(spaceColor)
		rect.Resize(fyne.NewSize(espacioAncho, espacioAlto))
		rect.Move(fyne.NewPos(posX, posY))
		s.container.Add(rect)

		s.dibujarLineas(posX, posY, espacioAncho, espacioAlto, lineColor)
	}

	for i := 0; i < blockHeight; i++ {
		for j := 0; j < blockWidth; j++ {
			posX := float32(60)
			posY := 200 + float32(i)*(espacioAlto+10)

			rect := canvas.NewRectangle(spaceColor)
			rect.Resize(fyne.NewSize(espacioAncho, espacioAlto))
			rect.Move(fyne.NewPos(posX, posY))
			s.container.Add(rect)

			s.dibujarLineas(posX, posY, espacioAncho, espacioAlto, lineColor)
		}

		for j := 0; j < blockWidth; j++ {
			posX := 60 + float32(blockWidth)*(espacioAncho+10)
			posY := 200 + float32(i)*(espacioAlto+10)

			rect := canvas.NewRectangle(spaceColor)
			rect.Resize(fyne.NewSize(espacioAncho, espacioAlto))
			rect.Move(fyne.NewPos(posX, posY))
			s.container.Add(rect)

			s.dibujarLineas(posX, posY, espacioAncho, espacioAlto, lineColor)
		}
	}
}

func (s *Scene) Reiniciar() {
	s.container.RemoveAll()
	s.container.Refresh()

	s.carros = make([]*views.Carro, 0)

	s.container.Add(s.startButton)
	s.container.Add(s.resetButton)

	s.startButton.Enable()
}
