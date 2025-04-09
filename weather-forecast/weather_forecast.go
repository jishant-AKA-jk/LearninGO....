// Package weather ...
package weather

// CurrentCondition is a string stores current condition for location.
var CurrentCondition string

// CurrentLocation is a string stores weather data.
var CurrentLocation string

// Forecast function will return prediction about weather.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
