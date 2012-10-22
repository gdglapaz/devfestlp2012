package restserver

// En esta seccion se indican los paquetes que pueden ser importados al proyecto
import (
  "fmt"
  "net/http"
  "io/ioutil"
  "strconv"
)

// Funcion para el listado de todos los items del recurso
// GET /recurso/
func (u *UsuarioCollection) Listar(p http.ResponseWriter, req *http.Request) {
	for _, usuario := range u.ListarTodos() {
		fmt.Fprintf(p, "<a href=\"%v\">%v</a>%v<br />", usuario.Id, usuario.Id, usuario.Nick)
	}
}

// Funcion para crear un recurso nuevo
// POST /recurso/
func (u *UsuarioCollection) Crear(p http.ResponseWriter, req *http.Request) {
	datos, valido := ioutil.ReadAll(req.Body)
	if valido != nil {
		RequerimientoInvalido(p, "El requerimiento es invalido")
		return
	}
	id := u.Agregar(string(datos))
	RecursoCreado(p, fmt.Sprintf("%v%v", req.URL.String(), id))
}

// Funcion para listar o buscar un recurso especifico
// GET /recurso/1
func (u *UsuarioCollection) Buscar(w http.ResponseWriter, idString string, req *http.Request) {
	id, valido := strconv.Atoi(idString)
	if valido != nil {
		NoEncontrado(w)
		return
	}
	usuario, ok := u.BuscarId(id)
	if !ok {
		NoEncontrado(w)
		return
	}
	fmt.Fprintf(w, "<h1>Usuario %v</h1><p>%v</p>", usuario.Id, usuario.Nick)
}

// Funcion para actualizar un recurso especifico
// PUT /recurso/1
func (u *UsuarioCollection) Actualizar(w http.ResponseWriter, idString string, req *http.Request) {
	id, valido := strconv.Atoi(idString)
	if valido != nil {
		NoEncontrado(w)
		return
	}
	var usuario *Usuario
	var ok bool
	usuario, ok = u.BuscarId(id)
	if !ok {
		NoEncontrado(w)
	}
	var err error
	var datos []byte
	datos, err = ioutil.ReadAll(req.Body)
	if err != nil {
		RequerimientoInvalido(w, "El requerimiento es invalido")
		return
	}
	usuario.Nick = string(datos)
	RecursoActualizado(w, req.URL.String())

}

// Funcion que elimina un recurso
// DELETE /recurso/id
func (u *UsuarioCollection) Eliminar(w http.ResponseWriter, idString string, req *http.Request) {
	var id int
	var err error
	if id, err = strconv.Atoi(idString); err != nil {
	    NoEncontrado(w)
    }
	u.EliminarId(id)
	SinContenido(w)
}
