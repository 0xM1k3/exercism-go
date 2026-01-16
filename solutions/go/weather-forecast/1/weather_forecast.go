// Package weather ...
package weather

var (
    // CurrentCondition, random text.
	CurrentCondition string
    // CurrentLocation, random text.
	CurrentLocation  string
)
// Forecast, random text.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
