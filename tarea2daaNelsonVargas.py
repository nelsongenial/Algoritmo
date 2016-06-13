from operator import itemgetter, attrgetter, methodcaller
from time import time
import random

#Recibo una tupla
def  ambicioso(A, K, W):
    #A es una tupla
    A = sorted(A, key=itemgetter(2), reverse=True)
    # Eliminar A.pop(0)
    Final = []
    evaluar(A,K,W, Final)
    return Final

def evaluar(A,K,W, arreglo):
    if len(A) ==0:
        return 
    cont = 1
    if len(A) > 1:
        while A[0][2] == A[cont][2]:
            cont = cont +1 
    aux1 = A[cont::]
    aux = sorted(A[0:cont],key=itemgetter(1), reverse=True)
    A = aux1
    K,W = poi(aux,K,W,arreglo)
    if len(A) == 0:
        return 
    evaluar(A,K,W,arreglo)

def poi(A,K,W,arreglo):
    if K <=0:
        return K,W
    for i in range(len(A)):
        if W > A[i][1] and K > 0:
            arreglo.append(A[i])
            W = W-A[i][1]
            K = K-1
    return K,W

A = [
        ('Tazas', 1, 4),
        ('Sillon', 4, 5),
        ('Televisor', 3, 1),
        ('Radio', 11, 4),
        ('Mesa Comedor', 3, 4),
]
#random.randrange(10)
#def llenar(A):
#    for i in range(100):
#        a = ('algo', 1,10)
#        A.append(a)
#    return A

inicio = time()
Final = ambicioso(A,5,10)
final = time() - inicio
print "Tiempo", final
print "Colocar en el container los siguientes objetos:"
for i in range(len(Final)):
    print "-", Final[i][0]
