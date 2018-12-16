package service

import(
	"encoding/json"
	"net/http"
)

func NotFound(w http.ResponseWriter, req * http.Request){
	response := map[string]string{"detail":"Not found"}
	rm,_ := json.Marshal(response)
	w.Write(rm)
}