package main
import (
    "fmt"
    "math"
    )

type tnodo struct{
    valor int
    altura float64
    izq, der *tnodo
    }
    
type avl struct{
    head *tnodo
    }

func nuevo() *avl {
	return new(avl)
}

func (t *avl) buscar(x int) bool{
    if t == nil{
        return false
        }
    aux := t.head
    for aux!= nil{
        if aux.valor == x{
            return true
            }
        if x < aux.valor{
            aux = aux.izq
            }else{
                aux = aux.der
                }
        }
    return false
    }

func (t *tnodo) rotarSimple() {
    if t.izq == nil{
        aux := t
        t = t.der 
        aux.der = nil 
        t.izq = aux
        return
        }else{
    aux := t 
    t = t.izq
    aux.izq = nil
    t.der = aux
    return
        }
    }

func (t *tnodo) rotarDoble(){
    if t.izq == nil{
        aux := t 
        aux2 := t.izq
        t = aux2
        t.der = aux
        return
        }else{
    aux := t 
    aux2 := t.der
    t = aux2
    t.izq = aux 
    return
        }
    }

func (t *tnodo) actualizarAltura(){
    if t == nil{return}
    t.izq.actualizarAltura()
    t.der.actualizarAltura()
    if t.der == nil && t.izq == nil{
        t.altura = 1
        return
    }
    if t.der == nil{
        t.altura = t.izq.altura +1
        return
        }
    if t.izq == nil{
        t.altura = t.der.altura +1
        return
        }
    t.altura = math.Max(t.izq.altura, t.der.altura) + 1
    }

func (t *tnodo) revisarAlturas(){
    if t == nil {return}
    t.izq.revisarAlturas()
    t.der.revisarAlturas()
    if t.der == nil && t.der == nil {return}
    if t.der == nil{
        if t.izq.altura >= 2 || t.izq.altura <= -2{
            if t.izq.der == nil{
                t.rotarSimple()
                //t.actualizarAltura()
                return
                }else{
                    t.rotarDoble()
                    //t.actualizarAltura()
                    return
                    }
            }
        }
    if t.izq == nil{
        if t.der.altura >=2 || t.der.altura <=-2{
            if t.der.izq == nil{
                t.rotarSimple()
                //t.actualizarAltura()
                return
                }else{
                    t.rotarDoble()
                    //t.actualizarAltura()
                    return
                    }
            }
        }
    if t.der != nil && t.izq != nil{
    if t.izq.altura - t.der.altura <=-2 || t.izq.altura - t.der.altura >=2{
        if t.der.izq == nil{
                t.rotarSimple()
                //t.actualizarAltura()
                return
                }else{
                    t.rotarDoble()
                    //t.actualizarAltura()
                    return
                    }
        if t.izq.altura >= 2 || t.izq.altura <= -2{
            if t.izq.der == nil{
                t.rotarSimple()
                //t.actualizarAltura()
                return
                }else{
                    t.rotarDoble()
                    //t.actualizarAltura()
                    return
                    }
        }
    }
}
}

func (t *avl) insertar(x int){
    //aux := t.head
    if t.head == nil{
        t.head = &tnodo{valor: x}
        t.head.actualizarAltura()
        t.head.revisarAlturas()
        return
        }
    aux := t.head
    for aux != nil{
        if aux.valor < x{
            if aux.der == nil{
                aux.der = &tnodo{valor: x}
                t.head.actualizarAltura()
                t.head.revisarAlturas()
                return
                }
            aux = aux.der
            }else{
                if aux.izq == nil{
                    aux.izq = &tnodo{valor: x}
                    t.head.actualizarAltura()
                    t.head.revisarAlturas()
                    return
                    }
                aux = aux.izq
                }
        }
    return
    }

func (t *tnodo) imprimir(){
    if t == nil{return}
    t.izq.imprimir()
    fmt.Print("Valor: ")
    fmt.Print(t.valor)
    fmt.Print(" Altura: ")
    fmt.Println(t.altura)
    t.der.imprimir()
    }
    
func main(){
    //fmt.Println("Todo esta bien =D")
    arbol := nuevo()
    arbol.insertar(1)
    arbol.insertar(2)
    arbol.insertar(3)
    //arbol.insertar(4)
    //arbol.insertar(5)
    //arbol.insertar(6)
    //arbol.insertar(7)
    //arbol.insertar(8)
    //arbol.insertar(9)
    //arbol.insertar(10)
    arbol.head.imprimir()
    }
