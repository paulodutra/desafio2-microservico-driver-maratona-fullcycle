package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

//struct(similar ao Model) para mapear os dados que vem do json
type Driver struct {
	Uuid string "json:uuid"
	Name string "json:name"
}
	
type Drivers struct {
	//slice (Array infinito, porque array no go tem tamanho definido)
	Drivers []Driver
}

//carrega os dados do arquvo json
func loadDrivers() []byte {
	
	jsonFile, err := os.Open("drivers.json")
	if err != nil {
		panic(err.Error())
	}

	defer jsonFile.Close()

	data, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		panic(err.Error())
	}

	return data

}

//obtem 1 driver por id
func getDriverById(w http.ResponseWriter, r *http.Request) {
	
	vars:= mux.Vars(r)
	data:= loadDrivers()

	var drivers Drivers
	//como se fosse o json encode do PHP
	json.Unmarshal(data, &drivers)

	for _, v := range drivers.Drivers {
		
		if v.Uuid == vars["id"] {
			//esse _ é caso de erro ignorar
			driver, _ := json.Marshal(v)
			w.Write([]byte(driver))
	
		}
	

	}


}

//lista todos os dados do json e escreve uma resposta http
func listDrivers(w http.ResponseWriter, r *http.Request) {
	drivers := loadDrivers()
	w.Write([]byte(drivers))
}

//cria o servidor na porta 8081 e registra as rotas
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/drivers", listDrivers)
	//endpoint de driver por id
	r.HandleFunc("/drivers/{id}", getDriverById)
	http.ListenAndServe(":8081", r)
}