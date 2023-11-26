package route

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

func ReadJsonRestAreas(filePath string) ([]RestArea, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var restAreas []RestArea
	err = json.Unmarshal(data, &restAreas)
	if err != nil {
		return nil, err
	}
	return restAreas, nil
}

func ReadJsonRoutes(filePath string) ([]Route, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var routes []Route
	err = json.Unmarshal(data, &routes)
	if err != nil {
		return nil, err
	}
	return routes, nil
}

func ReadJsonDescansos(filePath string) ([]Descansos, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var descansos []Descansos
	err = json.Unmarshal(data, &descansos)
	if err != nil {
		return nil, err
	}
	return descansos, nil
}

//Función para buscar restArea por ID
func BuscarRestAreaPorID(array []RestArea, idBuscado string) RestArea {
	for _, restArea := range array {
		if restArea.Id == idBuscado {
			return restArea
		}
	}
	return RestArea{}
}

//Función para ver es una de las paradas elegida para hacer el descanso, lo que hace es ver si está en el array de paradas
func ExisteParada(array []string, parada string) bool {
	for _, area := range array {
		if area == parada {
			return true
		}
	}
	return false
}

//Función que devuelve el horario de una área de servicio un dia concreto
func GetOpeningHoursByWeekday(restArea RestArea, weekday time.Weekday) *OpeningHours {
	for i := range restArea.OpeningHours {
		if restArea.OpeningHours[i].Weekday == weekday.String() {
			return &restArea.OpeningHours[i]
		}
	}
	return nil
}

//Función que comprueba si una área de servicio está abierta un dia concreto y en un horario
func ComprobarSiEstaAbiertaElArea(horario *OpeningHours, arrivalTime time.Time) bool {
	start, _ := time.Parse(time.RFC3339, horario.StartTime)
	close, _ := time.Parse(time.RFC3339, horario.CloseTime)

	if arrivalTime.Hour() >= start.Hour() && arrivalTime.Hour() < close.Hour() {
		return true
	}
	return false
}
