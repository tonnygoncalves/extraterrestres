package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/tonnygoncalves/extraterrestres/src/models"
)

func main() {
	http.HandleFunc("/clima", handleGetData)
	//Abriendo el puerto donde se escucha.
	http.ListenAndServe(":80", nil)
}

//handleGetData get clima router info
func handleGetData(w http.ResponseWriter, r *http.Request) {

	keys, ok := r.URL.Query()["dia"]

	if !ok || len(keys) < 1 {
		log.Println("Url Param 'dia' is missing")
		type ErrBack struct {
			Error string `json:"error"`
		}
		group := ErrBack{
			Error: "No ha colocado el parametro 'dia'!",
		}
		js, errError1 := json.Marshal(group)
		if errError1 != nil {
			http.Error(w, errError1.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
		return
	}

	if _, err := strconv.Atoi(keys[0]); err != nil {
		log.Println("%q - It looks like not a number.\n", keys[0])
		type ErrBack struct {
			Error string `json:"error"`
		}
		group := ErrBack{
			Error: "No ha ingresado un numero.",
		}
		js, errError1 := json.Marshal(group)
		if errError1 != nil {
			http.Error(w, errError1.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
		return
	}
	key := keys[0]
	type Day struct {
		Dia   string `json:"dia"`
		Clima string `json:"clima"`
	}

	i, err := strconv.Atoi(key)
	var wheader = models.GetDayWheather(i)
	group := Day{
		Dia:   key,
		Clima: wheader,
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
