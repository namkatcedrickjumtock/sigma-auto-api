package events

import (
	"fmt"
	"math"

	models "github.com/namkatcedrickjumtock/sigma-auto-api/internal/models/event"
)

func getTicketPrice(event []models.Ticket, ticketType string, qty int64) (int64, int64, error) {
	var ticketPrice int64

	var ticketSum int64

	for _, eventTickets := range event {
		if eventTickets.Name == ticketType {
			ticketSum = eventTickets.Price * qty
			ticketPrice = eventTickets.Price

			return ticketSum, ticketPrice, nil
		}
	}
	//nolint:goerr113
	return 0, 0, fmt.Errorf("tickets not found")
}

// helper function remove items from slice of any type orderly.
func Remove[T comparable](slice []T, index int) []T {
	return append(slice[:index], slice[index+1:]...)
}

// :::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
// :::                                                                         :::
// :::  This routine calculates the distance between two points (given the     :::
// :::  latitude/longitude of those points). It is based on free code used to  :::
// :::  calculate the distance between two locations using GeoDataSource(TM)   :::
// :::  products.                                                              :::
// :::                                                                         :::
// :::  Definitions:                                                           :::
// :::    South latitudes are negative, east longitudes are positive           :::
// :::                                                                         :::
// :::  Passed to function:                                                    :::
// :::    lat1, lon1 = Latitude and Longitude of point 1 (in decimal degrees)  :::
// :::    lat2, lon2 = Latitude and Longitude of point 2 (in decimal degrees)  :::
// :::    optional: unit = the unit you desire for results                     :::
// :::           where: 'M' is statute miles (default, or omitted)             :::
// :::                  'K' is kilometers                                      :::
// :::                  'N' is nautical miles                                  :::
// ::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::: .
func distance(lat1 float64, lng1 float64, lat2 float64, lng2 float64, unit ...string) float64 {
	radlat1 := math.Pi * lat1 / 180
	radlat2 := math.Pi * lat2 / 180

	theta := lng1 - lng2
	radtheta := math.Pi * theta / 180

	dist := math.Sin(radlat1)*math.Sin(radlat2) + math.Cos(radlat1)*math.Cos(radlat2)*math.Cos(radtheta)
	if dist > 1 {
		dist = 1
	}

	dist = math.Acos(dist)
	dist = dist * 180 / math.Pi
	dist = dist * 60 * 1.1515

	if len(unit) > 0 {
		if unit[0] == "K" {
			dist *= 1.609344
		} else if unit[0] == "N" {
			dist *= 0.8684
		}
	}

	return dist
}
