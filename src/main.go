package main

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
	"strconv"
)

var (
	vTime          int     = 3600 //10 years
	maxperimetro   float64 = 0
	diamaximo      int     = 0
	contadorSequia int64   = 0 // contador para saber cuantos días de sequía hay en el periodo a evaluar
	contadorLluvia int64   = 0 // contador para saber cuantos días de lluvia hay en el periodo a evaluar
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Ha ingresado al el API del sistema solar DELTA si no conoce nuestra civilización por favor no continue.")
	})

	http.HandleFunc("/clima", handleGetData)

	http.ListenAndServe(":9990", nil)
}

func handleGetData(w http.ResponseWriter, r *http.Request) {

	keys, ok := r.URL.Query()["dia"]

	if !ok || len(keys) < 1 {
		log.Println("Url Param 'dia' is missing")
		return
	}

	if _, err := strconv.Atoi(""); err == nil {
		fmt.Printf("%q looks like a number.\n", keys)
	}

	// Query()["key"] will return an array of items,
	// we only want the single item.
	key := keys[0]

	log.Println("Url Param 'dia' is: " + string(key))
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))

	type Day struct {
		Dia    int
		Tiempo string
	}
	group := Day{
		Dia:    1,
		Tiempo: "Reds",
	}

	js, err := json.Marshal(group)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
