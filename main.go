package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"fmt"
	"io/ioutil"
)

type Hostings struct{
	Hostings []host `json:"Hostings"`
}

type host struct{
	Id	string	`json:"Id"`
	Name	string	`json:"Nombre"`
	Cores	string	`json:"Cores"` 
	Mem	string	`json:"Memoria"`
	Disk	string	`json:"Disco"`
}

var h Hostings

//Listar todo
func GetHostings(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(h.Hostings)
}

//Listar un elemento
func GetHost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range h.Hostings {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&host{})
}

//Eliminar elemento
func DeleteHost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range h.Hostings {
		if item.Id == params["id"] {
			h.Hostings = append(h.Hostings[:index], h.Hostings[index+1:]...)
			break
		}else {fmt.Println("No existe el ID")}
	}
	json.NewEncoder(w).Encode(h.Hostings)
}

//Crear elemento
func CreateHost (w http.ResponseWriter, r *http.Request) {
	var newhost host
	_ = json.NewDecoder(r.Body).Decode(&newhost)
	h.Hostings = append(h.Hostings, newhost)
	json.NewEncoder(w).Encode(h.Hostings)
}

//Editar elemento
func EditHost (w http.ResponseWriter, r *http.Request) {
	var edit host
	_ = json.NewDecoder(r.Body).Decode(&edit)
	for index, item := range h.Hostings {
		if item.Id == edit.Id {
			h.Hostings[index].Id=edit.Id
			if edit.Name!="" {h.Hostings[index].Name=edit.Name}
			if edit.Cores!="" {h.Hostings[index].Cores=edit.Cores}
			if edit.Mem!="" {h.Hostings[index].Mem=edit.Mem}
			if edit.Disk!="" {h.Hostings[index].Disk=edit.Disk}
			break
		}else {fmt.Println("No existe el ID")}
	}
	json.NewEncoder(w).Encode(h.Hostings)
}

func main(){
	//Lectura del fichero
	byteValue, e := ioutil.ReadFile("hostings.json")
	if e != nil{
		fmt.Println(e)
	}
	json.Unmarshal(byteValue, &h)

	//Creación de rutas
	router := mux.NewRouter()
	router.HandleFunc("/hosting",GetHostings).Methods("GET")
	router.HandleFunc("/hosting/{id}", GetHost).Methods("GET")
	router.HandleFunc("/hosting/{id}",  DeleteHost).Methods("DELETE")
	router.HandleFunc("/hosting/{id}", CreateHost).Methods("POST")
	router.HandleFunc("/hosting/{id}", EditHost).Methods("PATCH")
	log.Fatal(http.ListenAndServe(":8000",router))
}
