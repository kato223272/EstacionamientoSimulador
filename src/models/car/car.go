
package car

import (
    "time"
    "simulador/src/models/Observador"
)

type Car struct {
    posX, posY   int32
    status       bool
    Observadores []Observador.Observador
}

func NuevoCar() *Car {
    return &Car{posX: 0, posY: 0, status: true}
}

func (c *Car) Movimiento(incX, incY int32) {
    c.status = true
    for c.status {
    
        if c.posX < 30 || c.posX > 770 {
            incX *= -1
        }
        if c.posY < 30 || c.posY > 470 {
            incY *= -1
        }

        c.posX += incX
        c.posY += incY
        c.NotificarPosicion()
        time.Sleep(300 * time.Millisecond)
    }
}



func (c *Car) EstablecerEstacionamiento() {
    c.status = false
    time.Sleep(5 * time.Second)
    c.NotificarPosicion()
}

func (c *Car) EnviarStatus(status bool) {
    c.status = status
}

func (c *Car) RegistrarObservador(observer Observador.Observador) {
    c.Observadores = append(c.Observadores, observer)
}

func (c *Car) EliminarObservador(observer Observador.Observador) {
    for i, o := range c.Observadores {
        if o == observer {
            c.Observadores = append(c.Observadores[:i], c.Observadores[i+1:]...)
            break
        }
    }
}

func (c *Car) NotificarPosicion() {
    for _, observer := range c.Observadores {
        observer.Update(Observador.Pos{X: c.posX, Y: c.posY})
    }
}
