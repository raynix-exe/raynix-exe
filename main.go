package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Resposta struct {
	Nome    string `json:"nome"`
	Horario string `json:"horario"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/projeto-korp" {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	resposta := Resposta{
		Nome:    "Projeto Korp",
		Horario: time.Now().UTC().Format(time.RFC3339),
	}
	json.NewEncoder(w).Encode(resposta)
}

func main() {
	// Sua rota original
	http.HandleFunc("/projeto-korp", handler)

	// A nova rota que o Prometheus vai ler
	http.Handle("/metrics", promhttp.Handler())

	// Inicializa o servidor na porta 8080
	http.ListenAndServe(":8080", nil)
}

