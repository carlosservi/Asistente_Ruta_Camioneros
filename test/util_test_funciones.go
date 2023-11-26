package test

import (
	"Asistente_Ruta_Camioneros/src/route"
	"testing"
)

//Función para comparar si dos restAreas son iguales
func RestAreaEqual(ra1, ra2 route.RestArea) bool {
	// Compara los campos uno por uno
	if ra1.Id != ra2.Id {
		return false
	}

	if ra1.Retraso != ra2.Retraso {
		return false
	}

	// Compara las listas de OpeningHours
	if len(ra1.OpeningHours) != len(ra2.OpeningHours) {
		return false
	}

	for i := range ra1.OpeningHours {
		if !OpeningHoursEqual(ra1.OpeningHours[i], ra2.OpeningHours[i]) {
			return false
		}
	}

	return true
}

//Funcion para comparar si dos OpeningHours son iguales
func OpeningHoursEqual(oh1, oh2 route.OpeningHours) bool {
	// Compara los campos uno por uno
	if oh1.StartTime != oh2.StartTime {
		return false
	}

	if oh1.CloseTime != oh2.CloseTime {
		return false
	}

	if oh1.Weekday != oh2.Weekday {
		return false
	}

	return true
}

//Función para comparar si dos rutas son iguales
func RouteEqual(r1, r2 route.Route) bool {
	// Compara los campos uno por uno
	if r1.Id != r2.Id {
		return false
	}

	if r1.TotalMinutesDistance != r2.TotalMinutesDistance {
		return false
	}

	// Compara las listas de restAreas
	if len(r1.RestAreas) != len(r2.RestAreas) {
		return false
	}

	for i := range r1.RestAreas {
		if r1.RestAreas[i] != r2.RestAreas[i] {
			return false
		}
	}

	// Compara las listas de routeKmDistances
	if len(r1.RouteKmDistances) != len(r2.RouteKmDistances) {
		return false
	}

	for i := range r1.RouteKmDistances {
		if r1.RouteKmDistances[i] != r2.RouteKmDistances[i] {
			return false
		}
	}

	// Compara las listas de routeMinutesTimes
	if len(r1.RouteMinutesTimes) != len(r2.RouteMinutesTimes) {
		return false
	}

	for i := range r1.RouteMinutesTimes {
		if r1.RouteMinutesTimes[i] != r2.RouteMinutesTimes[i] {
			return false
		}
	}

	return true
}

//Función para comparar si dos descansos son iguales
func DescansoEqual(d1, d2 route.Descansos) bool {
	// Compara los campos uno por uno
	if d1.Nombre != d2.Nombre {
		return false
	}

	if d1.Tiempo != d2.Tiempo {
		return false
	}

	if d1.Descanso != d2.Descanso {
		return false
	}

	return true
}

func assertEqual(t *testing.T, result, expected interface{}, message string) {
	if result != expected {
		t.Errorf("%s\nGot: %v\nExpected: %v", message, result, expected)
	}
}
