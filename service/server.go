package service


import(
	//"net/http"
	//"os"
	//"log"
	"github.com/codegangsta/negroni"
    "github.com/gorilla/mux"
    //"github.com/unrolled/render"
)

func NewServer () * negroni.Negroni{

	n := negroni.Classic()

	mx := mux.NewRouter()

	initRoutes(mx)

	n.UseHandler(mx)

	return n;
}

func initRoutes(mx *mux.Router) {
	//mx.HandleFunc("/people/?search={id}",SearchPeople).Methods("GET")
	mx.HandleFunc("/people/",GetPeople).Methods("GET")
	mx.HandleFunc("/people/{id}",GetPerson).Methods("GET")
	mx.HandleFunc("/api",GetAllApi).Methods("GET")
	//mx.Path("/people/").Queries("page","{page}").HandlerFunc(GetPeople).Name("GetPeople").Methods("GET")

}