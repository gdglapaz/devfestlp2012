package restserver

// En esta seccion se indican los paquetes que pueden ser importados al proyecto
import(
  "log"
)
// la funcion INIT es la que inicializa la aplicacion en AppEngine
func init() {
	misUsuarios := NuevoUsuarioCollection()
	log.Println(misUsuarios.Agregar("qennix"))
	
	log.Println(misUsuarios.Agregar("devfestlp2012"))
	
	registrarRecurso("usuario", misUsuarios)
}