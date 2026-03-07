package api

import (
	"encoding/json"
	"lavaproxy/model"
	"net/http"
	"sync"
)

var (
	Nodes     []model.Node
	Mutex     sync.Mutex
	lastIndex int // <--- Este recordará el turno
)

func GetNodeHandler(w http.ResponseWriter, r *http.Request) {
	Mutex.Lock()
	defer Mutex.Unlock()

	// 1. Filtramos solo los nodos que están vivos
	var healthyNodes []*model.Node
	for i := range Nodes {
		if Nodes[i].Alive {
			healthyNodes = append(healthyNodes, &Nodes[i])
		}
	}

	// 2. Si no hay ninguno sano, avisamos
	if len(healthyNodes) == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusServiceUnavailable)
		json.NewEncoder(w).Encode(map[string]string{"error": "No hay nodos disponibles"})
		return
	}

	// 3. Lógica Round Robin
	// Si el índice se pasó del tamaño de la lista de sanos, vuelve a 0
	if lastIndex >= len(healthyNodes) {
		lastIndex = 0
	}

	selectedNode := healthyNodes[lastIndex]

	// 4. Avanzamos el turno para la próxima petición
	lastIndex = (lastIndex + 1) % len(healthyNodes)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(selectedNode)
}

func UpTimeHandler(w http.ResponseWriter, r *http.Request) {
	//logica para obtener una respueta json falsa para el uptime robot

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status": "alive",
		"uuid":   "f47ac10b-58cc-4372-a567-0e02b2c3d479",
	})
}