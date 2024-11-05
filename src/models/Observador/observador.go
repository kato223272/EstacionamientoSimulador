package Observador

type Pos struct {
    X int32
    Y int32
}

type Observador interface {
    Update(pos Pos)
}

type Sujetos interface {
    RegistrarObservador(observer Observador) 
    EliminarObservador(observer Observador)  
    NotificarPosicion()
}
