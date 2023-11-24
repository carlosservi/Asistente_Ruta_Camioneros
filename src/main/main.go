package main

import (
	"Asistente_Ruta_Camioneros/src/route"
	"log"
)

func main() {
	// Ruta al archivo JSON
	filePath := "data/restArea.json"

	// Crear el objeto restArea
	restAreas, err := route.ReadJsonRestAreas(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// Mostrar por la terminal el objeto restAreas
	for _, restArea := range restAreas {
		log.Printf("ID: %s, Retraso: %d", restArea.Id, restArea.Retraso)
		for _, openingHour := range restArea.OpeningHours {
			log.Printf("  %s: %s - %s", openingHour.Weekday, openingHour.StartTime, openingHour.CloseTime)
		}
	}

	// Crear los objetos route
	filePath = "data/route.json"
	routes, err := route.ReadJsonRoutes(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// Mostrar por la terminal el objeto routes
	for _, route := range routes {
		log.Printf("ID: %d, Distancia total: %d", route.Id, route.TotalDistance)
		for _, areas := range route.RestAreas {
			log.Printf("Area: %s", areas)
		}
	}

	// Crear los objetos descansos
	filePath = "data/descansos.json"
	descansos, err := route.ReadJsonDescansos(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// Mostrar por la terminal el objeto descansos
	for _, descanso := range descansos {
		log.Printf("Nombre: %s, Tiempo: %d, Descanso: %d", descanso.Nombre, descanso.Tiempo, descanso.Descanso)
	}

}
