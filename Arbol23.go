package main
import (
    "fmt"
    )

type tnodo struct{
    vaizq, vader int
    izq, centro, der, arriba *tnodo
    }

type dostres struct{
    head *tnodo
    }

func nuevo() *dostres {
	return new(dostres)
}

func (t *dostres) buscar(x int)bool{
    aux := t.head
    if aux == nil{return false}
    for aux != nil{
        if aux.vaizq == x || aux.vader == x{return true}
        if x < aux.vaizq{
            aux = aux.izq
            }else{
            if aux.vaizq < x && x < aux.vader {
                aux = aux.centro
                }else{
                    aux = aux.der
                    }
            }
        }
    return false
    }

func (t *dostres) arreglar(){
    aux := t.head
    for aux != nil{
        if aux.arriba != nil{
            t.head = aux.arriba
            }
        aux = aux.arriba
        }
    }

func (t *tnodo) comprobar(){
    if t == nil{return}
    fmt.Println("Entro")
    //if t.vaizq != 0 && t.vader != 0{
        if t.izq != nil && t.der == nil{
            if t.izq.vaizq != 0 && t.izq.vader !=0{
                    t.der = &tnodo{vaizq: t.izq.vaizq, arriba : t}
                    t.izq.vader = 0
            }
    }
        if t.der != nil && t.izq == nil{
            if t.der.vaizq != 0 && t.der.vader !=0{
                t.izq = &tnodo{vaizq: t.der.vaizq, arriba: t}
                t.der.vaizq = t.der.vader
                t.der.vader = 0
            }
        }
    //}
    t.izq.comprobar()
    t.der.comprobar()
    }

func (t *tnodo) insertarArriba(x int){
    if t.arriba == nil{
        t.arriba = &tnodo{vaizq: x, izq: t}
        aux := t.vader
        t.vader = 0
        t.arriba.der = &tnodo{vaizq: aux, arriba: t.arriba}
        t.arriba.der.der = t.der
        t.der = t.centro
        t.centro = nil
        return
        }
    switch{
        case t.arriba.vader == 0:
            t.arriba.vader = x 
            t.arriba.centro = &tnodo{vaizq:t.vaizq, arriba: t.arriba}
            t.vaizq = t.vader
            t.vader = 0
            return
        case x < t.arriba.vaizq && t.arriba.vader != 0:
            aux := t.arriba.vaizq
            t.arriba.vaizq = x
            t.arriba.insertarArriba(aux)
            return
        case x> t.arriba.vaizq && x< t.arriba.vader:
            t.arriba.insertarArriba(x)
            return
        default:
            aux := t.arriba.vader
            t.arriba.vader = x
            t.arriba.insertarArriba(aux)
            return
    }
}



func (t *dostres) insertar(x int) {
    if t.head == nil{
        t.head = &tnodo{vaizq: x}
        return
        }
    aux := t.head
    for aux != nil{
        switch{
            case x< aux.vaizq:
                if aux.izq == nil{
                    if aux.vader == 0{
                        aux2 := aux.vaizq
                        aux.vaizq = x
                        aux.vader = aux2
                        t.arreglar()
                        t.head.comprobar()
                        return
                    }
                    aux2:= aux.vaizq
                    aux.vaizq = x
                    aux.insertarArriba(aux2)
                    t.arreglar()
                    t.head.comprobar()
                    return
                aux = aux.izq
                }
            case x > aux.vaizq && aux.vader == 0:
                if aux.der == nil{
                    aux.vader = x
                    t.arreglar()
                    t.head.comprobar()
                    return
                }
                aux = aux.der
            case x > aux.vaizq && x < aux.vader:
                if aux.centro == nil{
                    aux.insertarArriba(x)
                    t.arreglar()
                    t.head.comprobar()
                    return
                    }
                aux = aux.centro
            default:
                if aux.der == nil{
                    aux2:= aux.vader
                    aux.vader = x 
                    aux.insertarArriba(aux2)
                    t.arreglar()
                    t.head.comprobar()
                    return
                    }
                aux = aux.der
        }
    }
    t.arreglar()
    t.head.comprobar()
    return
}


func (t *tnodo) imprimir(){
    if t == nil{return}
    t.izq.imprimir()
    fmt.Print(t.vaizq)
    fmt.Print(" ")
    fmt.Println(t.vader)
    t.centro.imprimir()
    t.der.imprimir()
    }


func main(){
    fmt.Println("Todo esta bien =D")
    arbol := nuevo()
    arbol.insertar(1)
    arbol.insertar(2)
    arbol.insertar(3)
    arbol.insertar(4)
    arbol.insertar(5)
    arbol.insertar(6)
    arbol.insertar(7)
    arbol.insertar(8)
    arbol.insertar(9)
    arbol.insertar(10)
    arbol.head.imprimir()
    }
