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
        t.izq = &tnodo{valor: t.valor}
        t.valor = t.der.valor 
        t.der.valor = t.der.der.valor
        t.der.der = nil
        return
    }else{
    t.der = &tnodo{valor: t.valor}
    t.valor = t.izq.valor 
    t.izq.valor = t.izq.izq.valor
    t.izq.izq = nil
    return
    }
}

func (t *tnodo) rotarDoble(){
    fmt.Println("La funcion rotar doble nunca funciono, asi que la comente noma ya que no servia")
    fmt.Println("Pero la rotacion simple si funciona :3")
    }
    
/*func (t *tnodo) rotarDoble(){
    fmt.Println("Doblete")
    if t.izq.altura - t.der.altura == 2{//a la derecha
        R:= t.izq.der
        //Q := t.izq
        auxI := R.izq
        auxD := R.der
        P:= t.der 
        
        t.valor = R.valor
        
        //t.izq = Q 
        t.der = P
        t.der.izq = auxD
        t.izq.der = auxI
        return
    }else{ //izquierda
    R := t.der.izq 
    //Q := t.der
    auxD:= R.der 
    auxI := R.izq 
    P := t.izq
    
    t.valor = R.valor
    
    t.izq = P 
    //t.der = Q 
    t.izq.der = auxI
    t.der.izq = auxD
    return
    }
}*/

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
