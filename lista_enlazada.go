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
/*
func CrearListaEnlazada[T any]() Lista[T] {
	//...
}
*/
