package restserver

// En esta seccion se indican los paquetes que pueden ser importados al proyecto
import (
    "net/http"
    "strings"
    "log"
)

// Variable global que almacena los recursos disponibles
var recursos = make(map[string]interface{})

// A continuacion crearemos variables TYPE en las que se definen las funciones
// interface que necesitaremos para administrar los request.

// Devuelve una lista de todos los items del recurso
// GET /recurso/
type listado interface {
    Listar(http.ResponseWriter, *http.Request)
}

// Crea un nuevo recurso
// POST /recurso/
type crear interface {
    Crear(http.ResponseWriter, *http.Request)
}

// Devuelve un recurso de acuerdo al parametro string
// GET /recurso/id 
type buscar interface {
    Buscar(http.ResponseWriter, string, *http.Request)
}

// Actualiza un recurso con el parametro id
// PUT /recurso/id
type actualizar interface {
    Actualizar(http.ResponseWriter, string, *http.Request)
}

// Elimina un recurso con el codigo id
// DELETE /recurso/id
type eliminar interface {
    Eliminar(http.ResponseWriter, string, *http.Request)
}

// Obtiene la configuracion del servidor, esto es util cuando el REST server es
// distinto en dominio al client (CORS).
// OPTIONS /recurso/            si el string es nil (nulo)
// OPTIONS /recurso/id          si el string es valido
type opciones interface {
    Opciones(http.ResponseWriter, string, *http.Request)
}


// La siguiente funcion manejara el ruteo de las peticiones HTTP y las transformara
// en resquest a los diferentes verbos de REST
func manejarPeticion(p http.ResponseWriter, req *http.Request) {

    // Obtiene el nombre del recurso y el id
    var finalRecurso = strings.Index(req.URL.Path[1:], "/") + 1
    var nombreRecurso string
    if finalRecurso == -1 {
        nombreRecurso = req.URL.Path[1:]
    } else {
        nombreRecurso = req.URL.Path[1:finalRecurso]
    }
    var id = req.URL.Path[finalRecurso + 1:]

    log.Println(nombreRecurso)

    // Inicializa el Recurso Solicitado
    recurso, valido := recursos[nombreRecurso]



    // Comprobar que el nombre de recurso requerido sea valido
    if !valido {
        SinImplementar(p)
    }

    //log.Println("Valido")
    
    // Enrutador de Requerimientos por Method
    switch req.Method {
    // Procesar GET
    case "GET":
        if len(id) == 0 {
            if recGet, valido := recurso.(listado); valido {
                recGet.Listar(p, req)
            } else {
                log.Println("No encuentra")
                SinImplementar(p)
            }
        } else {
            if recGet, valido := recurso.(buscar); valido {
                recGet.Buscar(p, id, req)
            } else {
                SinImplementar(p)
            }
        }
    // Procesar POST
    case "POST":
        if recPost, valido := recurso.(crear); valido {
            recPost.Crear(p, req)
        } else {
            SinImplementar(p)
        }
    // Procesar PUT
    case "PUT":
        if recPut, valido := recurso.(actualizar); valido {
            recPut.Actualizar(p, id, req)
        } else {
            SinImplementar(p)
        }
    // Procesar DELETE
    case "DELETE":
        if recDelete, valido := recurso.(eliminar); valido {
            recDelete.Eliminar(p, id, req)
        } else {
            SinImplementar(p)
        }
    // Procesar OPTIONS
    case "OPTIONS":
        if recOptions, valido := recurso.(opciones); valido {
            recOptions.Opciones(p, id, req)
        } else {
            SinImplementar(p)
        }
    default:
        SinImplementar(p)
    }

}

// Define las funciones usadas para registar recursos, administrar respuestas e implementaciones

// Devuelve el error 404 cuando no se ha implementado el recurso
func SinImplementar(p http.ResponseWriter) {
    http.Error(p, "404 Not Found", http.StatusNotFound)
}

// Agrega un recurso
func registrarRecurso (nombre string, rec interface{}) {
   recursos[nombre] = rec
   http.Handle("/"+nombre+"/", http.HandlerFunc(manejarPeticion))
}
