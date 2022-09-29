package lista

type nodoLista[T any] struct {
	dato      T
	siguiente *nodoLista[T]
}

func crearNodo[T any]() *nodoLista[T] {
	nodo := new(nodoLista[T])
	return nodo
}

type listaEnlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo   int
}

// FUNCIÓN PARA CREAR LA LISTA ENLAZADA

func CrearListaEnlazada[T any]() Lista[T] {
	//...
	lista := new(listaEnlazada[T])

	return lista
}

func (lista *listaEnlazada[T]) EstaVacia() bool {
	return lista.primero == nil
}

func (lista *listaEnlazada[T]) InsertarPrimero(elem T) {

}

func (lista *listaEnlazada[T]) InsertarUltimo(elem T) {

}

func (lista *listaEnlazada[T]) BorrarPrimero() T {
	if lista.EstaVacia() {
		panic("lista vacia")
	}
	siguiente := lista.primero.siguiente
	lista.primero.siguiente = nil
	lista.primero = siguiente

	if lista.primero == nil {
		lista.ultimo = nil
	}
}

func (lista *listaEnlazada[T]) VerPrimero() T {
	if lista.primero != nil {
		return lista.primero.dato
	}
	panic("lista vacia")
}

func (lista *listaEnlazada[T]) VerUltimo() T {
	if lista.ultimo != nil {
		return lista.ultimo.dato
	}
	panic("lista vacia")
}

func (lista *listaEnlazada[T]) Largo() int {
	return lista.largo
}

/*
func (lista *listaEnlazada[T]) Iterar(visitar func(T) bool){

}
*/
