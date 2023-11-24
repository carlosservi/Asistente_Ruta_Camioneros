package main

import (
	"Asistente_Ruta_Camioneros/src/route"
	"log"
)

//Necesito el programa principal para ejecutar el programa
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
}
