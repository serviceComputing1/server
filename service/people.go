package service

import(
	"net/http"
	"github.com/873421427/server/model"
	"encoding/json"
	"strconv"
	"github.com/gorilla/mux"
	"fmt"
	"github.com/873421427/server/swapi"
	"bytes"
	//"strings"
)

type responseFormat struct{
	Count int `json:"count"`
	Next string `json:"next"`
	Previous string `json:"previous"`
	Results []swapi.People `json:"results"`
}

//default return for /people
//only return 10 pages

var perPage=10

func GetPeople(w http.ResponseWriter, req * http.Request){
	page :=1
	var res []swapi.People

	var resData responseFormat

	//var resData map[string]string
	
	//upage := mux.Vars(req)["page"]
	//key := req.FormValue("page")	

	vals := req.URL.Query()
	//param,_ := vals["page"]
	search := vals.Get("search")
	pp := vals.Get("page")
	callbackName := vals.Get("callback")
	if callbackName == ""{
		fmt.Fprintf(w,"Please give callback name in query string")
		return 
	}

	//fmt.Printf("%T, %T\n",search,pp)


	if pp != ""{

		page1,err := strconv.Atoi(pp)

		if err !=nil {
			fmt.Println(err)
			NotFound(w,req) 
			return
		}
		page = page1
	}	
	
	//handle queries search with page or without search
	if search!=""{
		res= model.Search("people",search,(page-1)*perPage+1,perPage)
		resData.Count = model.GetTotalNumOfSearch("people",search)
		//resData.Next = NextPage(page,"search=")
		resData.Next= NextPageInSearch(page,search)
		fmt.Println(resData.Next)
		resData.Previous = PrePageInSearch(page,search)

	} else{
		res = model.GetPeople((page-1)*perPage + 1,perPage)
		resData.Count = model.GetTotalNumOfPeople()
		resData.Next = NextPage(page)
		resData.Previous = PrePage(page)
	}
	
	w.Header().Set("Content-Type","application/javascript")
	if res != nil{		

		resData.Results = res

		b,err := json.Marshal(resData)
		b= bytes.Replace(b,[]byte("\\u0026"),[]byte("&"),-1)
		if err != nil{
			fmt.Println(err)
			return
		}
		//w.Write(b)
		fmt.Fprintf(w,"%s(%s);",callbackName,b) // jsonp call
	} else {
		NotFound(w,req) 
	}
	
}



func GetPerson(w http.ResponseWriter, req * http.Request) {
	params := mux.Vars(req)
	vals := req.URL.Query()

	callbackName := vals.Get("callback")
	if callbackName == ""{
		fmt.Fprintf(w,"Please give callback name in query string")
		return 
	}

	id,err := strconv.Atoi(params["id"])

	if err != nil{
		NotFound(w,req)
		return
	}

	res := model.GetPeople(id,1)
	buf,err := json.Marshal(res)
	if err !=nil{
		fmt.Println(err)
	}
	w.Header().Set("Content-Type","application/javascript")

	if res == nil{
		NotFound(w,req)
		return
	}
	//w.Write(buf)
	fmt.Fprintf(w,"%s(%s);",callbackName,buf)  //jsonp
}



func NextPage(curpage int) string{
	if curpage * perPage < model.GetTotalNumOfPeople(){
		return "localhost:8080/people/?page="+ strconv.Itoa(curpage+1)
	}
	return ""
}

func PrePage(curpage int) string{
	if curpage != 1{
		return "localhost:8080/people/?page="+ strconv.Itoa(curpage-1)
	}
	return ""
}

func NextPageInSearch(curpage int,search string) string{
	if curpage*perPage < model.GetTotalNumOfSearch("people",search){
		return "localhost:8080/people/?search=" + search + "&page=" + strconv.Itoa(curpage+1)
	}
	return ""
}

func PrePageInSearch(curpage int,search string)string{
	if curpage!=1{
		return "localhost:8080/people/?search=" + search + "&page=" + strconv.Itoa(curpage-1)
	}
	return ""
}




/*
func dataToArray( data []byte) string{
	s := string(data)
	s1 := append([]string{"["},s,"]")
	return strings.Join(s1,"")
}
*/

/*

func GetPeopleInPage(w http.ResponseWriter, req * http.Request){
	params := mux.Vars(req)

	page,err := strconv.Atoi(params["id"])

	if err != nil{
		NotFound(w,req)
		return
	}

	res := model.GetPeople((page-1)*perPage + 1,perPage)
	w.Header().Set("Content-Type","application/json")

	if res == nil{
		NotFound(w,req)
		return
	}
	w.Write(res)
	
}
*/
