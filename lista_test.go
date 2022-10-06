package lista_test

import (
	"github.com/stretchr/testify/require"
	TDALista "lista"
	"testing"
)

const VALOR_PRUEBA = 1
const DIVISOR_PRUEBA = 3

func TestListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()
	require.NotNil(t, lista)
	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
	require.EqualValues(t, 0, lista.Largo())
}

func TestListaInts(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i < VALOR_PRUEBA; i++ {
		lista.InsertarUltimo(i)
		require.EqualValues(t, i, lista.VerUltimo())
		require.EqualValues(t, i+1, lista.Largo())
		require.False(t, lista.EstaVacia())
	}
	require.EqualValues(t, 0, lista.VerPrimero())
	for i := 0; i < VALOR_PRUEBA; i++ {
		require.EqualValues(t, i, lista.VerPrimero())
		require.EqualValues(t, VALOR_PRUEBA-i, lista.Largo())
		lista.BorrarPrimero()
	}
	require.True(t, lista.EstaVacia())
}

func TestIteradorInterno(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	valor_comparacion := 0
	for i := 0; i < VALOR_PRUEBA; i++ {
		lista.InsertarUltimo(i)
		valor_comparacion += i
	}
	contador := 0
	contador_ptr := &contador
	lista.Iterar(func(elemento int) bool {
		*contador_ptr = *contador_ptr + elemento
		return true
	})
	require.EqualValues(t, valor_comparacion, contador)
}

func TestIteradorInternoTope(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	valor_comparacion := 0
	for i := 0; i < VALOR_PRUEBA; i++ {
		lista.InsertarUltimo(i)
		if i%DIVISOR_PRUEBA != 0 {
			valor_comparacion += i
		}
	}
	contador := 0
	contador_ptr := &contador
	lista.Iterar(func(elemento int) bool {
		if elemento%DIVISOR_PRUEBA == 0 {
			return false
		} else {
			*contador_ptr += elemento
			return true
		}
	})
	require.EqualValues(t, valor_comparacion, contador)
}

func TestIteradorExterno(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iterador := lista.Iterador()
	require.NotNil(t, iterador)
	require.False(t, iterador.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.Siguiente() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.Borrar() })
	iterador.Insertar(VALOR_PRUEBA)
	require.EqualValues(t, VALOR_PRUEBA, lista.VerPrimero())
	iterador.Insertar(VALOR_PRUEBA + 1)
	require.EqualValues(t, VALOR_PRUEBA+1, lista.VerPrimero())
	iterador.Borrar()
	require.EqualValues(t, VALOR_PRUEBA, lista.VerPrimero())
	require.EqualValues(t, VALOR_PRUEBA, iterador.Borrar())
	for i := 0; i < VALOR_PRUEBA; i++ {
		iterador.Insertar(i)
		iterador.Siguiente()
		if i%2 != 0 {
			iterador.Borrar()
			require.EqualValues(t, i, lista.VerUltimo())
		}
	}
}

func TestTresElementos(t *testing.T) {
	t.Log("Hacemos pruebas con listas de tres elementos")
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i < 3; i++ {
		lista.InsertarUltimo(i)
	}
	iterador := lista.Iterador()
	require.EqualValues(t, 0, iterador.VerActual())
	require.True(t, iterador.HaySiguiente())
	iterador.Siguiente()
	require.EqualValues(t, 1, iterador.VerActual())
	require.True(t, iterador.HaySiguiente())
	iterador.Siguiente()
	require.EqualValues(t, 2, iterador.VerActual())
}

func TestBorrarUltimo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i < 5; i++ {
		lista.InsertarUltimo(i)
	}
	iterador := lista.Iterador()
	require.EqualValues(t, 0, iterador.VerActual())
	iterador.Siguiente()
	require.EqualValues(t, 1, iterador.VerActual())

	iterador.Siguiente()
	require.EqualValues(t, 2, iterador.VerActual())

	iterador.Siguiente()
	require.EqualValues(t, 3, iterador.VerActual())

	iterador.Siguiente()
	require.EqualValues(t, 4, iterador.Borrar())
	require.EqualValues(t, 3, lista.VerUltimo())

}
