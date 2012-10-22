package restserver

import(
  "net/http"
  "log"
  "fmt"
)

// Definicion del Objeto Usuario
type Usuario struct {
	Id int
	Nick string
}

// Constructor del Objecto Usuario
func NuevoUsuario(id int, nick string) *Usuario {
	return &Usuario{id, nick}
}

// Definicion de la colleccion de Usuarios
type UsuarioCollection struct {
	usuarios map[int]string
	proximoId int
}

// Constructor de la colleccion de usuarios
func NuevoUsuarioCollection() *UsuarioCollection {
	return &UsuarioCollection {make(map[int]string), 0}
}


// Crear el Metodo Agregar al Objeto UsuarioCollection
func (u *UsuarioCollection) Agregar(nick string) int {
	id := u.proximoId;
	u.proximoId++
	u.usuarios[id] = nick
	return id
}

// Crear el Metodo Buscar al Objeto UsuarioCollection
func (u *UsuarioCollection) BuscarId(id int) (*Usuario, bool) {
	nick, valido := u.usuarios[id]
	if !valido {
		return nil, false
	}
	return &Usuario{id, nick}, true
}

// Crear el Metodo ListarTodos al Objeto UsuarioCollection
func (u *UsuarioCollection) ListarTodos() []*Usuario {
	todos := make([]*Usuario, len(u.usuarios))
	for id, nick := range u.usuarios {
		todos[id] = &Usuario{id, nick}
	}
	return todos
}

// Crear el Metodo Eliminar al Objeto UsuarioCollection
func (u *UsuarioCollection) Eliminar(id int) {
	delete(u.usuarios, id);
}


func (u *UsuarioCollection) Listar(p http.ResponseWriter, req *http.Request) {
	log.Println("LLego Aca")
	for _, usuario := range u.ListarTodos() {
		fmt.Fprintf(p, "<a href=\"%v\">%v</a>%v<br />", usuario.Id, usuario.Id, usuario.Nick)
	}
}
