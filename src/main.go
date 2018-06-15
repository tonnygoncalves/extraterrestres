package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var (
	vTime          int //10 years
	maxperimetro   float64
	diamaximo      int
	contadorSequia int64 // contador para saber cuantos días de sequía hay en el periodo a evaluar
	contadorLluvia int64 // contador para saber cuantos días de lluvia hay en el periodo a evaluar
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Ha ingresado al el API del sistema solar DELTA si no conoce nuestra civilización por favor no continue.")
	})

	http.HandleFunc("/clima", handleGetData)

	http.ListenAndServe(":9990", nil)
}

//handleGetData get clima router info
func handleGetData(w http.ResponseWriter, r *http.Request) {

	keys, ok := r.URL.Query()["dia"]

	if !ok || len(keys) < 1 {
		log.Println("Url Param 'dia' is missing")
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte("No ha colocado el parametro 'dia'!"))
		return
	}

	if _, err := strconv.Atoi(keys[0]); err != nil {
		log.Println("%q - It looks like not a number.\n", keys[0])
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte("No ha ingresado un numero.\n"))
		return
	}
	key := keys[0]
	type Day struct {
		dia   string
		clima string
	}
	group := Day{
		dia:   key,
		clima: "Reds",
	}

	js, err := json.Marshal(group)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func process() {
	vTime = 3600 //10 years
}
