package main

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
	distance     uint16
}

func (r RestArea) OpeningHours() []OpeningHours {
	return r.openingHours
}

func (r RestArea) Distance() uint16 {
	return r.distance
}

}

func NewRestArea(openingHours []OpeningHours, latitude float64, longitude float64) *RestArea {
	return &RestArea{
		openingHours: openingHours,
		distance:     distance,
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
