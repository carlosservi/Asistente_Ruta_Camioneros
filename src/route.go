package route

import (
	"time"
)

type OpeningHours struct {
	Day       string
	StartTime string
	EndTime   string
}

func NewOpeningHours(day string, startTime string, endTime string) *OpeningHours {
	return &OpeningHours{
		Day:       day,
		StartTime: startTime,
		EndTime:   endTime,
	}
}

// Definicion de la entidad Area de descanso
type RestArea struct {
	openingHours []OpeningHours
	timeToArrive time.Duration
	distance     uint16
}

func (r RestArea) OpeningHours() []OpeningHours {
	return r.openingHours
}

func (r RestArea) Distance() uint16 {
	return r.distance
}

func (r RestArea) TimeToArrive() time.Duration {
	return r.timeToArrive
}

func NewRestArea(openingHours []OpeningHours, distance uint16, timeToArrive time.Duration) *RestArea {
	return &RestArea{
		openingHours: openingHours,
		distance:     distance,
		timeToArrive: timeToArrive,
	}
}

// Definición de la entidad Ruta
type Route struct {
	id        string
	restAreas []RestArea
}

func (r Route) Id() string {
	return r.id
}

func (r Route) RestAreas() []RestArea {
	return r.restAreas
}

func NewRoute(id string, restAreas []RestArea) *Route {
	return &Route{
		id:        id,
		restAreas: restAreas,
	}
}
