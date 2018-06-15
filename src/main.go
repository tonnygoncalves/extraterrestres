package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/tonnygoncalves/extraterrestres/src/models"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		process()
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

// process get all info by day
func process() {
	models.ProcessPlanets()
}
