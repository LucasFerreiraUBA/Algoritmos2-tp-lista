package lista

type nodoLista[T any] struct {
	dato      T
	siguiente *nodoLista[T]
}

func crearNodo[T any](dato T) *nodoLista[T] {
	nodo := new(nodoLista[T])
	nodo.dato = dato
	return nodo
}

type listaEnlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo   int
}

type iteradorListaEnlazada[T any] struct {
	lista    *listaEnlazada[T]
	actual   *nodoLista[T]
	anterior *nodoLista[T]
}

// FUNCIÓN PARA CREAR LA LISTA ENLAZADA

func CrearListaEnlazada[T any]() Lista[T] {
	lista := new(listaEnlazada[T])
	lista.largo = 0
	return lista
}

func (lista listaEnlazada[T]) EstaVacia() bool {
	return lista.primero == nil
}

func (lista *listaEnlazada[T]) InsertarPrimero(elem T) {
	nuevo_nodo := crearNodo(elem)
	lista.largo++
	nuevo_nodo.siguiente = lista.primero
	lista.primero = nuevo_nodo
	if lista.largo == 1 {
		lista.ultimo = lista.primero
	}
}

func (lista *listaEnlazada[T]) InsertarUltimo(elem T) {
	if lista.EstaVacia() {
		lista.InsertarPrimero(elem)
		return
	}
	nuevo_nodo := crearNodo(elem)
	lista.ultimo.siguiente = nuevo_nodo
	lista.ultimo = nuevo_nodo
	lista.largo++
}

func (lista *listaEnlazada[T]) BorrarPrimero() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	dato := lista.primero.dato
	siguiente := lista.primero.siguiente
	lista.primero = siguiente
	if lista.primero == nil {
		lista.ultimo = nil
	}
	lista.largo--
	return dato
}

func (lista *listaEnlazada[T]) VerPrimero() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	return lista.primero.dato
}

func (lista *listaEnlazada[T]) VerUltimo() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	return lista.ultimo.dato
}

func (lista listaEnlazada[T]) Largo() int {
	return lista.largo
}

func (lista listaEnlazada[T]) Iterar(visitar func(T) bool) {
	actual := lista.primero
	for actual != nil {
		if !visitar(actual.dato) {
			return
		}
		actual = actual.siguiente
	}
}

func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {
	iterador := new(iteradorListaEnlazada[T])
	iterador.lista = lista
	iterador.actual = iterador.lista.primero
	return iterador
}

func (iterador *iteradorListaEnlazada[T]) Insertar(dato T) {
	if iterador.anterior == nil {
		iterador.lista.InsertarPrimero(dato)
		iterador.actual = iterador.lista.primero
	} else if iterador.actual == nil {
		iterador.lista.InsertarUltimo(dato)
		iterador.actual = iterador.lista.ultimo
	} else {
		nuevo := crearNodo(dato)
		iterador.anterior.siguiente = nuevo
		nuevo.siguiente = iterador.actual
		iterador.actual = nuevo
		iterador.lista.largo++
	}
}

func (iterador *iteradorListaEnlazada[T]) VerActual() T {
	if !iterador.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return iterador.actual.dato
}

func (iterador *iteradorListaEnlazada[T]) HaySiguiente() bool {

	return iterador.actual != nil
}

func (iterador *iteradorListaEnlazada[T]) Siguiente() T {
	if !iterador.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	dato := iterador.VerActual()
	iterador.anterior = iterador.actual
	iterador.actual = iterador.actual.siguiente
	iterador.anterior.siguiente = iterador.actual
	return dato
}

func (iterador *iteradorListaEnlazada[T]) Borrar() T {
	if !iterador.HaySiguiente() {
		panic("El iterador termino de iterar")
	} else if iterador.actual == iterador.lista.primero {
		dato := iterador.lista.BorrarPrimero()
		iterador.actual = iterador.lista.primero
		return dato
	} else if iterador.actual.siguiente == nil {
		dato := iterador.actual.dato
		iterador.lista.ultimo = iterador.anterior
		iterador.actual = nil
		iterador.anterior.siguiente = nil
		iterador.lista.largo--
		return dato
	} else {
		dato := iterador.actual.dato
		sig := iterador.actual.siguiente
		iterador.actual.siguiente = nil
		iterador.anterior.siguiente = sig
		iterador.actual = sig
		iterador.lista.largo--
		return dato
	}
}
