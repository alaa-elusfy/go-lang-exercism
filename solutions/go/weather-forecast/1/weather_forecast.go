// Package weather provides weather in Goblinocus country.
package weather

var (
    // CurrentCondition represents the current condition of weather. 
	CurrentCondition string
    // CurrentLocation represents the current location. 
	CurrentLocation  string
)

// Forecast returns a string that describes the currennt condition of weather in a location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
