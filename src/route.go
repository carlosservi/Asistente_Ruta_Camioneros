package route

import (
	"time"
)

type OpeningHours struct {
	StartTime time.Time
	CloseTime time.Time
	Weekday   time.Weekday
}

func NewOpeningHours(startTime time.Time, closeTime time.Time, weekday time.Weekday) *OpeningHours {
	return &OpeningHours{
		StartTime: startTime,
		CloseTime: closeTime,
		Weekday:   weekday,
	}
}

// Definicion de la entidad Area de descanso
type RestArea struct {
	openingHours []OpeningHours
	id           uint8
	retraso      time.Duration
}

func NewRestArea(openingHours []OpeningHours, id uint8, retraso time.Duration) *RestArea {
	return &RestArea{
		openingHours: openingHours,
		id:           id,
		retraso:      retraso,
	}
}

// Definición de la entidad Ruta
type Route struct {
	id             string
	arrivalTime    time.Time
	restAreas      []RestArea
	routeDistances [][]int16
	routeTimes     [][]time.Duration
	totalDistance  int16
}

func NewRoute(id string, arrivalTime time.Time, restAreas []RestArea, routeDistances [][]int16, routeTimes [][]time.Duration, totalDistance int16) *Route {
	return &Route{
		id:             id,
		arrivalTime:    arrivalTime,
		restAreas:      restAreas,
		routeDistances: routeDistances,
		routeTimes:     routeTimes,
		totalDistance:  totalDistance,
	}
}
