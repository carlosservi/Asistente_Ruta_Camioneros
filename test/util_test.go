package test

import (
	"Asistente_Ruta_Camioneros/src/route"
	"testing"
	"time"
)

func TestParseRestAreaJSON(t *testing.T) {
	t.Run("Parsing a valid RestArea JSON", func(t *testing.T) {
		// Given a RestArea JSON with id "Granada" and retraso 30
		filePath := "../data/rest_area_test.json"

		// When I parse the RestArea JSON
		parsedRestArea, err := route.ReadJsonRestAreas(filePath)
		if err != nil {
			t.Fatalf("Error parsing RestArea JSON: %v", err)
		}

		// Then the parsed RestArea should have id "Granada" and retraso 30
		expectedRestArea := route.RestArea{
			Id:      "Granada",
			Retraso: 30,
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

		if !RestAreaEqual(parsedRestArea[0], expectedRestArea) {
			t.Errorf("Parsed RestArea does not match the expected values")
		}
	})
}

func TestParseRouteJSON(t *testing.T) {
	t.Run("Parsing a valid Route JSON", func(t *testing.T) {
		// Given a Route JSON with id "1", restAreas "Granada", "Jaen" and "Ciudad Real", totalMinutesDistance 480, routeKmDistances 80, 200 and 200 and routeMinutesTimes 60, 150 and 150
		filePath := "../data/route_test.json"

		// When I parse the Route JSON
		parsedRoute, err := route.ReadJsonRoutes(filePath)
		if err != nil {
			t.Fatalf("Error parsing Route JSON: %v", err)
		}

		// Then the parsed Route should have id "1", restAreas "Granada", "Jaen" and "Ciudad Real", totalMinutesDistance 480, routeKmDistances 80, 200 and 200 and routeMinutesTimes 60, 150 and 150
		expectedRoute := route.Route{
			Id:                   1,
			RestAreas:            []string{"Granada", "Jaen", "Ciudad Real"},
			RouteKmDistances:     []uint16{80, 200, 200},
			RouteMinutesTimes:    []uint16{60, 150, 150},
			TotalMinutesDistance: 480,
		}

		if !RouteEqual(parsedRoute[0], expectedRoute) {
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

		if !DescansoEqual(parsedDescanso[0], expectedDescanso) {
			t.Errorf("Parsed Route does not match the expected values")
		}
	})
}

func TestBuscarRestAreaPorID(t *testing.T) {
	t.Run("Searching for RestArea by ID", func(t *testing.T) {
		// Given an array of RestAreas and an ID to search
		filePath := "../data/rest_area_test.json"
		restAreas, _ := route.ReadJsonRestAreas(filePath)

		idBuscado := "Granada" // ID to search

		// When searching for RestArea by ID
		foundRestArea := route.BuscarRestAreaPorID(restAreas, idBuscado)
		// Then the found RestArea should have the equals restArea
		expectedRestArea := route.RestArea{
			Id:      "Granada",
			Retraso: 30,
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

		if !RestAreaEqual(foundRestArea, expectedRestArea) {
			t.Errorf("Found RestArea does not match the expected values")
		}
	})
}

func TestExisteParada(t *testing.T) {
	t.Run("Check if the stop exists in the list of obtained stops ", func(t *testing.T) {
		// Given an array of stops and a stop to check
		array := []string{"Granada", "Jaen", "Ciudad Real"}
		parada := "Jaen"

		// When checking if the stop exists
		exists := route.ExisteParada(array, parada)

		// Then the result should be true, as "Jaen" exists in the array
		expectedResult := true

		if exists != expectedResult {
			t.Errorf("Expected %t but got %t", expectedResult, exists)
		}
	})

	t.Run("Check if the stop does not exist in the list of obtained stops", func(t *testing.T) {
		// Given an array of stops and a stop to check
		array := []string{"Granada", "Ciudad Real"}
		parada := "Jaen"

		// When checking if the stop exists
		exists := route.ExisteParada(array, parada)

		// Then the result should be false, as "Jaen" does not exist in the array
		expectedResult := false

		if exists != expectedResult {
			t.Errorf("Expected %t but got %t", expectedResult, exists)
		}
	})
}

func TestGetOpeningHoursByWeekday(t *testing.T) {
	// Given a RestArea with specific opening hours
	openingHours := []route.OpeningHours{
		{
			StartTime: "08:00",
			CloseTime: "18:00",
			Weekday:   "Monday",
		},
		{
			StartTime: "09:00",
			CloseTime: "17:00",
			Weekday:   "Tuesday",
		},
	}

	restArea := route.RestArea{
		OpeningHours: openingHours,
	}

	// When trying to get the opening hours for a specific weekday
	weekdayToCheck := time.Monday
	result := route.GetOpeningHoursByWeekday(restArea, weekdayToCheck)

	// Then the result should not be nil, indicating that the opening hours were found
	if result == nil {
		t.Errorf("Opening hours for %s should not be nil", weekdayToCheck.String())
	}

	// Optionally, you can check if the actual values match your expectations
	expectedOpeningHours := openingHours[0]
	if *result != expectedOpeningHours {
		t.Errorf("Opening hours for %s do not match the expected values", weekdayToCheck.String())
	}
}

func TestComprobarSiEstaAbiertaElArea(t *testing.T) {
	// Given an OpeningHours struct with specific start and close times
	horario := &route.OpeningHours{
		StartTime: "2023-11-23T08:00:00Z",
		CloseTime: "2023-11-23T18:00:00Z",
		Weekday:   "Monday", // No se usa en esta función, pero puedes incluirlo si es necesario
	}

	// When trying to check if the area is open at a specific arrival time
	arrivalTime := time.Date(2023, 11, 23, 12, 0, 0, 0, time.UTC) // Ajusta la hora según tus necesidades

	result := route.ComprobarSiEstaAbiertaElArea(horario, arrivalTime)

	// Then the result should be true, indicating that the area is open
	if !result {
		t.Errorf("The area should be open at %v", arrivalTime)
	}

	// Optionally, you can test with different arrival times and expected results
	arrivalTime2 := time.Date(2023, 11, 23, 19, 0, 0, 0, time.UTC)
	result2 := route.ComprobarSiEstaAbiertaElArea(horario, arrivalTime2)
	if result2 {
		t.Errorf("The area should not be open at %v", arrivalTime2)
	}
}
