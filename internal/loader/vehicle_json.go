package loader

import (
	"app/internal"
	"encoding/json"
	"os"
)

// NewVehicleJSONFile is a function that returns a new instance of VehicleJSONFile
func NewVehicleJSONFile(path string) *VehicleJSONFile {
	return &VehicleJSONFile{
		path: path,
	}
}

// VehicleJSONFile is a struct that implements the LoaderVehicle interface
type VehicleJSONFile struct {
	// path is the path to the file that contains the vehicles in JSON format
	path string
}

// VehicleJSON is a struct that represents a vehicle in JSON format
type VehicleJSON struct {
	Id              int     `json:"id"`
	Brand           string  `json:"brand"`
	Model           string  `json:"model"`
	Registration    string  `json:"registration"`
	Color           string  `json:"color"`
	FabricationYear int     `json:"year"`
	Capacity        int     `json:"passengers"`
	MaxSpeed        float64 `json:"max_speed"`
	FuelType        string  `json:"fuel_type"`
	Transmission    string  `json:"transmission"`
	Weight          float64 `json:"weight"`
	Height          float64 `json:"height"`
	Length          float64 `json:"length"`
	Width           float64 `json:"width"`
}

// Load is a method that loads the vehicles
func (l *VehicleJSONFile) Load() (v map[int]internal.Vehicle, err error) {
	// open file
	file, err := os.Open(l.path)
	if err != nil {
		return
	}
	defer file.Close()

	// decode file
	var vehiclesJSON []VehicleJSON
	err = json.NewDecoder(file).Decode(&vehiclesJSON)
	if err != nil {
		return
	}

	// serialize vehicles
	v = make(map[int]internal.Vehicle)
	for _, vh := range vehiclesJSON {
		v[vh.Id] = internal.Vehicle{
			Id: vh.Id,
			VehicleAttributes: internal.VehicleAttributes{
				Brand:           vh.Brand,
				Model:           vh.Model,
				Registration:    vh.Registration,
				Color:           vh.Color,
				FabricationYear: vh.FabricationYear,
				Capacity:        vh.Capacity,
				MaxSpeed:        vh.MaxSpeed,
				FuelType:        vh.FuelType,
				Transmission:    vh.Transmission,
				Weight:          vh.Weight,
				Dimensions: internal.Dimensions{
					Height: vh.Height,
					Length: vh.Length,
					Width:  vh.Width,
				},
			},
		}
	}

	return
}

// AddVehicle is a method that adds a vehicle to the file
/* func (l *VehicleJSONFile) AddVehicle(vh *internal.Vehicle) (err error) {
	// open file
	file, err := os.Open(l.path)
	if err != nil {
		return
	}
	defer file.Close()
	// defer prevents the file from being closed before the function returns

	// decode file
	var vehiclesJSON []VehicleJSON
	// decode file to vehiclesJSON
	err = json.NewDecoder(file).Decode(&vehiclesJSON)
	if err != nil {
		return
	}

	// serialize vehicles
	vehiclesJSON = append(vehiclesJSON, VehicleJSON{
		Id:              vh.Id,
		Brand:           vh.Brand,
		Model:           vh.Model,
		Registration:    vh.Registration,
		Color:           vh.Color,
		FabricationYear: vh.FabricationYear,
		Capacity:        vh.Capacity,
		MaxSpeed:        vh.MaxSpeed,
		FuelType:        vh.FuelType,
		Transmission:    vh.Transmission,
		Weight:          vh.Weight,
		Height:          vh.Dimensions.Height,
		Length:          vh.Dimensions.Length,
		Width:           vh.Dimensions.Width,
	})

	// encode file
	// file is a pointer to the file
	// file, err = os.OpenFile(l.path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	// os.Create creates the file if it does not exist
	// example of os.Create: file, err := os.Create("file.txt")
	os.Create(l.path)

	if err != nil {
		return
	}
	defer file.Close()
	err = json.NewEncoder(file).Encode(vehiclesJSON)
	if err != nil {
		return
	}

	return
} */
