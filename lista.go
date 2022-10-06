package lista

type IteradorLista[T any] interface {
	//Hace que el iterador apunte al siguiente nodo
	Siguiente() T
	// Devuelve true si el siguiente nodo no apunta a nil. Devuelve false en caso contrario
	HaySiguiente() bool
	//Devuelve el dato del nodo al cual el iterador esta apuntando
	VerActual() T
	//Recibe un dato e inserta un nodo con el dato recibido entre el nodo anterior del iterador y el actual. El iterador termina apuntando al nuevo nodo
	Insertar(T)
	// Se borra el nodo al cual el iterador está apuntando. El iterador termina apuntando al nodo siguiente
	Borrar() T
}

type Lista[T any] interface {

	// EstaVacia() devuelve true si no hay elementos en la lista, false en caso contrario
	EstaVacia() bool

	//Inserta un elemento en la primer posición de la lista
	InsertarPrimero(T)

	//Inserta un elemento en la última posición de la lista
	InsertarUltimo(T)

	//Borra el primer elemento de la lista y lo devuelve
	BorrarPrimero() T

	//Devuelve el primer elemento de la lista
	VerPrimero() T

	//Devuelve el último elemento de la lista
	VerUltimo() T

	//Devuelve la cantidad de elementos que hay en la lista
	Largo() int

	//Itera por todos los elementos de la lista
	Iterar(visitar func(T) bool)

	Iterador() IteradorLista[T]
}
