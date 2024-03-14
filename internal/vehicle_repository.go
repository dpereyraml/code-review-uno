package internal

// VehicleRepository is an interface that represents a vehicle repository
type VehicleRepository interface {
	// FindAll is a method that returns a map of all vehicles
	FindAll() (v map[int]Vehicle, err error)
	// AddVehicle is a method that adds a vehicle to the repository
	AddVehicle(v *Vehicle) (err error)
	// GetByColorAndYear is a method that returns vehicles by color and year
	GetByColorAndYear(color string, year int) (v map[int]Vehicle, err error)
}
