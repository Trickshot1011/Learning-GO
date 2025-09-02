// Package weather provides tools to predict weather conditions for a given location.
package weather

// CurrentCondition represents weather condition in string format.
var CurrentCondition string

// CurrentLocation represents a location in string format.
var CurrentLocation string

// Forecast return a string represeting the current weather condition for a provided location and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
