package restserver

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
func (u *UsuarioCollection) EliminarId(id int) {
	delete(u.usuarios, id);
}
