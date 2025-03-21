package main

import (
	"fmt"
	"net/http"

	"github.com/ZaxVaxZ/GoAPI/internal/handlers"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	fmt.Print("Setting Up Service Handlers...\n")
	handlers.Handler(r)
	fmt.Print("Starting GO API Service Listener...\n")
	err := http.ListenAndServe("localhost:8000", r)
	if err != nil {
		log.Error(err)
	}
}
