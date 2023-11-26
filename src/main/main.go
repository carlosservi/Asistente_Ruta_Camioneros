package main

import (
	"Asistente_Ruta_Camioneros/src/route"
	"fmt"
	"log"
	"time"
)

func main() {
	// Ruta al archivo JSON
	filePath := "data/rest_area.json"

	// Crear el objeto restArea
	restAreas, err := route.ReadJsonRestAreas(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// Crear los objetos route
	filePath = "data/route.json"
	routes, err := route.ReadJsonRoutes(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// Crear los objetos descansos
	filePath = "data/descansos.json"
	descansos, err := route.ReadJsonDescansos(filePath)
	if err != nil {
		log.Fatal(err)
	}

	dateTimeString := "2021-05-10T21:00:00Z"
	arrivalTime, err := time.Parse(time.RFC3339, dateTimeString)
	if err != nil {
		fmt.Println("Error al analizar la cadena de tiempo:", err)
		return
	}

	// Calcular la ruta óptima
	optimalRoute, departureTime, err := route.CalculateOptimalRoute(arrivalTime, routes[1], restAreas, descansos)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Áreas de descanso en las que parar:", optimalRoute)
	fmt.Println("Día y Hora de salida:", departureTime.Format("2006-01-02 15:04"))
	fmt.Println("Día y Hora de llegada:", arrivalTime.Format("2006-01-02 15:04"))

}
