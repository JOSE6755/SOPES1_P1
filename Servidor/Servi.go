package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"io/ioutil"

	"github.com/gorilla/mux"
)

type Todo struct {
	Datos     []Padre `json:"root"`
	Num_proce Num_Proc
}

type Padre struct {
	Proceso string `json:"Proceso"`
	PID     string `json:"PID"`
	RAM     string `json:"RAM"`
	Usuario string `json:"Usuario"`
	Estado  string `json:"Estado"`
	Hijos   []Kids `json:"hijos"`
}

type Kids struct {
	Proceso string `json:"Proceso"`
	PID     string `json:"PID"`
	Estado  string `json:"Estado"`
}

type Num_Proc struct {
	Ejecucion   int
	Suspendidos int
	Detenidos   int
	Zombies     int
	Total       int
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func homepage(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(http.StatusCreated)
	out, err := ioutil.ReadFile("/proc/Modulo de CPU")
	if err != nil {
		log.Fatal(err)
	}

	var procesos Todo
	var num_tot Num_Proc
	json.Unmarshal(out, &procesos)
	ejecucion := 0
	suspendidos := 0
	detenidos := 0
	zombies := 0
	total := 0
	for i := 0; i < len(procesos.Datos); i++ {
		if procesos.Datos[i].Estado == "0" {
			ejecucion += 1
		} else if procesos.Datos[i].Estado == "1" || procesos.Datos[i].Estado == "1026" {
			suspendidos += 1
		} else if procesos.Datos[i].Estado == "4" {
			zombies += 1
		} else {
			detenidos += 1
		}
		total += 1
		for j := 0; j < len(procesos.Datos[i].Hijos); j++ {
			if procesos.Datos[i].Hijos[j].Estado == "0" {
				ejecucion += 1
			} else if procesos.Datos[i].Hijos[j].Estado == "1" || procesos.Datos[i].Estado == "1026" {
				suspendidos += 1
			} else if procesos.Datos[i].Hijos[j].Estado == "4" {
				zombies += 1
			} else {
				detenidos += 1
			}
			total += 1
		}
	}
	num_tot.Ejecucion = ejecucion
	num_tot.Detenidos = detenidos
	num_tot.Suspendidos = suspendidos
	num_tot.Total = total
	num_tot.Zombies = zombies
	procesos.Num_proce = num_tot
	json.NewEncoder(w).Encode(procesos)

}

func main() {
	fmt.Println("es mi servidor we")
	router := mux.NewRouter()
	router.HandleFunc("/", homepage).Methods("GET")
	log.Fatal(http.ListenAndServe(":9101", router))
}
