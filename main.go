package main

import (
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"lavaproxy/api"
	"lavaproxy/model"

	"gopkg.in/yaml.v3"
)

func main() {
	// 1. Cargar la configuración desde el archivo YAML
	configFile, err := os.ReadFile("config.yml")
	if err != nil {
		log.Fatalf("Error: No se pudo leer config.yml: %v", err)
	}

	var config model.Config
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		log.Fatalf("Error: Formato de YAML inválido: %v", err)
	}

	// Pasamos los nodos cargados a la variable global de la API
	api.Nodes = config.Nodos
	log.Printf("Cargados %d nodos desde config.yml", len(api.Nodes))

	// 2. Iniciar el Health Check en una Goroutine (segundo plano)
	go func() {
		for {
			var fallos int

			for i := range api.Nodes {
				api.Mutex.Lock()
				host := api.Nodes[i].Host
				port := api.Nodes[i].Port
				wasAlive := api.Nodes[i].Alive
				api.Mutex.Unlock()

				hostPort := net.JoinHostPort(host, port)
				
				// Intentamos una conexión TCP rápida (timeout de 2s)
				conn, err := net.DialTimeout("tcp", hostPort, 2*time.Second)
				
				api.Mutex.Lock()
				if err != nil {
					if wasAlive {
						log.Printf("Nodo CAÍDO: %s", host)
					} else {
						log.Printf("Nodo sigue con error: %s", host)
					}
					api.Nodes[i].Alive = false
					fallos++
				} else {
					if !wasAlive {
						log.Printf("Nodo ACTIVO: %s", host)
					}
					api.Nodes[i].Alive = true
					conn.Close()
				}
				api.Mutex.Unlock()
			}
			
			if fallos == 0 {
				log.Printf("Todos los %d nodos funcionaron correctamente", len(api.Nodes))
			} else {
				log.Printf("Estado: %d de %d nodos con error", fallos, len(api.Nodes))
			}

			// Esperamos el tiempo de delay configurado
			delay := time.Duration(config.CheckDelay) * time.Second
			if delay <= 0 {
				delay = 30 * time.Second // por defecto 30 segundos
			}
			time.Sleep(delay)
		}
	}()

	// 3. Configurar las rutas y arrancar el servidor HTTP
	http.HandleFunc("/", api.GetNodeHandler)
	http.HandleFunc("/up-time", api.UpTimeHandler)
	
	port := ":" + config.Puerto
	log.Printf("Discovery Server corriendo en http://localhost%s", port)
	
	server := &http.Server{
		Addr:         port,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Error al iniciar el servidor: ", err)
	}
}