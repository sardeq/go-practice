package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/sardeq/go-practice/SomeAPI/internal/handlers"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting GO API...")
	//from https://edukits.co/text-art/
	fmt.Println(`
  ____   ___  __  __ _____      _    ____ ___
 / ___| / _ \|  \/  | ____|    / \  |  _ \_ _|
 \___ \| | | | |\/| |  _|     / _ \ | |_) | |
  ___) | |_| | |  | | |___   / ___ \|  __/| |
 |____/ \___/|_|  |_|_____| /_/   \_\_|  |___|`)

	err := http.ListenAndServe("localhost:8000", r)
	if err != nil {
		log.Error(err)
	}
}
