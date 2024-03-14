package repository

import (
	"app/internal"
	"encoding/json"
	"fmt"
	"os"
)

// NewVehicleMap is a function that returns a new instance of VehicleMap
func NewVehicleMap(db map[int]internal.Vehicle) *VehicleMap {
	// default db
	defaultDb := make(map[int]internal.Vehicle)
	if db != nil {
		defaultDb = db
	}
	return &VehicleMap{db: defaultDb}
}

// Dimensions is a struct that represents a dimension in 3d
type Dimensions struct {
	// Height is the height of the dimension
	Height float64
	// Length is the length of the dimension
	Length float64
	// Width is the width of the dimension
	Width float64
}

// VehicleAttributes is a struct that represents the attributes of a vehicle
type VehicleFormatJSON struct {
	// ID
	ID int `json:"id"`
	// Brand is the brand of the vehicle
	Brand string `json:"brand"`
	// Model is the model of the vehicle
	Model string `json:"model"`
	// Registration is the registration of the vehicle
	Registration string `json:"registration"`
	// Color is the color of the vehicle
	Color string `json:"color"`
	// FabricationYear is the fabrication year of the vehicle
	FabricationYear int `json:"year"`
	// Capacity is the capacity of people of the vehicle
	Capacity int `json:"passengers"`
	// MaxSpeed is the maximum speed of the vehicle
	MaxSpeed float64 `json:"max_speed"`
	// FuelType is the fuel type of the vehicle
	FuelType string `json:"fuel_type"`
	// Transmission is the transmission of the vehicle
	Transmission string `json:"transmission"`
	// Weight is the weight of the vehicle
	Weight float64 `json:"weight"`
	// Dimensions is the dimensions of the vehicle
	// Height is the height of the dimension
	Height float64 `json:"height"`
	// Length is the length of the dimension
	Length float64 `json:"length"`
	// Width is the width of the dimension
	Width float64 `json:"width"`
}

// VehicleMap is a struct that represents a vehicle repository
type VehicleMap struct {
	// db is a map of vehicles
	db map[int]internal.Vehicle
}

// FindAll is a method that returns a map of all vehicles
func (r *VehicleMap) FindAll() (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	// copy db
	for key, value := range r.db {
		v[key] = value
	}

	return
}

func (r *VehicleMap) AddVehicle(v *internal.Vehicle) (err error) {
	// set id
	(*v).Id = len(r.db) + 1 // auto increment id
	r.db[(*v).Id] = *v      // add vehicle to db

	file, err := os.Create("../docs/db/vehicles.json")
	if err != nil {
		fmt.Println("Error creating file")
		return
	}
	defer file.Close()

	// r.db to json
	var vehicleFormat = []VehicleFormatJSON{}

	for _, value := range r.db {

		vehicleFormat = append(vehicleFormat, VehicleFormatJSON{
			ID:              value.Id,
			Brand:           value.Brand,
			Model:           value.Model,
			Registration:    value.Registration,
			Color:           value.Color,
			FabricationYear: value.FabricationYear,
			Capacity:        value.Capacity,
			MaxSpeed:        value.MaxSpeed,
			FuelType:        value.FuelType,
			Transmission:    value.Transmission,
			Weight:          value.Weight,
			Height:          value.Dimensions.Height,
			Length:          value.Dimensions.Length,
			Width:           value.Dimensions.Width,
		})
	}
	fmt.Println(vehicleFormat)

	vehiclesJSON, err := json.Marshal(vehicleFormat)
	if err != nil {
		fmt.Println("Error marshalling")
		return
	}

	// write to file
	_, err = file.Write(vehiclesJSON)
	if err != nil {
		fmt.Println("Error writing to file")
		return
	}
	return

}

// GetByColorAndYear is a method that returns vehicles by color and fabrication year
func (r *VehicleMap) GetByColorAndYear(color string, year int) (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	// filter vehicles
	for key, value := range r.db {
		if value.Color == color && value.FabricationYear == year {
			v[key] = value
		}
	}

	return
}
