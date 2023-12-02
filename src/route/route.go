package route

import (
	"time"
)

// Definicion de la entidad Horarios de apertura
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
	Retraso      int            `json:"retraso"`
}

func NewRestArea(openingHours []OpeningHours, id string, retraso int) *RestArea {
	return &RestArea{
		OpeningHours: openingHours,
		Id:           id,
		Retraso:      retraso,
	}
}

// Definición de la entidad Ruta
type Route struct {
	Id                   int      `json:"id"`
	RestAreas            []string `json:"restAreas"`            //Areas de descanso con horas de apertura y cierre
	RouteKmDistances     []int    `json:"routeKmDistances"`     //Distancia entre areas de descanso en kilometros, el último valor es la distancia entre el último area de descanso y el destino
	RouteMinutesTimes    []int    `json:"routeMinutesTimes"`    //Tiempo entre areas de descanso en minutos, el último valor es la distancia en minutos entre el último area de descanso y el destino
	TotalMinutesDistance int      `json:"totalMinutesDistance"` //Distancia total de la ruta en minutos
}

func NewRoute(id int, arrivalTime time.Time, restAreas []string, routeDistances []int, routeTimes []int, totalMinutesDistance int) *Route {
	return &Route{
		Id:                   id,
		RestAreas:            restAreas,
		RouteKmDistances:     routeDistances,
		RouteMinutesTimes:    routeTimes,
		TotalMinutesDistance: totalMinutesDistance,
	}
}

type Descansos struct {
	Nombre   string `json:"nombre"`
	Tiempo   int    `json:"tiempo"`   //Tiempo en minutos hasta el siguiente descanso
	Descanso int    `json:"descanso"` //Tiempo en minutos del descanso
}

func NewDescansos(nombre string, tiempo int, descanso int) *Descansos {
	return &Descansos{
		Nombre:   nombre,
		Tiempo:   tiempo,
		Descanso: descanso,
	}
}

//Funcion de optimización de ruta
func CalculateOptimalRoute(arrivalTime time.Time, route Route, restAreas []RestArea, descansos []Descansos) ([]string, time.Time, error) {
	for _, tiempo := range route.RouteMinutesTimes {
		for _, descanso := range descansos {
			if tiempo > descanso.Tiempo {
				return []string{"No se puede hacer la ruta"}, time.Time{}, nil
			}
		}
	}
	if descansos[0].Tiempo < route.TotalMinutesDistance {
		indices := []int{}
		sumaDescansos := 0
		suma := 0
		j := 0
		i := 0
		for sumaDescansos < route.TotalMinutesDistance || i < len(route.RouteMinutesTimes) {
			if (sumaDescansos + suma + route.RouteMinutesTimes[i]) <= route.TotalMinutesDistance {
				if (suma + route.RouteMinutesTimes[i]) < descansos[j].Tiempo {
					suma += route.RouteMinutesTimes[i]
					i++
					if i == len(route.RouteMinutesTimes) {
						break
					}
				} else {
					sumaDescansos += suma
					suma = 0
					indices = append(indices, i-1)
					j = (j + 1) % len(descansos)
				}
			} else {
				break
			}
		}

		//Introducir el nombre de las restAreas en una lista
		listaParadas := []string{}
		for _, indice := range indices {
			listaParadas = append(listaParadas, route.RestAreas[indice])
		}

		//Calcular la hora de salida.
		h := 0
		tiempoTotalRuta := int(route.TotalMinutesDistance)
		for _, parada := range listaParadas {
			area := BuscarRestAreaPorID(restAreas, parada)
			if area.Retraso > descansos[h].Descanso {
				tiempoTotalRuta += area.Retraso
			} else {
				tiempoTotalRuta += descansos[h].Descanso
			}
			h = (h + 1) % len(descansos)
		}
		departureTime := arrivalTime.Add(-time.Minute * time.Duration(tiempoTotalRuta))
		departureTimeAux := departureTime
		abierto := false
		//Comprobar si están abiertas las áreas de descanso, en caso contrario, llamar recursivamente a la función, eliminando de las posibles areas las que están cerradas y ajustando el tiempo y los kilometros de la ruta
		for i, parada := range route.RestAreas {
			if ExisteParada(listaParadas, parada) {
				departureTimeAux = departureTimeAux.Add(time.Minute * time.Duration(route.RouteMinutesTimes[i]))
				dia := departureTimeAux.Weekday()
				area := BuscarRestAreaPorID(restAreas, parada)
				horario := GetOpeningHoursByWeekday(area, dia)
				if horario == nil {
					abierto = false
				} else {
					abierto = ComprobarSiEstaAbiertaElArea(horario, departureTimeAux)
				}
				if !abierto {
					//Necesito eliminar de la ruta la parada que está cerrada, además de incrementarle el tiempo a la siguiente y los kilometros
					km := route.RouteKmDistances[i]
					minutos := route.RouteMinutesTimes[i]

					//Eliminar de la ruta
					route.RouteKmDistances = append(route.RouteKmDistances[:i], route.RouteKmDistances[i+1:]...)
					route.RouteMinutesTimes = append(route.RouteMinutesTimes[:i], route.RouteMinutesTimes[i+1:]...)
					//Incrementar el tiempo y los kilometros de la siguiente parada
					route.RouteKmDistances[i] += km
					route.RouteMinutesTimes[i] += minutos
					//Llamar recursivamente a la función
					return CalculateOptimalRoute(arrivalTime, route, restAreas, descansos)
				}
			} else {
				departureTimeAux = departureTimeAux.Add(time.Minute * time.Duration(route.RouteMinutesTimes[i]))
			}
		}

		return listaParadas, departureTime, nil
	} else {
		//Se puede hacer la ruta sin descanso
		departureTime := arrivalTime.Add(-time.Minute * time.Duration(route.TotalMinutesDistance))
		return []string{"No es necesario parar"}, departureTime, nil
	}
}
