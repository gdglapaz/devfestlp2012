package restserver

// la funcion INIT es la que inicializa la aplicacion en AppEngine
func init() {
	misUsuarios := NuevoUsuarioCollection()
	misUsuarios.Agregar("qennix")
	misUsuarios.Agregar("devfestlp2012")
	RegistrarRecurso("usuario", misUsuarios)
}