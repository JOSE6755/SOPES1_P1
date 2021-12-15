package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
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

type Mensaje struct {
	Prueba string `json:"Prueba"`
}

type Memoria struct {
	Total      string `json:"Total"`
	Consumida  string `json:"Consumida"`
	ConsumidaP string
}

var (
	wsUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func homepage(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	wsUpgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	wsConn, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("NO se pudo actualizar: %s\n", err.Error())
		return
	}
	out, err := ioutil.ReadFile("/proc/Modulo de RAM")
	if err != nil {
		log.Fatal(err)

	}
	var memo Memoria
	json.Unmarshal(out, &memo)
	guard, err := strconv.ParseFloat(memo.Total, 64)
	if err != nil {
		fmt.Println(err.Error())
	}
	s := fmt.Sprintf("%f", guard/1024)
	memo.Total = s
	//var mensaje string

	go func() {
		for {
			_, message, err := wsConn.ReadMessage()
			if err != nil {
				fmt.Println(string(message))
				fmt.Printf("Ocurrio un error: %s", err.Error())
				break
			}

		}
	}()

	go func() {
		for {
			w.Header().Set("Content-Type", "application-json")
			out, err := ioutil.ReadFile("/proc/Modulo de CPU")
			if err != nil {
				log.Fatal(err)
				break
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
				memo, err := strconv.ParseFloat(memo.Total, 64)
				if err != nil {
					fmt.Println(err.Error())
				}
				cambio, err := strconv.ParseFloat(procesos.Datos[i].RAM, 64)
				if err != nil {
					fmt.Println(err.Error())
				}
				s := fmt.Sprintf("%f", (cambio/memo)*100)
				procesos.Datos[i].RAM = s
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

			}
			num_tot.Ejecucion = ejecucion
			num_tot.Detenidos = detenidos
			num_tot.Suspendidos = suspendidos
			num_tot.Total = total
			num_tot.Zombies = zombies
			procesos.Num_proce = num_tot
			//x := json.NewEncoder(w).Encode(procesos)
			error := wsConn.WriteJSON(procesos)
			if error != nil {
				log.Println(error)

			}
			fmt.Println("Adios")
			time.Sleep(3 * time.Second)
		}
	}()

}

func CPU(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	wsUpgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	wsConn, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("NO se pudo actualizar: %s\n", err.Error())
		return
	}

	go func() {
		for {
			_, message, err := wsConn.ReadMessage()
			if err != nil {
				fmt.Println(string(message))
				fmt.Printf("Ocurrio un error: %s", err.Error())
				break

			}
		}
	}()

	go func() {
		for {
			cmd := exec.Command("sh", "-c", "ps -eo pcpu | sort | sort -k 1 -r | head -50")
			out, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Println(err.Error())
			}
			output := string(out[:])
			pruebas := strings.Split(output, "\n")
			var suma float64
			for i := 1; i < len(pruebas); i++ {
				s := strings.TrimSpace(pruebas[i])
				aux, err := strconv.ParseFloat(s, 64)
				if err != nil {
					fmt.Println(err)
				}
				if aux == 0 {
					break
				}
				suma += aux
				//fmt.Println(pruebas[i])
				//fmt.Println(pruebas[i])
			}
			fmt.Println(suma / 8)
			s := fmt.Sprintf("%f", suma/8)
			error := wsConn.WriteMessage(websocket.TextMessage, []byte(s))
			if error != nil {
				log.Println(error)

			}
			fmt.Println("Adios")
			time.Sleep(1 * time.Second)
		}
	}()

}

func RAM(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	wsUpgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	wsConn, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("NO se pudo actualizar: %s\n", err.Error())
		return
	}

	go func() {
		for {
			_, message, err := wsConn.ReadMessage()
			if err != nil {
				fmt.Println(string(message))
				fmt.Printf("Ocurrio un error: %s", err.Error())
				break

			}
		}
	}()

	go func() {
		for {
			w.Header().Set("Content-Type", "application-json")
			out2, err2 := ioutil.ReadFile("/proc/Modulo de RAM")
			if err2 != nil {
				fmt.Println(err2.Error())
			}
			var memoinf Memoria
			json.Unmarshal(out2, &memoinf)
			cmd := exec.Command("sh", "-c", "free -m | head -2 | tail -1 | awk '{print $6}' ")
			out, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Println(err.Error())
			}
			output := strings.Split(string(out[:]), "\n")
			pruebas, err := strconv.ParseFloat(output[0], 64)
			if err != nil {
				fmt.Println(err)
			}

			r, erro := strconv.ParseFloat(memoinf.Consumida, 64)
			if erro != nil {
				fmt.Println(erro)
			}
			resta := r - pruebas
			memoinf.Consumida = fmt.Sprintf("%f", resta)
			r, erro = strconv.ParseFloat(memoinf.Total, 64)
			if erro != nil {

				fmt.Println(erro)
			}
			resta = (resta / r) * 100
			memoinf.ConsumidaP = fmt.Sprintf("%f", resta)

			//s := fmt.Sprintf("%f", suma/8)
			error := wsConn.WriteJSON(memoinf)
			if error != nil {
				log.Println(error)

			}
			fmt.Println("Adios")
			time.Sleep(1 * time.Second)
		}
	}()

}

//free -m | head -2 | tail -1 | awk '{print $6}'

func main() {
	fmt.Println("es mi servidor we")

	router := mux.NewRouter()
	router.HandleFunc("/", homepage).Methods("GET")
	router.HandleFunc("/CPU", CPU).Methods("GET")
	router.HandleFunc("/RAM", RAM).Methods("GET")
	log.Fatal(http.ListenAndServe(":9100", router))
}
