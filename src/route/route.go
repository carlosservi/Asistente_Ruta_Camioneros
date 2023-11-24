package route

import (
	"time"
)

type OpeningHours struct {
	StartTime string `json:"startTime"`
	CloseTime string `json:"closeTime"`
	Weekday   string `json:"weekday"`
}

func NewOpeningHours(startTime string, closeTime string, weekday string) *OpeningHours {
	return &OpeningHours{
		StartTime: startTime,
		CloseTime: closeTime,
		Weekday:   weekday,
	}
}

// Definicion de la entidad Area de descanso
type RestArea struct {
	OpeningHours []OpeningHours `json:"openingHours"`
	Id           string         `json:"id"`
	Retraso      time.Duration  `json:"retraso"`
}

func NewRestArea(openingHours []OpeningHours, id string, retraso uint16) *RestArea {
	return &RestArea{
		OpeningHours: openingHours,
		Id:           id,
		Retraso:      time.Duration(retraso),
	}
}

// Definición de la entidad Ruta
type Route struct {
	Id             uint16    `json:"id"`
	ArrivalTime    time.Time `json:"arrivalTime"`
	RestAreas      []string  `json:"restAreas"`
	RouteDistances []int16   `json:"routeDistances"`
	RouteTimes     []int16   `json:"routeTimes"`
	TotalDistance  int16     `json:"totalDistance"`
}

func NewRoute(id uint16, arrivalTime time.Time, restAreas []string, routeDistances []int16, routeTimes []int16, totalDistance int16) *Route {
	return &Route{
		Id:             id,
		ArrivalTime:    arrivalTime,
		RestAreas:      restAreas,
		RouteDistances: routeDistances,
		RouteTimes:     routeTimes,
		TotalDistance:  totalDistance,
	}
}

type Descansos struct {
	Nombre   string `json:"nombre"`
	Tiempo   int16  `json:"tiempo"`
	Descanso int16  `json:"descanso"`
}

func NewDescansos(nombre string, tiempo int16, descanso int16) *Descansos {
	return &Descansos{
		Nombre:   nombre,
		Tiempo:   tiempo,
		Descanso: descanso,
	}
}

//Necesito un algoritmo que pasandole una ruta, una lista de areas de descanso y , me devuelva las areas de descanso que se pueden visitar, teniendo en cuenta sus horarios de apertura y cierre
