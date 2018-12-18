package service

import(
	"encoding/json"
	"net/http"
	"fmt"
)

func NotFound(w http.ResponseWriter, req * http.Request){
	vals := req.URL.Query()
	callbackName := vals.Get("callback")
	if callbackName == ""{
		fmt.Fprintf(w,"Please give a callback name in query string.")
		return 
	}

	response := map[string]string{"detail":"Not found"}
	rm,_ := json.Marshal(response)
	//w.Write(rm)
	fmt.Fprintf(w,"%s(%s)",callbackName,rm)
}