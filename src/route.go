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
}

func (r RestArea) OpeningHours() []OpeningHours {
	return r.openingHours
}

func (r RestArea) Id() uint8 {
	return r.id
}

func NewRestArea(openingHours []OpeningHours, id uint8) *RestArea {
	return &RestArea{
		openingHours: openingHours,
		id:           id,
	}
}

// Definicion de las Distancias entre areas de descanso
type RestAreaDistances struct {
	distances [][]uint16
}

func (r RestAreaDistances) Distance(origin uint8, destination uint8) uint16 {
	return r.distances[origin][destination]
}

func NewRestAreaDistances(distances [][]uint16) *RestAreaDistances {
	return &RestAreaDistances{
		distances: distances,
	}
}

// Definición de la entidad Ruta
type Route struct {
	id               string
	arrivalTime      time.Time
	restAreas        []RestArea
	timesToRestAreas []time.Duration
}

func (r Route) Id() string {
	return r.id
}

func (r Route) ArrivalTime() time.Time {
	return r.arrivalTime
}

func (r Route) RestAreas() []RestArea {
	return r.restAreas
}

func (r Route) TimesToRestAreas() []time.Duration {
	return r.timesToRestAreas
}

func NewRoute(id string, arrivalTime time.Time, restAreas []RestArea, timesToRestAreas []time.Duration) *Route {
	return &Route{
		id:               id,
		arrivalTime:      arrivalTime,
		restAreas:        restAreas,
		timesToRestAreas: timesToRestAreas,
	}
}
