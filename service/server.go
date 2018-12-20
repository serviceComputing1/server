package service


import(
	"net/http"
	"encoding/json"
	"fmt"
	//"github.com/serviceComputing1/server/routes"
    //"github.com/unrolled/render"
)

/*
func NewServer () * negroni.Negroni{

	n := negroni.Classic()

	mx := routes.NewRouter()

	//initRoutes(mx)

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
*/
func GetAllApi(w http.ResponseWriter, req * http.Request){
	var res1 map[string]string
	res1 = make(map[string]string)
	res1["people"] = "localhost:8080/people/"
	res1["films"] = "localhost:8080/films/"
	res1["vehicles"] =	"localhost:8080/vehicles/"
	res1["planets"] = "localhost:8080/planets"
	res1["species"] = "localhost:8080/species"
	res1["starships"] = "localhost:8080/starships"

	b,err := json.Marshal(res1)
	if err !=nil{
		fmt.Println(err)
		NotFound(w,req)
		return 
	}


	w.Header().Set("Content-Type","application/json")
	w.Write(b)
}

func GetIndex(w http.ResponseWriter, req * http.Request) {
		http.ServeFile(w, req, "./dist/index.html")	
		fmt.Println("index request ")
}

