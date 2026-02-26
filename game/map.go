package game

type WeatherState string

const (
	Rain         WeatherState = "rain"
	FairSkies    WeatherState = "fair skies"
	Thunderstorm WeatherState = "thunderstorm"
)

type Map struct {
	Size        int
	Coordinates [][]int
	Name        string
	Weather     WeatherState
}

func InitializeMap() *Map {
	return &Map{
		Size:        0,
		Coordinates: [][]int{},
		Name:        "",
		Weather:     FairSkies,
	}
}
