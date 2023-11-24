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
	// layout := "2006-01-02T15:04:05Z"

	// startTimeParsed, err := time.Parse(layout, startTime)
	// if err != nil {
	// 	fmt.Println("Error al analizar la cadena de tiempo:", err)
	// 	return nil
	// }

	// closeTimeParsed, err := time.Parse(layout, closeTime)
	// if err != nil {
	// 	fmt.Println("Error al analizar la cadena de tiempo:", err)
	// 	return nil
	// }

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
	Id             uint16            `json:"id"`
	ArrivalTime    time.Time         `json:"arrivalTime"`
	RestAreas      []RestArea        `json:"restAreas"`
	RouteDistances [][]int16         `json:"routeDistances"`
	RouteTimes     [][]time.Duration `json:"routeTimes"`
	TotalDistance  int16             `json:"totalDistance"`
}

func NewRoute(id uint16, arrivalTime time.Time, restAreas []RestArea, routeDistances [][]int16, routeTimes [][]time.Duration, totalDistance int16) *Route {
	return &Route{
		Id:             id,
		ArrivalTime:    arrivalTime,
		RestAreas:      restAreas,
		RouteDistances: routeDistances,
		RouteTimes:     routeTimes,
		TotalDistance:  totalDistance,
	}
}
