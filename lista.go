package lista

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

	// CREAMOS EL ITERADOR - DESCOMENTAR LUEGO
	//Iterador() IteradorLista[T]
}
