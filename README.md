# LavaProxy

LavaProxy es un servidor de descubrimiento ligero para nodos de [Lavalink](https://github.com/lavalink-devs/Lavalink). Está diseñado para monitorear una lista de nodos Lavalink, comprobar periódicamente si están en línea (a través de una verificación TCP), y proporcionar una API HTTP simple para obtener un nodo que esté activo y listo para ser utilizado.

## Características

*   **Comprobación de salud en segundo plano:** Verifica el estado de los nodos cada cierto tiempo (configurable).
*   **Balanceo sencillo:** Actúa como un punto de descubrimiento para que tu bot obtenga siempre un nodo funcional.
*   **Logs limpios:** Reporta qué nodos están caídos y confirma cuando todos los nodos están funcionando correctamente sin hacer spam en la consola.
*   **Configuración simple en YAML:** Fácil de agregar y quitar nodos, y definir el puerto del proxy.

## Requisitos

*   [Go](https://golang.org/dl/) (Golang) instalado en tu sistema.

## Configuración

Toda la configuración principal se realiza en el archivo `config.yml` ubicado en la raíz del proyecto.

### Ejemplo de `config.yml`

```yaml
# Puerto en el que se ejecutará la API HTTP del Proxy
puerto: "3001"

# Tiempo en segundos entre cada comprobación de los nodos
check_delay: 300 

nodos:
  - host: "lavalink.mi-servidor.com"
    port: "443"
    password: "la-contrasena-secreta"
    secure: true
    
  - host: "otro-nodo.local"
    port: "2333"
    password: "youshallnotpass"
    secure: false
```

## Instalación y Ejecución

1. Clona o descarga el proyecto.
2. Asegúrate de tener tu archivo `config.yml` configurado con tus nodos reales.
3. Abre una terminal en la carpeta del proyecto y ejecuta el servidor:

```bash
# Para ejecutarlo directamente:
go run main.go

# O para compilarlo en un ejecutable y luego correrlo:
go build
./lavaproxi # o lavaproxi.exe en Windows
```

Al iniciar, verás que el servidor carga los nodos y te dice en qué puerto está corriendo. En segundo plano comenzará a verificar qué nodos están vivos o muertos.

## Uso de la API (Ejemplo de Petición)

Una vez que el servidor esté corriendo, expondrá un endpoint para obtener un nodo disponible. 

**Endpoint:** `GET /get-node`

### Ejemplo usando `curl` o tu navegador de internet:
Si lo ejecutas en tu misma PC y dejaste el puerto en `3001`, simplemente entra a:

```bash
http://localhost:3001/get-node
```

Respuesta (JSON con los datos del nodo devuelto):
```json
{
  "host": "lavalink.mi-servidor.com",
  "port": "443",
  "password": "la-contrasena-secreta",
  "secure": true
}
```

(Nota: el comportamiento exacto de `/get-node` dependerá de cómo esté implementado internamente el balanceador en tu `handler.go`, pero siempre te devolverá información de un nodo disponible).
