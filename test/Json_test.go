package test

import (
	"Asistente_Ruta_Camioneros/src/route"
	"testing"
	"time"
)

func TestParseRestAreaJSON(t *testing.T) {
	t.Run("Parsing a valid RestArea JSON", func(t *testing.T) {
		// Given a RestArea JSON with id "Granada" and retraso 30
		filePath := "../data/restAreaTest.json"

		// When I parse the RestArea JSON
		parsedRestArea, err := route.ReadJsonRestAreas(filePath)
		if err != nil {
			t.Fatalf("Error parsing RestArea JSON: %v", err)
		}

		// Then the parsed RestArea should have id "Granada" and retraso 30
		expectedRoute := route.RestArea{
			Id:      "Granada",
			Retraso: time.Duration(30),
			OpeningHours: []route.OpeningHours{
				{
					Weekday:   "time.Saturday",
					StartTime: "2023-11-23T10:00:00Z",
					CloseTime: "2023-11-23T23:00:00Z",
				},
				{
					Weekday:   "time.Sunday",
					StartTime: "2023-11-23T10:00:00Z",
					CloseTime: "2023-11-23T17:00:00Z",
				},
			},
		}

		if !restAreaEqual(parsedRestArea[0], expectedRoute) {
			t.Errorf("Parsed RestArea does not match the expected values")
		}
	})
}

func TestParseRouteJSON(t *testing.T) {
	t.Run("Parsing a valid Route JSON", func(t *testing.T) {
		// Given a Route JSON with id "1", restAreas "Granada", "Jaen" and "Ciudad Real", totalDistance 480, routeDistances 80, 200 and 200 and routeTimes 60, 150 and 150
		filePath := "../data/routeTest.json"

		// When I parse the Route JSON
		parsedRoute, err := route.ReadJsonRoutes(filePath)
		if err != nil {
			t.Fatalf("Error parsing Route JSON: %v", err)
		}

		// Then the parsed Route should have id "1", restAreas "Granada", "Jaen" and "Ciudad Real", totalDistance 480, routeDistances 80, 200 and 200 and routeTimes 60, 150 and 150
		expectedRoute := route.Route{
			Id:             1,
			RestAreas:      []string{"Granada", "Jaen", "Ciudad Real"},
			RouteDistances: []int16{80, 200, 200},
			RouteTimes:     []int16{60, 150, 150},
			TotalDistance:  480,
		}

		if !routeEqual(parsedRoute[0], expectedRoute) {
			t.Errorf("Parsed Route does not match the expected values")
		}
	})
}

func TestParseDescansosJSON(t *testing.T) {
	t.Run("Parsing a valid Descansos JSON", func(t *testing.T) {
		// Given a Descanso JSON with Nombre "Descanso 1" and Tiempo 270 and Descanso 45
		filePath := "../data/descansos.json"

		// When I parse the Descansos JSON
		parsedDescanso, err := route.ReadJsonDescansos(filePath)
		if err != nil {
			t.Fatalf("Error parsing Descansos JSON: %v", err)
		}

		// Then the parsed Descansos should have Nombre "Descanso 1" and Tiempo 270 and Descanso 45
		expectedDescanso := route.Descansos{
			Nombre:   "Descanso 1",
			Tiempo:   270,
			Descanso: 45,
		}

		if !descansoEqual(parsedDescanso[0], expectedDescanso) {
			t.Errorf("Parsed Route does not match the expected values")
		}
	})
}

func restAreaEqual(ra1, ra2 route.RestArea) bool {
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
		if !openingHoursEqual(ra1.OpeningHours[i], ra2.OpeningHours[i]) {
			return false
		}
	}

	return true
}

func openingHoursEqual(oh1, oh2 route.OpeningHours) bool {
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

func routeEqual(r1, r2 route.Route) bool {
	// Compara los campos uno por uno
	if r1.Id != r2.Id {
		return false
	}

	if r1.TotalDistance != r2.TotalDistance {
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

	// Compara las listas de routeDistances
	if len(r1.RouteDistances) != len(r2.RouteDistances) {
		return false
	}

	for i := range r1.RouteDistances {
		if r1.RouteDistances[i] != r2.RouteDistances[i] {
			return false
		}
	}

	// Compara las listas de routeTimes
	if len(r1.RouteTimes) != len(r2.RouteTimes) {
		return false
	}

	for i := range r1.RouteTimes {
		if r1.RouteTimes[i] != r2.RouteTimes[i] {
			return false
		}
	}

	return true
}

func descansoEqual(d1, d2 route.Descansos) bool {
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
