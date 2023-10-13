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

type RestArea struct {
	openingHours []OpeningHours
	latitude     float64
	longitude    float64
}

func (r RestArea) OpeningHours() []OpeningHours {
	return r.openingHours
}

func (r RestArea) Latitude() float64 {
	return r.latitude
}

func (r RestArea) Longitude() float64 {
	return r.longitude
}

func NewRestArea(openingHours []OpeningHours, latitude float64, longitude float64) *RestArea {
	return &RestArea{
		openingHours: openingHours,
		latitude:     latitude,
		longitude:    longitude,
	}
}
