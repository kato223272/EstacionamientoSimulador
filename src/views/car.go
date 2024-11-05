package views

import (
    "math/rand"
    "simulador/src/models/Observador"
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/canvas"
    "fyne.io/fyne/v2/storage"
)

type Carro struct {
    Image    *canvas.Image
    imagenes []string
}

func NuevoCarro() *Carro {
    carro := &Carro{
        imagenes: []string{
            "./assets/carro1.png",
            "./assets/taxi.png",
            "./assets/ambulancia.png",
            "./assets/morado.png",
            "./assets/vocho.png",
        },
    }
    return carro
}

func (c *Carro) AgregarCarro(container *fyne.Container) {
    indice := rand.Intn(len(c.imagenes)) 
    carImage := canvas.NewImageFromURI(storage.NewFileURI(c.imagenes[indice]))
    carImage.Resize(fyne.NewSize(50, 30))
    carImage.Move(fyne.NewPos(100, 100))  
    c.Image = carImage

    container.Add(carImage) 
    container.Refresh()    
}

func (c *Carro) Update(pos Observador.Pos) {
  
    if pos.X >= 0 && pos.X <= 800 && pos.Y >= 0 && pos.Y <= 500 { 
        c.Image.Move(fyne.NewPos(float32(pos.X), float32(pos.Y)))
    }
}
